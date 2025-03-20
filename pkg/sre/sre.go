package sre

import (
	"container/list"
	"context"
	"errors"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type googleSlide struct {
	sreSlide *list.List
	//滑动窗口大小
	interval int64
	mutex    sync.Mutex
	//客户端成功请求量的系数
	k float64
	// 缓存当前窗口的统计数据
	totalReq    float64
	totalAccept float64
}

type slideVal struct {
	//客户端请求时间
	time int64
	//客户端的总请求量
	req float64
	//客户端成功请求量
	accept float64
}

type SlideValOptions func(val *slideVal)

func NewSlideval(options ...SlideValOptions) *slideVal {
	t := &slideVal{
		time: time.Now().UnixNano(),
	}
	for _, option := range options {
		option(t)
	}
	return t
}

func WithReqOption(req float64) SlideValOptions {
	return func(val *slideVal) {
		val.req = req
	}
}

func WithAcceptReqOption(accept float64) SlideValOptions {
	return func(val *slideVal) {
		val.accept = accept
	}
}

func NewGoogleSlide(interval time.Duration, k float64) *googleSlide {
	return &googleSlide{
		sreSlide:    list.New(),
		interval:    interval.Nanoseconds(),
		k:           k,
		totalReq:    0,
		totalAccept: 0,
	}
}

func (g *googleSlide) Sre() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		g.mutex.Lock()
		now := time.Now().UnixNano()
		front := g.sreSlide.Front()
		//调整滑动窗口
		for front != nil && front.Value.(*slideVal).time+g.interval < now {
			val := front.Value.(*slideVal)
			// 从缓存的统计数据中减去过期数据
			g.totalReq -= val.req
			g.totalAccept -= val.accept
			g.sreSlide.Remove(front)
			front = g.sreSlide.Front()
		}

		//客户端请求被拒绝的概率（(requests − K×accepts) / (requests + 1)）
		tail := (g.totalReq - g.k*g.totalAccept) / (g.totalReq + 1)

		// 提前创建新的请求记录
		newReq := NewSlideval(WithReqOption(1))
		g.sreSlide.PushBack(newReq)
		g.totalReq += 1

		// 在执行实际请求前释放锁
		g.mutex.Unlock()

		if tail > 0 {
			return errors.New("request is fail")
		}

		err := invoker(ctx, method, req, reply, cc, opts...)

		if err != nil {
			sta, ok := status.FromError(err)
			if !ok {
				return err
			}
			if sta.Code() == codes.ResourceExhausted && sta.Code() == codes.Unavailable && sta.Code() == codes.Unknown {
				return err
			}

		}
		// 其他情况视为成功
		g.mutex.Lock()
		newAccept := NewSlideval(WithAcceptReqOption(1))
		g.sreSlide.PushBack(newAccept)
		g.totalAccept += 1
		g.mutex.Unlock()

		return err
	}
}

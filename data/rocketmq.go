package data

import (
	"time"

	rocketmq "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/apache/rocketmq-clients/golang/v5/credentials"
)

type RocketMQConfig struct {
	Endpoint string `yaml:"endpoint"`
	Group    string `yaml:"group"`
}

func NewRocketMQProducer(config *RocketMQConfig) rocketmq.Producer {
	//os.Setenv("mq.consoleAppender.enabled", "true")
	//rocketmq.ResetLogger()
	producer, err := rocketmq.NewProducer(
		&rocketmq.Config{
			Endpoint: config.Endpoint,
			Credentials: &credentials.SessionCredentials{
				AccessKey:    "xxx",
				AccessSecret: "xxx",
			},
		},
		rocketmq.WithTopics("im"),
	)
	if err != nil {
		panic(err)
	}

	producer.Start()

	return producer
}

func NewRocketMQConsumer(config *RocketMQConfig) rocketmq.SimpleConsumer {
	//os.Setenv("mq.consoleAppender.enabled", "true")
	//rocketmq.ResetLogger()
	consumer, err := rocketmq.NewSimpleConsumer(
		&rocketmq.Config{
			Endpoint: config.Endpoint,
			Credentials: &credentials.SessionCredentials{
				AccessKey:    "xxx",
				AccessSecret: "xxx",
			},
			ConsumerGroup: config.Group,
		},
		rocketmq.WithAwaitDuration(time.Second*10),
		rocketmq.WithSubscriptionExpressions(map[string]*rocketmq.FilterExpression{
			"im": rocketmq.SUB_ALL,
		}),
	)
	if err != nil {
		panic(err)
	}

	err = consumer.Start()
	if err != nil {
		panic(err)
	}

	return consumer
}

package internal

import (
	"fmt"
	"os"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	_logger *zap.Logger
	_once   sync.Once
)

func NewLogger() *zap.Logger {
	_once.Do(func() {
		file, err := os.Create(fmt.Sprintf("../../auth-%s.log", time.Now().Format("2006-01-02")))
		if err != nil {
			panic(err)
		}

		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(file)),
			zap.NewAtomicLevelAt(zap.DebugLevel),
		)

		_logger = zap.New(core)
	})

	return _logger
}

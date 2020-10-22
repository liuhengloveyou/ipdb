package common

import (
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger(logDir, logLevel string) error {
	writer, err := rotatelogs.New(
		logDir+"log.%Y%m%d%H%M",
		rotatelogs.WithLinkName("log"),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)
	if err != nil {
		return err
	}

	level := zapcore.InfoLevel
	if logLevel != "" {
		if err := level.UnmarshalText([]byte(logLevel)); err != nil {
			return err
		}

	}

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
		zapcore.AddSync(writer),
		level)

	Logger = zap.New(core, zap.AddCaller())

	return nil
}

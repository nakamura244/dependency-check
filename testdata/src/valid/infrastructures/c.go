package infrastructures

import (
	"os"

	"github.com/nakamura244/dependency-check/testdata/src/valid/config"

	"github.com/nakamura244/dependency-check/testdata/src/valid/domain"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is struct
type Logger struct {
	Logger *zap.Logger
}

// NewLogger is constructor
func NewLogger(cfg *config.Config) *Logger {
	encoderCfg := zapcore.EncoderConfig{}
	encoder := zapcore.NewJSONEncoder(encoderCfg)
	core := zapcore.NewCore(encoder, os.Stdout, nil)
	opts := []zap.Option{zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel)}
	return &Logger{Logger: zap.New(core, opts...)}
}

// Info is logging info
func (l *Logger) Info(msg string, fields ...domain.B) {
}

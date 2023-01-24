package pkg

import (
	"context"
	"github.com/newrelic/go-agent/v3/integrations/logcontext-v2/nrlogrus"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
)

const (
	LoggerContextKey = "loggerTransaction"
)

type Logger struct {
	entry *logrus.Entry
}

type ILogger interface {
	Info(...interface{})
}

func NewLoggerWithContext(ctx context.Context, app *newrelic.Application) *Logger {
	nrlogrusFormatter := nrlogrus.NewFormatter(app, &logrus.TextFormatter{})
	logger := logrus.New()
	logger.SetFormatter(nrlogrusFormatter)

	entry := logger.WithContext(ctx)
	return &Logger{entry: entry}
}

func (l Logger) Info(args ...interface{}) {
	l.entry.Info(args...)
}

func GetLogger(ctx context.Context) ILogger {
	return ctx.Value(LoggerContextKey).(*Logger)
}

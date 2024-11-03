package httpclient

import (
	"fmt"
	"go.uber.org/zap"
)

type restyLogger struct {
	l *zap.Logger
}

func (r restyLogger) Errorf(format string, v ...interface{}) {
	r.l.Error(fmt.Sprintf(format, v...))
}

func (r restyLogger) Warnf(format string, v ...interface{}) {
	r.l.Warn(fmt.Sprintf(format, v...))
}

func (r restyLogger) Debugf(format string, v ...interface{}) {
	r.l.Info(fmt.Sprintf(format, v...))
}

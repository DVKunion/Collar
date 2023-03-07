package log

import (
	"context"

	"github.com/gogf/gf/v2/os/glog"
)

// Log
// a set of chain/setter/logger
type Log interface {
	Logger
	Chain
	Setter
}

type defaultLog struct {
	*glog.Logger
}

func New() Log {
	logger := glog.New()
	return &defaultLog{Logger: logger}
}

var DefaultLog = New()

func Clone() Log {
	return DefaultLog.Clone()
}

func Skip(skip int) Log {
	return DefaultLog.Skip(skip)
}

func Line(long ...bool) Log {
	return DefaultLog.Line(long...)
}

func Prefix(prefix string) Log {
	return DefaultLog.Prefix(prefix)
}

func Cat(category interface{}) Log {
	return DefaultLog.Cat(category)
}

func Info(ctx context.Context, v ...interface{}) {
	DefaultLog.Info(ctx, v...)
}

func Infof(ctx context.Context, format string, v ...interface{}) {
	DefaultLog.Infof(ctx, format, v...)
}

func Debug(ctx context.Context, v ...interface{}) {
	DefaultLog.Debug(ctx, v...)
}

func Debugf(ctx context.Context, format string, v ...interface{}) {
	DefaultLog.Debugf(ctx, format, v...)
}

func Notice(ctx context.Context, v ...interface{}) {
	DefaultLog.Notice(ctx, v...)
}

func Noticef(ctx context.Context, format string, v ...interface{}) {
	DefaultLog.Noticef(ctx, format, v...)
}

func Warning(ctx context.Context, v ...interface{}) {
	DefaultLog.Warning(ctx, v...)
}

func Warningf(ctx context.Context, format string, v ...interface{}) {
	DefaultLog.Warningf(ctx, format, v...)
}

func Error(ctx context.Context, v ...interface{}) {
	DefaultLog.Error(ctx, v...)
}

func Errorf(ctx context.Context, format string, v ...interface{}) {
	DefaultLog.Errorf(ctx, format, v...)
}

func Critical(ctx context.Context, v ...interface{}) {
	DefaultLog.Critical(ctx, v...)
}

func Criticalf(ctx context.Context, format string, v ...interface{}) {
	DefaultLog.Criticalf(ctx, format, v...)
}

func Panic(ctx context.Context, v ...interface{}) {
	DefaultLog.Panic(ctx, v...)
}

func Panicf(ctx context.Context, format string, v ...interface{}) {
	DefaultLog.Panicf(ctx, format, v...)
}

func Fatal(ctx context.Context, v ...interface{}) {
	DefaultLog.Fatal(ctx, v...)
}

func Fatalf(ctx context.Context, format string, v ...interface{}) {
	DefaultLog.Fatalf(ctx, format, v...)
}

func (l *defaultLog) Info(ctx context.Context, v ...interface{}) {
	l.Logger.Info(ctx, v...)
}

func (l *defaultLog) Infof(ctx context.Context, format string, v ...interface{}) {
	l.Logger.Infof(ctx, format, v...)
}

func (l *defaultLog) Debug(ctx context.Context, v ...interface{}) {
	l.Logger.Debug(ctx, v...)
}

func (l *defaultLog) Debugf(ctx context.Context, format string, v ...interface{}) {
	l.Logger.Debugf(ctx, format, v...)
}

func (l *defaultLog) Notice(ctx context.Context, v ...interface{}) {
	l.Logger.Notice(ctx, v...)
}

func (l *defaultLog) Noticef(ctx context.Context, format string, v ...interface{}) {
	l.Logger.Noticef(ctx, format, v...)
}

func (l *defaultLog) Warning(ctx context.Context, v ...interface{}) {
	l.Logger.Warning(ctx, v...)
}

func (l *defaultLog) Warningf(ctx context.Context, format string, v ...interface{}) {
	l.Logger.Warningf(ctx, format, v...)
}

func (l *defaultLog) Error(ctx context.Context, v ...interface{}) {
	l.Logger.Error(ctx, v...)
}

func (l *defaultLog) Errorf(ctx context.Context, format string, v ...interface{}) {
	l.Logger.Errorf(ctx, format, v...)
}

func (l *defaultLog) Critical(ctx context.Context, v ...interface{}) {
	l.Logger.Critical(ctx, v...)
}

func (l *defaultLog) Criticalf(ctx context.Context, format string, v ...interface{}) {
	l.Logger.Criticalf(ctx, format, v...)
}

func (l *defaultLog) Panic(ctx context.Context, v ...interface{}) {
	l.Logger.Panic(ctx, v...)
}

func (l *defaultLog) Panicf(ctx context.Context, format string, v ...interface{}) {
	l.Logger.Panicf(ctx, format, v...)
}

func (l *defaultLog) Fatal(ctx context.Context, v ...interface{}) {
	l.Logger.Fatal(ctx, v...)
}

func (l *defaultLog) Fatalf(ctx context.Context, format string, v ...interface{}) {
	l.Logger.Fatalf(ctx, format, v...)
}

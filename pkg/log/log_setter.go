package log

import (
	"context"
	"time"
)

type Handler func(ctx context.Context)

// Setter
// set logger config
type Setter interface {
	SetPath(path string) error
	SetFile(path string)
	SetLevel(level Level)
	SetPrefix(prefix string)
	SetHeaderPrint(enabled bool)
	SetStdoutPrint(enabled bool)

	SetRotateSize(size int64)
	SetRotateExpire(t time.Duration)
	SetRotateBackupLimit(limit int)
	SetRotateBackupExpire(t time.Duration)
	SetRotateBackupCompress(n int)
}

func (l *defaultLog) SetPath(path string) error {
	return l.Logger.SetPath(path)
}

func (l *defaultLog) SetFile(pattern string) {
	l.Logger.SetFile(pattern)
}

func (l *defaultLog) SetLevel(level Level) {
	l.Logger.SetLevel(int(level))
}

func (l *defaultLog) SetPrefix(prefix string) {
	l.Logger.SetPrefix(prefix)
}

func (l *defaultLog) SetHeaderPrint(enabled bool) {
	l.Logger.SetHeaderPrint(enabled)
}

func (l *defaultLog) SetStdoutPrint(enabled bool) {
	l.Logger.SetStdoutPrint(enabled)
}

func (l *defaultLog) SetRotateSize(size int64) {
	config := l.Logger.GetConfig()
	config.RotateSize = size
	l.Logger.SetConfig(config)
}

func (l *defaultLog) SetRotateExpire(t time.Duration) {
	config := l.Logger.GetConfig()
	config.RotateExpire = t
	l.Logger.SetConfig(config)
}

func (l *defaultLog) SetRotateBackupLimit(limit int) {
	config := l.Logger.GetConfig()
	config.RotateBackupLimit = limit
	l.Logger.SetConfig(config)
}

func (l *defaultLog) SetRotateBackupExpire(t time.Duration) {
	config := l.Logger.GetConfig()
	config.RotateBackupExpire = t
	l.Logger.SetConfig(config)
}

func (l *defaultLog) SetRotateBackupCompress(n int) {
	config := l.Logger.GetConfig()
	config.RotateBackupCompress = n
	l.Logger.SetConfig(config)
}

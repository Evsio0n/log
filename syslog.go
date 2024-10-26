//go:build !windows
// +build !windows

package log

import (
	"log/syslog"
)

// SyslogWriter 实现 LogWriter 接口，输出到 syslog
type SyslogWriter struct {
	writer *syslog.Writer
}

func NewSyslogWriter(tag string) (*SyslogWriter, error) {
	w, err := syslog.Dial("", "", syslog.LOG_DAEMON, tag)
	if err != nil {
		return nil, err
	}
	return &SyslogWriter{writer: w}, nil
}

func (s *SyslogWriter) Write(level LogLevel, message string) {
	switch level {
	case LevelDebug:
		s.writer.Debug(message)
	case LevelInfo:
		s.writer.Info(message)
	case LevelWarn:
		s.writer.Warning(message)
	case LevelError:
		s.writer.Err(message)
	case LevelFatal, LevelPanic:
		s.writer.Crit(message)
	}
}

func (s *SyslogWriter) Close() error {
	return s.writer.Close()
}

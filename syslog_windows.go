//go:build windows
// +build windows

package log

import (
	"fmt"
)

// NewSyslogWriter 返回 ConsoleWriter，并提示 syslog 不可用
func NewSyslogWriter(tag string) (LogWriter, error) {
	fmt.Println("Syslog is not supported on Windows, using ConsoleWriter instead")
	return &ConsoleWriter{}, nil
}

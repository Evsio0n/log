package log

import (
	"fmt"
)

// ConsoleWriter 实现 LogWriter 接口，输出到控制台
type ConsoleWriter struct{}

func (c *ConsoleWriter) Write(level LogLevel, message string) {
	colorCode := ""
	switch level {
	case LevelDebug:
		colorCode = "\033[1;36m" // 青色
	case LevelInfo:
		colorCode = "\033[1;32m" // 绿色
	case LevelWarn:
		colorCode = "\033[1;33m" // 黄色
	case LevelError, LevelFatal, LevelPanic:
		colorCode = "\033[1;31m" // 红色
	}
	fmt.Printf("%s%s\033[0m\n", colorCode, message)
}

func (c *ConsoleWriter) Close() error {
	return nil
}

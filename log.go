package log

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

// LogLevel 定义日志级别
type LogLevel int

const (
	LevelDebug LogLevel = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

// LogWriter 定义日志输出接口
type LogWriter interface {
	Write(level LogLevel, message string)
	Close() error
}

// Logger 定义日志器结构体
type Logger struct {
	mu              sync.Mutex
	isDebug         bool
	showDate        bool
	showLogCategory bool
	logWriter       LogWriter
	dateFormat      string
	callDepth       int
}

// NewLogger 创建一个新的日志器
func NewLogger() *Logger {
	return &Logger{
		isDebug:         false,
		showDate:        false,
		showLogCategory: false,
		logWriter:       &ConsoleWriter{},
		dateFormat:      "2006-01-02 15:04:05.000000",
		callDepth:       3,
	}
}

// SetDebug 设置是否启用调试模式
func (l *Logger) SetDebug(debug bool) {
	l.isDebug = debug
}

// SetShowDate 设置是否显示日期
func (l *Logger) SetShowDate(show bool) {
	l.showDate = show
}

// SetShowLogCategory 设置是否显示日志类别
func (l *Logger) SetShowLogCategory(show bool) {
	l.showLogCategory = show
}

// SetDateFormat 设置日期格式
func (l *Logger) SetDateFormat(format string) {
	l.dateFormat = format
}

// SetLogWriter 设置日志输出
func (l *Logger) SetLogWriter(writer LogWriter) {
	l.logWriter = writer
}

// 通用的日志记录函数
func (l *Logger) log(level LogLevel, levelStr string, args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	var sb strings.Builder

	if l.showDate {
		sb.WriteString(time.Now().Format(l.dateFormat))
		sb.WriteString(" ")
	}

	if l.showLogCategory && levelStr != "" {
		sb.WriteString("[")
		sb.WriteString(levelStr)
		sb.WriteString("] ")
	}

	sb.WriteString(l.getPosition())
	sb.WriteString(" -> ")
	sb.WriteString(fmt.Sprint(args...))

	message := sb.String()

	l.logWriter.Write(level, message)

	switch level {
	case LevelFatal:
		panic("Fatal error occurred")
	case LevelPanic:
		panic(message)
	}
}

// Debug 级别日志
func (l *Logger) Debug(args ...interface{}) {
	if l.isDebug {
		l.log(LevelDebug, "DEBUG", args...)
	}
}

// Info 级别日志
func (l *Logger) Info(args ...interface{}) {
	l.log(LevelInfo, "INFO", args...)
}

// Warn 级别日志
func (l *Logger) Warn(args ...interface{}) {
	l.log(LevelWarn, "WARN", args...)
}

// Error 级别日志
func (l *Logger) Error(args ...interface{}) {
	l.log(LevelError, "ERROR", args...)
}

// Fatal 级别日志
func (l *Logger) Fatal(args ...interface{}) {
	l.log(LevelFatal, "FATAL", args...)
}

// Panic 级别日志
func (l *Logger) Panic(args ...interface{}) {
	l.log(LevelPanic, "PANIC", args...)
}

// 获取调用位置
func (l *Logger) getPosition() string {
	_, file, line, ok := runtime.Caller(l.callDepth)
	if !ok {
		file = "???"
		line = 0
	}
	fileName := filepath.Base(file)
	return fmt.Sprintf("%s:%d", fileName, line)
}

// SetSyslog 设置使用 SyslogWriter
func (l *Logger) SetSyslog(tag string) error {
	syslogWriter, err := NewSyslogWriter(tag)
	if err != nil {
		// 在这里处理错误，例如切换回控制台输出
		fmt.Println("Syslog is not available, using ConsoleWriter")
		return err
	}
	l.SetLogWriter(syslogWriter)
	return nil
}

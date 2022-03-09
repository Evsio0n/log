package log

import (
	"fmt"
	"log/syslog"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var isDEBUG = false
var isShowDate = false
var isSyslog = false
var syslogTag = "syslog"
var dial *syslog.Writer
var err error
var isShowLogCatagory = false

func SetSyslog(debug bool, tagName string) error {
	isSyslog = debug
	syslogTag = tagName
	dial, err = syslog.Dial("", "", syslog.LOG_DAEMON, syslogTag)
	if err != nil {
		return err
	}
	return nil
}

func SetDebug(debug bool) {
	isDEBUG = debug
}

func IsShowDate(ShowDate bool) {
	isShowDate = ShowDate
}

func IsShowLogCatagory(ShowLogCatagory bool) {
	isShowLogCatagory = ShowLogCatagory
}

func Info(arg ...interface{}) {
	var s string
	if isShowLogCatagory {
		s = buildString(" [INFO]  ", arg)
	} else {
		s = buildString("", arg)
	}
	if isSyslog {
		dial.Info(s)
	} else {
		fmt.Printf("%c[1;32m%s%c[0m\n", 0x1B, s, 0x1B)
	}

}

func Debug(arg ...interface{}) {
	if isDEBUG {
		var s string
		if isShowLogCatagory {
			s = buildString(" [DEBUG] ", arg)
		} else {
			s = buildString("", arg)
		}
		if isSyslog {
			dial.Debug(s)
		} else {
			fmt.Printf("%c[1;36m%s%c[0m\n", 0x1B, s, 0x1B)
		}
	}
}

func Warn(arg ...interface{}) {
	var s string
	if isShowLogCatagory {
		s = buildString(" [WARN]  ", arg)
	} else {
		s = buildString("", arg)
	}
	if isSyslog {
		dial.Warning(s)
	} else {
		fmt.Printf("%c[1;33m%s%c[0m\n", 0x1B, s, 0x1B)
	}
}

func Error(arg ...interface{}) {
	var s string
	if isShowLogCatagory {
		s = buildString(" [ERROR] ", arg)
	} else {
		s = buildString("", arg)
	}
	if isSyslog {
		dial.Err(s)
	} else {
		fmt.Printf("%c[1;31m%s%c[0m\n", 0x1B, s, 0x1B)
	}
}

func Fatal(arg ...interface{}) {
	var s string
	if isShowLogCatagory {
		s = buildString(" [FATAL] ", arg)
	} else {
		s = buildString("", arg)
	}
	if isSyslog {
		dial.Crit(s)
	} else {
		fmt.Printf("%c[1;31m%s%c[0m\n", 0x1B, s, 0x1B)
	}
}

func Panic(arg ...interface{}) {
	var s string

	switch isShowLogCatagory {
	case true:
		s = buildString(" [PANIC] ", arg)
	default:
		s = buildString("", arg)
	}

	switch isSyslog {
	case true:
		dial.Crit(s)
	default:
		fmt.Printf("%c[1;31m%s%c[0m\n", 0x1B, s, 0x1B)
	}
	panic(s)
}

func buildString(level string, args []interface{}) string {
	var tag []interface{}
	switch isShowDate {
	case true:
		tag = append(tag, time.Now().Format("2006-01-02 15:04:05.000000"), level, getPosition(), " -> ")
	case false:
		tag = append(tag, level, getPosition(), " -> ")
	}
	s := fmt.Sprint(tag...) + fmt.Sprint(args...)
	return s
}

func getPosition() string {
	var file string
	var line int
	var ok bool
	_, file, line, ok = runtime.Caller(3)
	if !ok {
		file = "???"
		line = 0
	}
	path := strings.Split(file, "/")
	index := len(path) - 1
	return path[index] + ":" + strconv.Itoa(line)
}

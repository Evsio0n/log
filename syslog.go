//go:build darwin || linux || android
// +build darwin linux android

package log

import (
	"fmt"
	"log/syslog"
	_ "time"
)

var isSyslog = false
var syslogTag = "syslog"
var dial *syslog.Writer

func SetSyslog(debug bool, tagName string) error {
	isSyslog = debug
	syslogTag = tagName
	dial, err = syslog.Dial("", "", syslog.LOG_DAEMON, syslogTag)
	if err != nil {
		return err
	}
	return nil
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

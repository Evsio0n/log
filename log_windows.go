//go:build windows
// +build windows

package log

var isSyslog = false

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
	fmt.Printf("%c[1;32m%s%c[0m\n", 0x1B, s, 0x1B)

}

func Debug(arg ...interface{}) {
	if isDEBUG {
		var s string
		if isShowLogCatagory {
			s = buildString(" [DEBUG] ", arg)
		} else {
			s = buildString("", arg)
		}
		fmt.Printf("%c[1;36m%s%c[0m\n", 0x1B, s, 0x1B)

	}
}

func Warn(arg ...interface{}) {
	var s string
	if isShowLogCatagory {
		s = buildString(" [WARN]  ", arg)
	} else {
		s = buildString("", arg)
	}
	fmt.Printf("%c[1;33m%s%c[0m\n", 0x1B, s, 0x1B)

}

func Error(arg ...interface{}) {
	var s string
	if isShowLogCatagory {
		s = buildString(" [ERROR] ", arg)
	} else {
		s = buildString("", arg)
	}
	fmt.Printf("%c[1;31m%s%c[0m\n", 0x1B, s, 0x1B)

}

func Fatal(arg ...interface{}) {
	var s string
	if isShowLogCatagory {
		s = buildString(" [FATAL] ", arg)
	} else {
		s = buildString("", arg)
	}
	fmt.Printf("%c[1;31m%s%c[0m\n", 0x1B, s, 0x1B)

}

func Panic(arg ...interface{}) {
	var s string

	switch isShowLogCatagory {
	case true:
		s = buildString(" [PANIC] ", arg)
	default:
		s = buildString("", arg)
	}

	fmt.Printf("%c[1;31m%s%c[0m\n", 0x1B, s, 0x1B)

	panic(s)
}

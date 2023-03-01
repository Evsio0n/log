//go:build darwin || linux || android
// +build darwin linux android

package log

import "log/syslog"

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

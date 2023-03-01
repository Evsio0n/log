package log

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var isDEBUG = false
var isShowDate = false

// only for syslog available system

var err error
var isShowLogCatagory = false

func SetDebug(debug bool) {
	isDEBUG = debug
}

func IsShowDate(ShowDate bool) {
	isShowDate = ShowDate
}

func IsShowLogCatagory(ShowLogCatagory bool) {
	isShowLogCatagory = ShowLogCatagory
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

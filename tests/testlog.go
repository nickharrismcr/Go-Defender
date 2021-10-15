package tests

import (
	"regexp"
)

var loglist []string

func InitTestLog() {
	loglist = []string{}
}

func AddTestLog(s string) {
	loglist = append(loglist, s)

}

func GetTestLog() []string {
	return loglist
}

func TestLogCountPattern(patt string) (c int) {
	c = 0
	valid := regexp.MustCompile(patt)
	for _, s := range loglist {
		if valid.MatchString(s) {
			c++
		}
	}
	return
}

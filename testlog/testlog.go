package testlog

import "regexp"

var loglist []string

func Init() {
	loglist = []string{}
}

func Add(s string) {
	loglist = append(loglist, s)
}

func Get() []string {
	return loglist
}

func CountPattern(patt string) (c int) {
	c = 0
	valid := regexp.MustCompile(patt)
	for _, s := range loglist {
		if valid.MatchString(s) {
			c++
		}
	}
	return
}

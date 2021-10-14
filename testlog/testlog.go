package testlog

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

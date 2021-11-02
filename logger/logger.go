package logger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()
var f os.File

func init() {

	logger.SetLevel(logrus.PanicLevel)
	f, _ := os.OpenFile("./defender.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	logger.Out = f
}

func Debug(format string, args ...interface{}) {
	logger.Debug(fmt.Sprintf(format, args...))
}

func Info(format string, args ...interface{}) {
	logger.Info(fmt.Sprintf(format, args...))
}

func Close() {
	f.Close()
}

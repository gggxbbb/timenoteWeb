package log

import (
	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"os"
)

var Logger *logrus.Logger

func init() {
	// setup logger
	Logger = logrus.New()
	Logger.SetLevel(logrus.DebugLevel)
	Logger.SetFormatter(&formatter.Formatter{
		HideKeys:        false,
		TimestampFormat: "[2006-01-02 15:04:05]",
	})
	Logger.Out = os.Stdout
}

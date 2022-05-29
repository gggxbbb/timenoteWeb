package log

import (
	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"os"
)

// Logger 公共 logger
var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	Logger.SetLevel(logrus.DebugLevel)
	Logger.SetFormatter(&formatter.Formatter{
		HideKeys:        false,
		TimestampFormat: "[2006-01-02 15:04:05]",
	})
	Logger.Out = os.Stdout
}

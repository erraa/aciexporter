package acilogger

import "github.com/erraa/aciexporter/config"
import "github.com/sirupsen/logrus"

var Log = logrus.New()
var Fields = logrus.Fields{}

func init() {
	Log = logrus.New()
	var loglevel = config.ConfigInstance.Loglevel
	switch loglevel {
	case "debug":
		Log.SetLevel(logrus.DebugLevel)
	case "info":
		Log.SetLevel(logrus.InfoLevel)
	default:
		Log.SetLevel(logrus.DebugLevel)
	}
}

package configs

import "github.com/sirupsen/logrus"

// Set custom config options for logrus
func SetUpLogging() {
	customFormatter := new(logrus.TextFormatter)
	customFormatter.DisableLevelTruncation = true
	customFormatter.FullTimestamp = true
	logrus.SetFormatter(customFormatter)
}

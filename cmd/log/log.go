package log

import (
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	log     *logrus.Logger
	logOnce sync.Once
)

func Provide() *logrus.Logger {
	logOnce.Do(func() {

		logLevel, ok := os.LookupEnv("LOG_LEVEL")

		// LOG_LEVEL not set, let's default to debug
		if !ok {
			logLevel = "debug"
		}

		// parse string, this is built-in feature of logrus
		logrusLevel, err := logrus.ParseLevel(logLevel)
		if err != nil {
			logrusLevel = logrus.DebugLevel
		}

		// set global log level
		logrus.SetLevel(logrusLevel)

		// TimestampFormat: "2006-01-02 150405"
		logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

		logrus.SetOutput(os.Stdout)

		logrus.Infoln("[LOGRUS-CONFIG] Log level: ", logrus.GetLevel())
		log = logrus.New()
	})
	return log
}

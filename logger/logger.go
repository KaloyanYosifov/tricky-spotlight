package logger

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func Init(env string) {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)

	if env == "debug" {
		log.SetLevel(log.TraceLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}

func Trace(args ...interface{}) {
	log.Trace(args)
}

func Debug(args ...interface{}) {
	log.Debug(args)
}

func Info(args ...interface{}) {
	log.Info(args)
}

func Warn(args ...interface{}) {
	log.Warn(args)
}

func Error(args ...interface{}) {
	log.Error(args)
}

func Fatal(args ...interface{}) {
	log.Fatal(args)
}

func Panic(args ...interface{}) {
	log.Panic(args)
}

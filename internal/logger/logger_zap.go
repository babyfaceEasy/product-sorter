package logger

import (
	"go.uber.org/zap"
)

var log *zap.Logger

func Init(isProduction bool) {
	var err error
	if isProduction {
		log, err = zap.NewProduction()
	} else {
		log, err = zap.NewDevelopment()
	}
	if err != nil {
		panic(err)
	}
}

func L() *zap.Logger {
	if log == nil {
		Init(true) // default to production
	}
	return log
}

package logging

import (
	"fmt"
	"go.uber.org/zap"
)

func NewZapSugarLogger(logName string) *zap.SugaredLogger {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"stdout",
		"./var/log/" + logName + ".log",
	}

	b, err := cfg.Build()
	if err != nil {
		fmt.Println(err.Error())
	}

	defer b.Sync()

	sugar := b.Sugar()
	return sugar
}

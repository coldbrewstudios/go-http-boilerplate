package logging

import (
	"fmt"
	"go.uber.org/zap"
)

func NewZapSugarLogger(env string) *zap.SugaredLogger {
	fmt.Println(env)
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"stdout",
		"./var/log/file.log",
	}

	b, err := cfg.Build()
	if err != nil {
		fmt.Println(err.Error())
	}

	defer b.Sync()
	sugar := b.Sugar()
	return sugar
}

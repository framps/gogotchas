package main

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func hookD(e zapcore.Entry) error {
	fmt.Print("*** DDD Hook called ***\n")
	return nil
}

func hookP(e zapcore.Entry) error {
	fmt.Print("*** PPP Hook called ***\n")
	return nil
}

func main() {

	zap.L().Info("ZAP")

	loggerd, err := zap.NewDevelopment(
		zap.Fields(zap.String("name", "Dev")),
		zap.Hooks(hookD))
	if err != nil {
		panic(err)
	}

	loggerd.Info("Starting DDD")

	loggerp, err := zap.NewProduction(
		zap.Fields(zap.String("name", "Prod")),
		zap.Hooks(hookP))
	if err != nil {
		panic(err)
	}

	loggerp.Info("Starting PPP")

}

package main

import (
	"github.com/vietbui1502/RestAPIGolang/app"
	"github.com/vietbui1502/RestAPIGolang/logger"
)

func main() {
	logger.Debug("Application starting ...")
	app.Start()
}

package main

import (
	"github.com/aticio/banking/app"
	"github.com/aticio/banking/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}

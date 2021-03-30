package main

import "go.uber.org/zap"

var logger, loggerErr = zap.NewProduction()

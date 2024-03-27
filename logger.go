package main

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/tymcgee/go-starter/config"
	"gopkg.in/natefinch/lumberjack.v2"
)

func setupLogger() {
	file := &lumberjack.Logger{
		Filename:   config.Config.LogFile,
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28,    // days
		Compress:   false, // disabled by default
	}

	// show the pretty logs to stdout, only show the structured logs to the log file
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	multi := zerolog.MultiLevelWriter(consoleWriter, file)
	log.Logger = log.Output(multi)

	if config.Config.Environment == config.Local || config.Config.Environment == config.Dev {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}

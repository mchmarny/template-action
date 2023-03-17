package main

import (
	"os"

	action "github.com/mchmarny/action/cmd/action/cli"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	logLevelEnvVar = "debug"
)

var (
	version = "v0.0.1-default"
)

func main() {
	initLogging()
	err := action.Execute(version, os.Args)
	if err != nil {
		log.Error().Msg(err.Error())
	}
}

func initLogging() {
	level := zerolog.InfoLevel
	levStr := os.Getenv(logLevelEnvVar)
	if levStr == "true" {
		level = zerolog.DebugLevel
	}

	zerolog.SetGlobalLevel(level)

	out := zerolog.ConsoleWriter{
		Out: os.Stderr,
		PartsExclude: []string{
			zerolog.TimestampFieldName,
		},
	}

	log.Logger = zerolog.New(out)
}

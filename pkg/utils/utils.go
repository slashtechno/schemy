package utils

import (
	"strings"

	"github.com/charmbracelet/log"
)

func SetupLogger(logLevel string) {
	switch strings.ToLower(logLevel) {
	case "debug":
		// log.Info("Setting log level to debug")
		log.SetReportCaller(true)
		log.SetLevel(log.DebugLevel)
		log.Debug("Set log level to debug and enabled report caller")
	case "info":
		// log.Info("Setting log level to info")
		log.SetLevel(log.InfoLevel)
		// No point in logging an Info message for WarnLevel and ErrorLevel as the user probably wants to see warnings and errors only
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.InfoLevel)
		log.Info("Invalid log level passed, using InfoLevel", "passed", logLevel)
	}
}

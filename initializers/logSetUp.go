package initializers

import (
	"git.recombee.net/jtrojak/server-db/config"
	log "github.com/sirupsen/logrus"
)

func init() {
	cfg := config.LoggerConfig{
		LogFormat: "json",
		LogLevel: "info",
	}
	LogSetUp(cfg)
}

func LogSetUp(logCfg config.LoggerConfig) {
	
	if logCfg.LogFormat == "json" {
		log.SetFormatter(&log.JSONFormatter{})
	} else {
		log.SetFormatter(&log.TextFormatter{})
	}

	level, err := log.ParseLevel(logCfg.LogLevel)
	if err != nil {
		log.SetLevel(log.InfoLevel)
		return
	}
	log.SetLevel(level)
}

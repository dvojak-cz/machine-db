package main

import (
	"git.recombee.net/jtrojak/server-db/config"
	"git.recombee.net/jtrojak/server-db/constants"
	"git.recombee.net/jtrojak/server-db/initializers"
	"git.recombee.net/jtrojak/server-db/repository"
	httpEngine "git.recombee.net/jtrojak/server-db/router/http"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func init() {
	initializers.LoadEnvVariables()
	config.LoadFromEnv()
	initializers.LogSetUp(config.Configs.LoggerConfig)
	repository.Init()
}

func main() {
	log.Infof("Application %s, version %s\n", constants.AppName, constants.Version)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go httpEngine.Run(config.Configs.WebServerConfig.Http)
	go signalsHandler()
	wg.Wait()

}

func signalsHandler() {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	select {
	case <-sigc:
		// wg.Done()
		// TODO...
		os.Exit(1)
	}
}

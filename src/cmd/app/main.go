package main

import (
	"os"
	"os/signal"
	"syscall"

	"projecto/app"
	"projecto/config"
	"projecto/db"
	"projecto/service/item"
	"projecto/service/item/itemrepo"
	"projecto/web"

	log "github.com/sirupsen/logrus"
)

func main() {
	logger := log.Fields{"app": true}
	app := app.App{Logger: log.WithFields(logger)}
	initialize(&app)
	if err := app.Start(); err != nil {
		log.WithFields(logger).Fatalf("app %v %v cant start error: %v", app.Name(), app.Version(), err)
	}
	log.WithFields(logger).Infof("app %v %v started", app.Name(), app.Version())
	shutdownsignals := make(chan os.Signal, 1)
	signal.Notify(shutdownsignals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	sig := <-shutdownsignals
	log.WithFields(logger).Errorf("app %v %v received signal: %v", app.Name(), app.Version(), sig)
	if err := app.Close(); err != nil {
		log.WithFields(logger).Errorf("cant stop app: %v", err)
	}
}

func initialize(a *app.App) {
	a.Register(config.New()).
		Register(web.New()).
		Register(db.New()).
		Register(itemrepo.New()).
		Register(item.New())
}

package main

import (
	"context"
	"errors"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/himmel520/uoffer/require/internal/config"
	httphandler "github.com/himmel520/uoffer/require/internal/handler/http"
	"github.com/himmel520/uoffer/require/internal/server/cron"

	"github.com/himmel520/uoffer/require/internal/repository/postgres"
	"github.com/himmel520/uoffer/require/internal/repository/redis"
	"github.com/himmel520/uoffer/require/internal/server"
	"github.com/himmel520/uoffer/require/internal/service"
	"github.com/himmel520/uoffer/require/internal/service/parser"
)

// @title API Documentation
// @version 1.0
// @description API для сервиса требований
// @host localhost:8081
// @BasePath /api/v1

func main() {
	logLevel := flag.String("loglevel", "info", "log level: debug, info, warn, error")
	flag.Parse()

	log := server.SetupLogger(*logLevel)

	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.NewPG(cfg.DB.DBConn)
	if err != nil {
		log.Fatalf("unable to connect to pool: %v", err)
	}
	defer db.Close()

	// Слои
	cache, err := redis.New(cfg.Cache)
	if err != nil {
		log.Fatalf("unable to connect to redis: %v", err)
	}

	repo := postgres.NewRepo(db)
	srv := service.New(repo, cache, cfg.Srv.MediaServiceURL, log)
	handler := httphandler.New(srv, &cfg.Srv.JWT, log)

	// сервер
	parser := parser.NewParser(cfg.API_HH, repo, cache, log)
	cron := cron.NewCron(context.Background(), cfg.API_HH.Interval, parser, log)
	app := server.New(handler.InitRoutes(), cron, cfg.Srv.Addr)

	go func() {
		log.Infof("the server is starting on %v", cfg.Srv.Addr)

		if err := app.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Errorf("error occured while running http server: %s", err.Error())
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGTERM, syscall.SIGINT)
	<-done

	if err := app.Shutdown(context.Background()); err != nil {
		log.Errorf("error occured on server shutting down: %s", err)
	}

	log.Info("the server is shut down")
}

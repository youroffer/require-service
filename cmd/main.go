package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/himmel520/uoffer/require/config"
	"github.com/himmel520/uoffer/require/internal/controller/ogen"
	analyticHandler "github.com/himmel520/uoffer/require/internal/controller/ogen/analytic"
	authHandler "github.com/himmel520/uoffer/require/internal/controller/ogen/auth"
	categoryHandler "github.com/himmel520/uoffer/require/internal/controller/ogen/category"
	errHandler "github.com/himmel520/uoffer/require/internal/controller/ogen/error"
	filterHandler "github.com/himmel520/uoffer/require/internal/controller/ogen/filter"
	positionHandler "github.com/himmel520/uoffer/require/internal/controller/ogen/position"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/postgres"

	log "github.com/youroffer/logger"
)

func init() {
	logLevel := flag.String("loglevel", "info", "log level: debug, info, warn, error")
	flag.Parse()

	log.SetupLogger(*logLevel)
}

func main() {
	// config
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	// db
	pool, err := postgres.NewPG(cfg.DB.DBConn)
	if err != nil {
		log.FatalMsg(err, "unable to connect to pool")
	}
	defer pool.Close()
	dbtx := repository.NewDBTX(pool)
	_ = dbtx

	// rdb, err := redis.New(cfg.Cache.Conn)
	// if err != nil {
	// 	log.Fatalf("unable to connect to cache: %v", err)
	// }
	// defer rdb.Close()

	// cache := cache.New(rdb, cfg.Cache.Exp)
	// repo := repository.New(pool)
	// uc := usecase.New(repo, cache, cfg.Srv.JWT.PublicKey, log)
	// handler := httpctrl.New(uc, log)

	// сервер
	// parser := parser.NewParser(cfg.API_HH, repo, cache, log)
	// cron := cron.NewCron(context.Background(), cfg.API_HH.Interval, parser, log)
	// app := server.New(handler.InitRoutes(), cron, cfg.Srv.Addr)

	handler := ogen.NewHandler(ogen.HandlerParams{
		Auth:     authHandler.New(nil),
		Error:    errHandler.New(),
		Analytic: analyticHandler.New(),
		Category: categoryHandler.New(),
		Filter:   filterHandler.New(),
		Position: positionHandler.New(),
	})

	app, err := ogen.NewServer(handler, cfg.Srv.Addr)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		log.Infof("the server is starting on %v", cfg.Srv.Addr)

		if err := app.Run(); err != nil {
			log.FatalMsg(err, "error occured while running http server")
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGTERM, syscall.SIGINT)
	<-done

	if err := app.Shutdown(context.Background()); err != nil {
		log.ErrMsg(err, "error occured on server shutting down")
	}

	log.Info("the server is shut down")
}

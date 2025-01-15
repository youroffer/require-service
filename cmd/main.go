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
	"github.com/himmel520/uoffer/require/internal/infrastructure/cache"
	"github.com/himmel520/uoffer/require/internal/infrastructure/parser"
	"github.com/himmel520/uoffer/require/internal/infrastructure/parser/cron"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/postgres"
	analyticRepo "github.com/himmel520/uoffer/require/internal/infrastructure/repository/postgres/analytic"
	categoryRepo "github.com/himmel520/uoffer/require/internal/infrastructure/repository/postgres/category"
	filterRepo "github.com/himmel520/uoffer/require/internal/infrastructure/repository/postgres/filter"
	positionRepo "github.com/himmel520/uoffer/require/internal/infrastructure/repository/postgres/positions"
	analyticUC "github.com/himmel520/uoffer/require/internal/usecase/analytic"
	categoryUC "github.com/himmel520/uoffer/require/internal/usecase/category"
	filterUC "github.com/himmel520/uoffer/require/internal/usecase/filter"
	positionUC "github.com/himmel520/uoffer/require/internal/usecase/positions"

	"github.com/rs/zerolog/log"
	logSetup "github.com/youroffer/logger"
)

func init() {
	logLevel := flag.String("loglevel", "info", "log level: debug, info, warn, error")
	flag.Parse()

	logSetup.SetupLogger(*logLevel)
}

func main() {
	// config
	cfg, err := config.New()
	if err != nil {
		log.Fatal().Err(err).Msg("invalid env")
	}

	// db
	pool, err := postgres.NewPG(cfg.DB.DBConn)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to connect to pool")
	}
	defer pool.Close()
	dbtx := repository.NewDBTX(pool)

	// cache
	rdb, err := cache.NewRedis(cfg.Cache.Conn)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to connect to cache")
	}
	defer rdb.Close()
	cache := cache.NewCache(rdb, cfg.Cache.Exp)

	// repo
	filterRepo := filterRepo.New()
	categoryRepo := categoryRepo.New()
	analyticRepo := analyticRepo.New()
	positionRepo := positionRepo.New()

	// uc
	filterUC := filterUC.New(dbtx, filterRepo)
	categoryUC := categoryUC.New(dbtx, categoryRepo)
	analyticUC := analyticUC.New(dbtx, analyticRepo, cache)
	positionUC := positionUC.New(dbtx, positionRepo)

	// handler
	handler := ogen.NewHandler(ogen.HandlerParams{
		Auth:     authHandler.New(nil),
		Error:    errHandler.New(),
		Analytic: analyticHandler.New(analyticUC),
		Category: categoryHandler.New(categoryUC),
		Filter:   filterHandler.New(filterUC),
		Position: positionHandler.New(positionUC),
	})

	// parser
	parser := parser.NewParser(parser.ParserParams{
		Cfg:          cfg.API_HH,
		AnalyticRepo: analyticRepo,
		FilterRepo:   filterRepo,
		Cache:        cache,
		DBTX:         dbtx,
	})
	cron := cron.NewCron(context.Background(), cfg.API_HH.Interval, parser)
	go cron.Start()
	defer cron.Stop()

	// server
	app, err := ogen.NewServer(handler, cfg.Srv.Addr)
	if err != nil {
		log.Fatal().Err(err).Msg("create server")
	}

	go func() {
		log.Info().Msgf("the server is starting on %v", cfg.Srv.Addr)

		if err := app.Run(); err != nil {
			log.Fatal().Err(err).Msg("error occured while running http server")
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGTERM, syscall.SIGINT)
	<-done

	if err := app.Shutdown(context.Background()); err != nil {
		log.Err(err).Msg("error occured on server shutting down")
	}
}

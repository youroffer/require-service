package cron

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type Parser interface {
	Parse(ctx context.Context)
}

type Cron struct {
	interval time.Duration
	ctx      context.Context
	cancel   context.CancelFunc
	parser   Parser
	log      *logrus.Logger
}

func NewCron(ctx context.Context, interval time.Duration, parser Parser, log *logrus.Logger) *Cron {
	ctx, cancel := context.WithCancel(ctx)
	return &Cron{
		interval: interval,
		ctx:      ctx,
		cancel:   cancel,
		parser:   parser,
		log:      log,
	}
}

func (c *Cron) Start() {
	c.log.Info("cron is running")
	// c.parser.Parse(c.ctx)
	// ticker := time.NewTicker(c.interval)
	// for {
	// 	select {
	// 	case <-c.ctx.Done():
	// 		ticker.Stop()
	// 		return
	// 	case <-ticker.C:
	// 		c.parser.Parse(c.ctx)
	// 	}
	// }
}

func (c *Cron) Stop() {
	c.log.Info("cron stopped")
	c.cancel()
}

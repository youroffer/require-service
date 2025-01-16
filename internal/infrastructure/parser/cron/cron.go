package cron

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
)

type Parser interface {
	Parse(ctx context.Context)
}

type Cron struct {
	parser   Parser
	interval time.Duration

	ctx    context.Context
	cancel context.CancelFunc
}

func NewCron(ctx context.Context, interval time.Duration, parser Parser) *Cron {
	ctx, cancel := context.WithCancel(ctx)
	return &Cron{
		interval: interval,
		ctx:      ctx,
		cancel:   cancel,
		parser:   parser,
	}
}

func (c *Cron) Start() {
	log.Info().Msg("cron is running")
	// c.parser.Parse(c.ctx)
	
	ticker := time.NewTicker(c.interval)
	for {
		select {
		case <-c.ctx.Done():
			ticker.Stop()
			return
		case <-ticker.C:
			c.parser.Parse(c.ctx)
		}
	}
}

func (c *Cron) Stop() {
	log.Info().Msg("cron stopped")
	c.cancel()
}

package server

import (
	"context"
	"net/http"
	"time"
)

type Cron interface {
	Start()
	Stop()
}

type Server struct {
	httpServer *http.Server
	cron       Cron
}

func New(handler http.Handler, cron Cron, addr string) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         addr,
			Handler:      handler,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
		cron: cron,
	}
}

func (s *Server) Run() error {
	go s.cron.Start()
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.cron.Stop()
	return s.httpServer.Shutdown(ctx)
}

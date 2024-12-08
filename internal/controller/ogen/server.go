package ogen

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	api "github.com/himmel520/uoffer/require/api/oas"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(h *Handler, addr string) (*Server, error) {
	apiServer, err := api.NewServer(
		h,
		h.Auth,
	)
	if err != nil {
		return nil, fmt.Errorf("init http server: %w", err)
	}

	mux := chi.NewMux()
	mux.Use(SetCors())
	mux.Use(middleware.RequestID)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)

	swaggerUI(mux)

	mux.Mount("/", apiServer)

	return &Server{
		httpServer: &http.Server{
			Addr:         addr,
			Handler:      mux,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}, nil
}

func (s *Server) Run() error {
	err := s.httpServer.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}
	return err
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

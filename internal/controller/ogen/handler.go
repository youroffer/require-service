package ogen

import (
	"context"

	api "github.com/himmel520/uoffer/require/api/oas"
)

// Default pagination
const (
	Page    = 0
	PerPage = 20
)

type (
	Handler struct {
		Auth
		Error
	}

	Auth interface {
	}

	Error interface {
		NewError(ctx context.Context, err error) *api.ErrorStatusCode
	}
)

type HandlerParams struct {
	Auth
	Error
}

func NewHandler(params HandlerParams) *Handler {
	return &Handler{
		Auth:  params.Auth,
		Error: params.Error,
	}
}

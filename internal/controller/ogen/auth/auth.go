package auth

import (
	"context"
	"errors"

	api "github.com/himmel520/uoffer/require/api/oas"
	"github.com/himmel520/uoffer/require/internal/entity"
)

func (h *Handler) HandleAdminBearerAuth(ctx context.Context, operationName string, t api.AdminBearerAuth) (context.Context, error) {
	role, err := h.uc.GetUserRoleFromToken(t.GetToken())
	if err != nil {
		return ctx, err
	}

	if role < entity.RoleAdmin {
		return nil, errors.New("not enough permissions")
	}

	return ctx, nil
}

func (h *Handler) HandleUserBearerAuth(ctx context.Context, operationName string, t api.UserBearerAuth) (context.Context, error) {
	role, err := h.uc.GetUserRoleFromToken(t.GetToken())
	if err != nil {
		return ctx, err
	}

	if role < entity.RoleUser {
		return ctx, errors.New("not enough permissions")
	}

	return ctx, nil
}

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

var _ api.Handler = new(Handler)

type (
	Handler struct {
		Auth
		Error
		Analytic
		Category
		Filter
		Position
	}

	Auth interface {
		HandleAdminBearerAuth(ctx context.Context, operationName string, t api.AdminBearerAuth) (context.Context, error)
		HandleUserBearerAuth(ctx context.Context, operationName string, t api.UserBearerAuth) (context.Context, error)
	}

	Error interface {
		NewError(ctx context.Context, err error) *api.ErrorStatusCode
	}

	Analytic interface {
		V1AdminAnalyticsAnalyticIDDelete(ctx context.Context, params api.V1AdminAnalyticsAnalyticIDDeleteParams) (api.V1AdminAnalyticsAnalyticIDDeleteRes, error)
		V1AdminAnalyticsAnalyticIDPut(ctx context.Context, req *api.AnalyticPut, params api.V1AdminAnalyticsAnalyticIDPutParams) (api.V1AdminAnalyticsAnalyticIDPutRes, error)
		V1AdminAnalyticsGet(ctx context.Context, params api.V1AdminAnalyticsGetParams) (api.V1AdminAnalyticsGetRes, error)
		V1AdminAnalyticsPost(ctx context.Context, req *api.AnalyticPost) (api.V1AdminAnalyticsPostRes, error)
		V1AnalyticsAnalyticIDGet(ctx context.Context, params api.V1AnalyticsAnalyticIDGetParams) (api.V1AnalyticsAnalyticIDGetRes, error)
		V1AnalyticsAnalyticIDLimitGet(ctx context.Context, params api.V1AnalyticsAnalyticIDLimitGetParams) (api.V1AnalyticsAnalyticIDLimitGetRes, error)
	}

	Category interface {
		V1AdminCategoriesCategoryIDDelete(ctx context.Context, params api.V1AdminCategoriesCategoryIDDeleteParams) (api.V1AdminCategoriesCategoryIDDeleteRes, error)
		V1AdminCategoriesCategoryIDPut(ctx context.Context, req *api.CategoryPut, params api.V1AdminCategoriesCategoryIDPutParams) (api.V1AdminCategoriesCategoryIDPutRes, error)
		V1AdminCategoriesGet(ctx context.Context, params api.V1AdminCategoriesGetParams) (api.V1AdminCategoriesGetRes, error)
		V1AdminCategoriesPost(ctx context.Context, req *api.CategoryPost) (api.V1AdminCategoriesPostRes, error)
		V1CategoriesGet(ctx context.Context) (api.V1CategoriesGetRes, error)
	}

	Filter interface {
		V1AdminFiltersFilterIDDelete(ctx context.Context, params api.V1AdminFiltersFilterIDDeleteParams) (api.V1AdminFiltersFilterIDDeleteRes, error)
		V1AdminFiltersGet(ctx context.Context, params api.V1AdminFiltersGetParams) (api.V1AdminFiltersGetRes, error)
		V1AdminFiltersPost(ctx context.Context, req *api.V1AdminFiltersPostReq) (api.V1AdminFiltersPostRes, error)
	}

	Position interface {
		V1AdminPositionsGet(ctx context.Context, params api.V1AdminPositionsGetParams) (api.V1AdminPositionsGetRes, error)
		V1AdminPositionsPositionIDDelete(ctx context.Context, params api.V1AdminPositionsPositionIDDeleteParams) (api.V1AdminPositionsPositionIDDeleteRes, error)
		V1AdminPositionsPositionIDPut(ctx context.Context, req *api.PositionPut, params api.V1AdminPositionsPositionIDPutParams) (api.V1AdminPositionsPositionIDPutRes, error)
		V1AdminPositionsPost(ctx context.Context, req *api.PositionPost) (api.V1AdminPositionsPostRes, error)
	}
)

type HandlerParams struct {
	Auth
	Error
	Analytic
	Category
	Filter
	Position
}

func NewHandler(params HandlerParams) *Handler {
	return &Handler{
		Auth:     params.Auth,
		Error:    params.Error,
		Analytic: params.Analytic,
		Category: params.Category,
		Filter:   params.Filter,
		Position: params.Position,
	}
}

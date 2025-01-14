// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// V1AdminAnalyticsAnalyticIDDelete implements DELETE /v1/admin/analytics/{analyticID} operation.
//
// Удаляет аналитику по уникальному идентификатору.
//
// DELETE /v1/admin/analytics/{analyticID}
func (UnimplementedHandler) V1AdminAnalyticsAnalyticIDDelete(ctx context.Context, params V1AdminAnalyticsAnalyticIDDeleteParams) (r V1AdminAnalyticsAnalyticIDDeleteRes, _ error) {
	return r, ht.ErrNotImplemented
}

// V1AdminAnalyticsAnalyticIDPut implements PUT /v1/admin/analytics/{analyticID} operation.
//
// Обновляет аналитику по ее уникальному
// идентификатору.
//
// PUT /v1/admin/analytics/{analyticID}
func (UnimplementedHandler) V1AdminAnalyticsAnalyticIDPut(ctx context.Context, req *AnalyticPut, params V1AdminAnalyticsAnalyticIDPutParams) (r V1AdminAnalyticsAnalyticIDPutRes, _ error) {
	return r, ht.ErrNotImplemented
}

// V1AdminAnalyticsGet implements GET /v1/admin/analytics operation.
//
// Возвращает список всех аналитик с возможностью
// пагинации.
//
// GET /v1/admin/analytics
func (UnimplementedHandler) V1AdminAnalyticsGet(ctx context.Context, params V1AdminAnalyticsGetParams) (r V1AdminAnalyticsGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// V1AdminAnalyticsPost implements POST /v1/admin/analytics operation.
//
// Создает новую запись аналитики.
//
// POST /v1/admin/analytics
func (UnimplementedHandler) V1AdminAnalyticsPost(ctx context.Context, req *AnalyticPost) (r V1AdminAnalyticsPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// V1AdminCategoriesCategoryIDDelete implements DELETE /v1/admin/categories/{categoryID} operation.
//
// Удаляет категорию по ее уникальному идентификатору.
//
// DELETE /v1/admin/categories/{categoryID}
func (UnimplementedHandler) V1AdminCategoriesCategoryIDDelete(ctx context.Context, params V1AdminCategoriesCategoryIDDeleteParams) (r V1AdminCategoriesCategoryIDDeleteRes, _ error) {
	return r, ht.ErrNotImplemented
}

// V1AdminCategoriesCategoryIDPut implements PUT /v1/admin/categories/{categoryID} operation.
//
// Обновляет категорию по ее уникальному
// идентификатору.
//
// PUT /v1/admin/categories/{categoryID}
func (UnimplementedHandler) V1AdminCategoriesCategoryIDPut(ctx context.Context, req *CategoryPut, params V1AdminCategoriesCategoryIDPutParams) (r V1AdminCategoriesCategoryIDPutRes, _ error) {
	return r, ht.ErrNotImplemented
}

// V1AdminCategoriesGet implements GET /v1/admin/categories operation.
//
// Возвращает список всех категорий с возможностью
// пагинации.
//
// GET /v1/admin/categories
func (UnimplementedHandler) V1AdminCategoriesGet(ctx context.Context, params V1AdminCategoriesGetParams) (r V1AdminCategoriesGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// V1AdminCategoriesPost implements POST /v1/admin/categories operation.
//
// Создает новую категорию.
//
// POST /v1/admin/categories
func (UnimplementedHandler) V1AdminCategoriesPost(ctx context.Context, req *CategoryPost) (r V1AdminCategoriesPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// V1AdminFiltersFilterIDDelete implements DELETE /v1/admin/filters/{filterID} operation.
//
// Удаляет фильтр по его уникальному идентификатору.
//
// DELETE /v1/admin/filters/{filterID}
func (UnimplementedHandler) V1AdminFiltersFilterIDDelete(ctx context.Context, params V1AdminFiltersFilterIDDeleteParams) (r V1AdminFiltersFilterIDDeleteRes, _ error) {
	return r, ht.ErrNotImplemented
}

// V1AdminFiltersGet implements GET /v1/admin/filters operation.
//
// Возвращает список всех фильтров с возможностью
// пагинации.
//
// GET /v1/admin/filters
func (UnimplementedHandler) V1AdminFiltersGet(ctx context.Context, params V1AdminFiltersGetParams) (r V1AdminFiltersGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// V1AdminFiltersPost implements POST /v1/admin/filters operation.
//
// Создает новый фильтр с уникальным словом.
//
// POST /v1/admin/filters
func (UnimplementedHandler) V1AdminFiltersPost(ctx context.Context, req *V1AdminFiltersPostReq) (r V1AdminFiltersPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// V1AdminPositionsGet implements GET /v1/admin/positions operation.
//
// Возвращает список всех должностей с возможностью
// пагинации.
//
// GET /v1/admin/positions
func (UnimplementedHandler) V1AdminPositionsGet(ctx context.Context, params V1AdminPositionsGetParams) (r V1AdminPositionsGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// V1AdminPositionsPositionIDDelete implements DELETE /v1/admin/positions/{positionID} operation.
//
// Удаляет должность по уникальному идентификатору.
//
// DELETE /v1/admin/positions/{positionID}
func (UnimplementedHandler) V1AdminPositionsPositionIDDelete(ctx context.Context, params V1AdminPositionsPositionIDDeleteParams) (r V1AdminPositionsPositionIDDeleteRes, _ error) {
	return r, ht.ErrNotImplemented
}

// V1AdminPositionsPositionIDPut implements PUT /v1/admin/positions/{positionID} operation.
//
// Обновляет должность по уникальному идентификатору.
//
// PUT /v1/admin/positions/{positionID}
func (UnimplementedHandler) V1AdminPositionsPositionIDPut(ctx context.Context, req *PositionPut, params V1AdminPositionsPositionIDPutParams) (r V1AdminPositionsPositionIDPutRes, _ error) {
	return r, ht.ErrNotImplemented
}

// V1AdminPositionsPost implements POST /v1/admin/positions operation.
//
// Создает новую должность.
//
// POST /v1/admin/positions
func (UnimplementedHandler) V1AdminPositionsPost(ctx context.Context, req *PositionPost) (r V1AdminPositionsPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// V1AnalyticsAnalyticIDGet implements GET /v1/analytics/{analyticID} operation.
//
// Возвращает аналитику со всеми словами по уникальному
// идентификатору аналитики.
//
// GET /v1/analytics/{analyticID}
func (UnimplementedHandler) V1AnalyticsAnalyticIDGet(ctx context.Context, params V1AnalyticsAnalyticIDGetParams) (r V1AnalyticsAnalyticIDGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// V1AnalyticsAnalyticIDLimitGet implements GET /v1/analytics/{analyticID}/limit operation.
//
// Возвращает аналитику с ограничением на слова по
// уникальному идентификатору аналитики.
//
// GET /v1/analytics/{analyticID}/limit
func (UnimplementedHandler) V1AnalyticsAnalyticIDLimitGet(ctx context.Context, params V1AnalyticsAnalyticIDLimitGetParams) (r V1AnalyticsAnalyticIDLimitGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// V1CategoriesGet implements GET /v1/categories operation.
//
// Возвращает все категории с публичными должностями.
//
// GET /v1/categories
func (UnimplementedHandler) V1CategoriesGet(ctx context.Context) (r V1CategoriesGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// NewError creates *ErrorStatusCode from error returned by handler.
//
// Used for common default response.
func (UnimplementedHandler) NewError(ctx context.Context, err error) (r *ErrorStatusCode) {
	r = new(ErrorStatusCode)
	return r
}

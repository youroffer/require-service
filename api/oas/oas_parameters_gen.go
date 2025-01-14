// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"net/url"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// V1AdminAnalyticsAnalyticIDDeleteParams is parameters of DELETE /v1/admin/analytics/{analyticID} operation.
type V1AdminAnalyticsAnalyticIDDeleteParams struct {
	// Уникальный идентификатор аналитики.
	AnalyticID int
}

func unpackV1AdminAnalyticsAnalyticIDDeleteParams(packed middleware.Parameters) (params V1AdminAnalyticsAnalyticIDDeleteParams) {
	{
		key := middleware.ParameterKey{
			Name: "analyticID",
			In:   "path",
		}
		params.AnalyticID = packed[key].(int)
	}
	return params
}

func decodeV1AdminAnalyticsAnalyticIDDeleteParams(args [1]string, argsEscaped bool, r *http.Request) (params V1AdminAnalyticsAnalyticIDDeleteParams, _ error) {
	// Decode path: analyticID.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "analyticID",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.AnalyticID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "analyticID",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// V1AdminAnalyticsAnalyticIDPutParams is parameters of PUT /v1/admin/analytics/{analyticID} operation.
type V1AdminAnalyticsAnalyticIDPutParams struct {
	// Уникальный идентификатор аналитики.
	AnalyticID int
}

func unpackV1AdminAnalyticsAnalyticIDPutParams(packed middleware.Parameters) (params V1AdminAnalyticsAnalyticIDPutParams) {
	{
		key := middleware.ParameterKey{
			Name: "analyticID",
			In:   "path",
		}
		params.AnalyticID = packed[key].(int)
	}
	return params
}

func decodeV1AdminAnalyticsAnalyticIDPutParams(args [1]string, argsEscaped bool, r *http.Request) (params V1AdminAnalyticsAnalyticIDPutParams, _ error) {
	// Decode path: analyticID.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "analyticID",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.AnalyticID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "analyticID",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// V1AdminAnalyticsGetParams is parameters of GET /v1/admin/analytics operation.
type V1AdminAnalyticsGetParams struct {
	// Номер страницы для пагинации.
	Page OptInt
	// Количество записей на странице.
	PerPage OptInt
}

func unpackV1AdminAnalyticsGetParams(packed middleware.Parameters) (params V1AdminAnalyticsGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "page",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Page = v.(OptInt)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "per_page",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.PerPage = v.(OptInt)
		}
	}
	return params
}

func decodeV1AdminAnalyticsGetParams(args [0]string, argsEscaped bool, r *http.Request) (params V1AdminAnalyticsGetParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Set default value for query: page.
	{
		val := int(0)
		params.Page.SetTo(val)
	}
	// Decode query: page.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPageVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Page.SetTo(paramsDotPageVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.Page.Get(); ok {
					if err := func() error {
						if err := (validate.Int{
							MinSet:        true,
							Min:           0,
							MaxSet:        false,
							Max:           0,
							MinExclusive:  false,
							MaxExclusive:  false,
							MultipleOfSet: false,
							MultipleOf:    0,
						}).Validate(int64(value)); err != nil {
							return errors.Wrap(err, "int")
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "page",
			In:   "query",
			Err:  err,
		}
	}
	// Set default value for query: per_page.
	{
		val := int(20)
		params.PerPage.SetTo(val)
	}
	// Decode query: per_page.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "per_page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPerPageVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotPerPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.PerPage.SetTo(paramsDotPerPageVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.PerPage.Get(); ok {
					if err := func() error {
						if err := (validate.Int{
							MinSet:        true,
							Min:           1,
							MaxSet:        false,
							Max:           0,
							MinExclusive:  false,
							MaxExclusive:  false,
							MultipleOfSet: false,
							MultipleOf:    0,
						}).Validate(int64(value)); err != nil {
							return errors.Wrap(err, "int")
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "per_page",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// V1AdminCategoriesCategoryIDDeleteParams is parameters of DELETE /v1/admin/categories/{categoryID} operation.
type V1AdminCategoriesCategoryIDDeleteParams struct {
	// Уникальный идентификатор категории.
	CategoryID int
}

func unpackV1AdminCategoriesCategoryIDDeleteParams(packed middleware.Parameters) (params V1AdminCategoriesCategoryIDDeleteParams) {
	{
		key := middleware.ParameterKey{
			Name: "categoryID",
			In:   "path",
		}
		params.CategoryID = packed[key].(int)
	}
	return params
}

func decodeV1AdminCategoriesCategoryIDDeleteParams(args [1]string, argsEscaped bool, r *http.Request) (params V1AdminCategoriesCategoryIDDeleteParams, _ error) {
	// Decode path: categoryID.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "categoryID",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.CategoryID = c
				return nil
			}(); err != nil {
				return err
			}
			if err := func() error {
				if err := (validate.Int{
					MinSet:        true,
					Min:           1,
					MaxSet:        false,
					Max:           0,
					MinExclusive:  false,
					MaxExclusive:  false,
					MultipleOfSet: false,
					MultipleOf:    0,
				}).Validate(int64(params.CategoryID)); err != nil {
					return errors.Wrap(err, "int")
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "categoryID",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// V1AdminCategoriesCategoryIDPutParams is parameters of PUT /v1/admin/categories/{categoryID} operation.
type V1AdminCategoriesCategoryIDPutParams struct {
	// Уникальный идентификатор категории.
	CategoryID int
}

func unpackV1AdminCategoriesCategoryIDPutParams(packed middleware.Parameters) (params V1AdminCategoriesCategoryIDPutParams) {
	{
		key := middleware.ParameterKey{
			Name: "categoryID",
			In:   "path",
		}
		params.CategoryID = packed[key].(int)
	}
	return params
}

func decodeV1AdminCategoriesCategoryIDPutParams(args [1]string, argsEscaped bool, r *http.Request) (params V1AdminCategoriesCategoryIDPutParams, _ error) {
	// Decode path: categoryID.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "categoryID",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.CategoryID = c
				return nil
			}(); err != nil {
				return err
			}
			if err := func() error {
				if err := (validate.Int{
					MinSet:        true,
					Min:           1,
					MaxSet:        false,
					Max:           0,
					MinExclusive:  false,
					MaxExclusive:  false,
					MultipleOfSet: false,
					MultipleOf:    0,
				}).Validate(int64(params.CategoryID)); err != nil {
					return errors.Wrap(err, "int")
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "categoryID",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// V1AdminCategoriesGetParams is parameters of GET /v1/admin/categories operation.
type V1AdminCategoriesGetParams struct {
	// Номер страницы для пагинации.
	Page OptInt
	// Количество категорий на странице.
	PerPage OptInt
}

func unpackV1AdminCategoriesGetParams(packed middleware.Parameters) (params V1AdminCategoriesGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "page",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Page = v.(OptInt)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "per_page",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.PerPage = v.(OptInt)
		}
	}
	return params
}

func decodeV1AdminCategoriesGetParams(args [0]string, argsEscaped bool, r *http.Request) (params V1AdminCategoriesGetParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Set default value for query: page.
	{
		val := int(0)
		params.Page.SetTo(val)
	}
	// Decode query: page.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPageVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Page.SetTo(paramsDotPageVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.Page.Get(); ok {
					if err := func() error {
						if err := (validate.Int{
							MinSet:        true,
							Min:           0,
							MaxSet:        false,
							Max:           0,
							MinExclusive:  false,
							MaxExclusive:  false,
							MultipleOfSet: false,
							MultipleOf:    0,
						}).Validate(int64(value)); err != nil {
							return errors.Wrap(err, "int")
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "page",
			In:   "query",
			Err:  err,
		}
	}
	// Set default value for query: per_page.
	{
		val := int(20)
		params.PerPage.SetTo(val)
	}
	// Decode query: per_page.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "per_page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPerPageVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotPerPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.PerPage.SetTo(paramsDotPerPageVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.PerPage.Get(); ok {
					if err := func() error {
						if err := (validate.Int{
							MinSet:        true,
							Min:           1,
							MaxSet:        false,
							Max:           0,
							MinExclusive:  false,
							MaxExclusive:  false,
							MultipleOfSet: false,
							MultipleOf:    0,
						}).Validate(int64(value)); err != nil {
							return errors.Wrap(err, "int")
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "per_page",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// V1AdminFiltersFilterIDDeleteParams is parameters of DELETE /v1/admin/filters/{filterID} operation.
type V1AdminFiltersFilterIDDeleteParams struct {
	// Уникальный идентификатор фильтра.
	FilterID int
}

func unpackV1AdminFiltersFilterIDDeleteParams(packed middleware.Parameters) (params V1AdminFiltersFilterIDDeleteParams) {
	{
		key := middleware.ParameterKey{
			Name: "filterID",
			In:   "path",
		}
		params.FilterID = packed[key].(int)
	}
	return params
}

func decodeV1AdminFiltersFilterIDDeleteParams(args [1]string, argsEscaped bool, r *http.Request) (params V1AdminFiltersFilterIDDeleteParams, _ error) {
	// Decode path: filterID.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "filterID",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.FilterID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "filterID",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// V1AdminFiltersGetParams is parameters of GET /v1/admin/filters operation.
type V1AdminFiltersGetParams struct {
	// Номер страницы для пагинации.
	Page OptInt
	// Количество фильтров на странице.
	PerPage OptInt
}

func unpackV1AdminFiltersGetParams(packed middleware.Parameters) (params V1AdminFiltersGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "page",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Page = v.(OptInt)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "per_page",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.PerPage = v.(OptInt)
		}
	}
	return params
}

func decodeV1AdminFiltersGetParams(args [0]string, argsEscaped bool, r *http.Request) (params V1AdminFiltersGetParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Set default value for query: page.
	{
		val := int(0)
		params.Page.SetTo(val)
	}
	// Decode query: page.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPageVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Page.SetTo(paramsDotPageVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.Page.Get(); ok {
					if err := func() error {
						if err := (validate.Int{
							MinSet:        true,
							Min:           0,
							MaxSet:        false,
							Max:           0,
							MinExclusive:  false,
							MaxExclusive:  false,
							MultipleOfSet: false,
							MultipleOf:    0,
						}).Validate(int64(value)); err != nil {
							return errors.Wrap(err, "int")
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "page",
			In:   "query",
			Err:  err,
		}
	}
	// Set default value for query: per_page.
	{
		val := int(20)
		params.PerPage.SetTo(val)
	}
	// Decode query: per_page.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "per_page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPerPageVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotPerPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.PerPage.SetTo(paramsDotPerPageVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.PerPage.Get(); ok {
					if err := func() error {
						if err := (validate.Int{
							MinSet:        true,
							Min:           1,
							MaxSet:        false,
							Max:           0,
							MinExclusive:  false,
							MaxExclusive:  false,
							MultipleOfSet: false,
							MultipleOf:    0,
						}).Validate(int64(value)); err != nil {
							return errors.Wrap(err, "int")
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "per_page",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// V1AdminPositionsGetParams is parameters of GET /v1/admin/positions operation.
type V1AdminPositionsGetParams struct {
	// Номер страницы для пагинации.
	Page OptInt
	// Количество должностей на странице.
	PerPage OptInt
}

func unpackV1AdminPositionsGetParams(packed middleware.Parameters) (params V1AdminPositionsGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "page",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Page = v.(OptInt)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "per_page",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.PerPage = v.(OptInt)
		}
	}
	return params
}

func decodeV1AdminPositionsGetParams(args [0]string, argsEscaped bool, r *http.Request) (params V1AdminPositionsGetParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Set default value for query: page.
	{
		val := int(0)
		params.Page.SetTo(val)
	}
	// Decode query: page.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPageVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Page.SetTo(paramsDotPageVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.Page.Get(); ok {
					if err := func() error {
						if err := (validate.Int{
							MinSet:        true,
							Min:           0,
							MaxSet:        false,
							Max:           0,
							MinExclusive:  false,
							MaxExclusive:  false,
							MultipleOfSet: false,
							MultipleOf:    0,
						}).Validate(int64(value)); err != nil {
							return errors.Wrap(err, "int")
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "page",
			In:   "query",
			Err:  err,
		}
	}
	// Set default value for query: per_page.
	{
		val := int(20)
		params.PerPage.SetTo(val)
	}
	// Decode query: per_page.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "per_page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPerPageVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotPerPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.PerPage.SetTo(paramsDotPerPageVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.PerPage.Get(); ok {
					if err := func() error {
						if err := (validate.Int{
							MinSet:        true,
							Min:           1,
							MaxSet:        false,
							Max:           0,
							MinExclusive:  false,
							MaxExclusive:  false,
							MultipleOfSet: false,
							MultipleOf:    0,
						}).Validate(int64(value)); err != nil {
							return errors.Wrap(err, "int")
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "per_page",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// V1AdminPositionsPositionIDDeleteParams is parameters of DELETE /v1/admin/positions/{positionID} operation.
type V1AdminPositionsPositionIDDeleteParams struct {
	// Уникальный идентификатор должности.
	PositionID int
}

func unpackV1AdminPositionsPositionIDDeleteParams(packed middleware.Parameters) (params V1AdminPositionsPositionIDDeleteParams) {
	{
		key := middleware.ParameterKey{
			Name: "positionID",
			In:   "path",
		}
		params.PositionID = packed[key].(int)
	}
	return params
}

func decodeV1AdminPositionsPositionIDDeleteParams(args [1]string, argsEscaped bool, r *http.Request) (params V1AdminPositionsPositionIDDeleteParams, _ error) {
	// Decode path: positionID.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "positionID",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.PositionID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "positionID",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// V1AdminPositionsPositionIDPutParams is parameters of PUT /v1/admin/positions/{positionID} operation.
type V1AdminPositionsPositionIDPutParams struct {
	// Уникальный идентификатор должности.
	PositionID int
}

func unpackV1AdminPositionsPositionIDPutParams(packed middleware.Parameters) (params V1AdminPositionsPositionIDPutParams) {
	{
		key := middleware.ParameterKey{
			Name: "positionID",
			In:   "path",
		}
		params.PositionID = packed[key].(int)
	}
	return params
}

func decodeV1AdminPositionsPositionIDPutParams(args [1]string, argsEscaped bool, r *http.Request) (params V1AdminPositionsPositionIDPutParams, _ error) {
	// Decode path: positionID.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "positionID",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.PositionID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "positionID",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// V1AnalyticsAnalyticIDGetParams is parameters of GET /v1/analytics/{analyticID} operation.
type V1AnalyticsAnalyticIDGetParams struct {
	// Уникальный идентификатор аналитики.
	AnalyticID int
}

func unpackV1AnalyticsAnalyticIDGetParams(packed middleware.Parameters) (params V1AnalyticsAnalyticIDGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "analyticID",
			In:   "path",
		}
		params.AnalyticID = packed[key].(int)
	}
	return params
}

func decodeV1AnalyticsAnalyticIDGetParams(args [1]string, argsEscaped bool, r *http.Request) (params V1AnalyticsAnalyticIDGetParams, _ error) {
	// Decode path: analyticID.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "analyticID",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.AnalyticID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "analyticID",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// V1AnalyticsAnalyticIDLimitGetParams is parameters of GET /v1/analytics/{analyticID}/limit operation.
type V1AnalyticsAnalyticIDLimitGetParams struct {
	// Уникальный идентификатор аналитики.
	AnalyticID int
}

func unpackV1AnalyticsAnalyticIDLimitGetParams(packed middleware.Parameters) (params V1AnalyticsAnalyticIDLimitGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "analyticID",
			In:   "path",
		}
		params.AnalyticID = packed[key].(int)
	}
	return params
}

func decodeV1AnalyticsAnalyticIDLimitGetParams(args [1]string, argsEscaped bool, r *http.Request) (params V1AnalyticsAnalyticIDLimitGetParams, _ error) {
	// Decode path: analyticID.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "analyticID",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.AnalyticID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "analyticID",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

package entity

import (
	"time"

	api "github.com/himmel520/uoffer/require/api/oas"
	"github.com/himmel520/uoffer/require/internal/lib/convert"
)

type Analytic struct {
	ID           int
	PostID       int
	SearchQuery  string
	ParseAt      time.Time
	VacanciesNum int
}

type AnalyticUpdate struct {
	PostID      Optional[int]
	SearchQuery Optional[string]
}

type AnalyticResp struct {
	ID           int
	PostTitle    string
	SearchQuery  string
	ParseAt      Optional[time.Time]
	VacanciesNum Optional[int]
}

type TopWords struct {
	Word     string
	Mentions int
}

type AnalyticWithWords struct {
	Analytic *AnalyticResp
	Skills   []*TopWords
	Keywords []*TopWords
}

func AnalyticRespToApi(a *AnalyticResp) *api.Analytic {
	return &api.Analytic{
		ID:           a.ID,
		PostTitle:    a.PostTitle,
		SearchQuery:  a.SearchQuery,
		ParseAt:      api.OptDateTime{Value: a.ParseAt.Value, Set: a.ParseAt.Set},
		VacanciesNum: api.OptInt{Value: a.VacanciesNum.Value, Set: a.VacanciesNum.Set},
	}
}

func (a *AnalyticUpdate) IsSet() bool {
	return a.PostID.Set || a.SearchQuery.Set
}

type AnalyticsResp struct {
	Data    []*AnalyticResp
	Page    uint64
	Pages   uint64
	PerPage uint64
}

func (c *AnalyticsResp) ToApi() *api.AnalyticsResp {
	return &api.AnalyticsResp{
		Data:    convert.ApplyPointerToSlice(c.Data, AnalyticRespToApi),
		Page:    int(c.Page),
		Pages:   int(c.Pages),
		PerPage: int(c.PerPage),
	}
}

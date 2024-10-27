package entity

import (
	"time"
)

type Analytic struct {
	ID           int       `json:"id,omitempty"`
	PostID       int       `json:"posts_id" binding:"required,min=1"`
	SearchQuery  string    `json:"search_query" binding:"required"`
	LastUpdated  time.Time `json:"last_updated,omitempty"`
	VacanciesNum int       `json:"vacancies_num,omitempty"`
}

type AnalyticUpdate struct {
	PostID      *int    `json:"posts_id" binding:"omitempty,min=1"`
	SearchQuery *string `json:"search_query,omitempty"`
}

type AnalyticResp struct {
	ID           int        `json:"id"`
	SearchQuery  string     `json:"search_query"`
	LastUpdated  *time.Time `json:"last_updated,omitempty"`
	VacanciesNum *int       `json:"vacancies_num,omitempty"`
}

type TopWords struct {
	Word      string `json:"word"`
	Reference int    `json:"reference"`
}

type AnalyticWithWords struct {
	Analytic *AnalyticResp `json:"analytic"`
	Skills   []*TopWords   `json:"skills"`
	Keywords []*TopWords   `json:"keywords"`
}

func (a *AnalyticUpdate) IsEmpty() bool {
	return a.PostID == nil && a.SearchQuery == nil
}

package entity

import api "github.com/himmel520/uoffer/require/api/oas"

type Position struct {
	ID           int    `json:"id,omitempty"`
	CategoriesID int    `json:"categories_id"`
	LogoID       int    `json:"logo_id"`
	Title        string `json:"title"`
	Public       bool   `json:"public"`
}

type PositionResp struct {
	ID       int      `json:"id,omitempty"`
	Category Category `json:"categories_id"`
	LogoID   int      `json:"logo_id"`
	Title    string   `json:"title"`
	Public   bool     `json:"public"`
}

func PositionRespToApi(p *PositionResp) *api.Position {
	return &api.Position{
		ID: api.NewOptInt(p.ID),
		Category: api.NewOptCategory(api.Category{
			ID:     p.Category.ID,
			Title:  p.Category.Title,
			Public: p.Category.Public,
		}),
		LogoID: p.LogoID,
		Title:  p.Title,
		Public: p.Public,
	}
}

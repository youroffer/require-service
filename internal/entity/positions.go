package entity

import (
	api "github.com/himmel520/uoffer/require/api/oas"
	"github.com/himmel520/uoffer/require/internal/lib/convert"
)

type Position struct {
	ID           int    `json:"id,omitempty"`
	CategoriesID int    `json:"categories_id"`
	LogoID       int    `json:"logo_id"`
	Title        string `json:"title"`
	Public       bool   `json:"public"`
}

type PositionUpdate struct {
	CategoriesID Optional[int]    `json:"categories_id"`
	LogoID       Optional[int]    `json:"logo_id"`
	Title        Optional[string] `json:"title"`
	Public       Optional[bool]   `json:"public"`
}

func (p *PositionUpdate) IsSet() bool {
	return p.CategoriesID.Set || p.LogoID.Set || p.Title.Set || p.Public.Set
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

type PositionsResp struct {
	Data    []*PositionResp `json:"data"`
	Page    uint64          `json:"page"`
	Pages   uint64          `json:"pages"`
	PerPage uint64          `json:"per_page"`
}

func (c *PositionsResp) ToApi() *api.PositionsResp {
	return &api.PositionsResp{
		Data:    convert.ApplyPointerToSlice(c.Data, PositionRespToApi),
		Page:    api.NewOptInt(int(c.Page)),
		Total:   api.NewOptInt(int(c.Pages)),
		PerPage: api.NewOptInt(int(c.PerPage)),
	}
}

package entity

import (
	api "github.com/himmel520/uoffer/require/api/oas"
	"github.com/himmel520/uoffer/require/internal/lib/convert"
)

type CategoryUpdate struct {
	Title  Optional[string]
	Public Optional[bool]
}

func (c *CategoryUpdate) IsSet() bool {
	return c.Title.Set || c.Public.Set
}

type Category struct {
	ID     int
	Title  string
	Public bool
}

func ConvertCategoryToApi(f *Category) *api.Category {
	return &api.Category{
		ID:     f.ID,
		Title:  f.Title,
		Public: f.Public,
	}
}

type CategoriesResp struct {
	Data    []*Category
	Page    uint64
	Pages   uint64
	PerPage uint64
}

func (c *CategoriesResp) ToApi() *api.CategoriesResp {
	return &api.CategoriesResp{
		Data:    convert.ApplyPointerToSlice(c.Data, ConvertCategoryToApi),
		Page:    int(c.Page),
		Pages:   int(c.Pages),
		PerPage: int(c.PerPage),
	}
}
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

func ConvertCategoryToApi(c *Category) *api.Category {
	return &api.Category{
		ID:     c.ID,
		Title:  c.Title,
		Public: c.Public,
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

type CategoryPosition struct {
	ID     int    `json:"id"`
	LogoID int    `json:"logo_id"`
	Title  string `json:"title"`
	Public bool   `json:"public"`
}

func ConvertCategoryPositionToApi(position CategoryPosition) api.CategoryPosition {
	return api.CategoryPosition{
		ID:     position.ID,
		LogoID: position.LogoID,
		Title:  position.Title,
		Public: position.Public,
	}
}

type CategoriesPublicPostsResp map[string][]CategoryPosition

func ConvertCategoriesPublicPostsRespToApi(categories CategoriesPublicPostsResp) *api.CategoriesPostsResp {
	apiCategories := api.CategoriesPostsResp{}

	for category, position := range categories {
		apiCategories[category] = convert.ApplyToSlice(position, ConvertCategoryPositionToApi)
	}

	return &apiCategories
}

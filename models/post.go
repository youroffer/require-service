package models

type Post struct {
	ID           int    `json:"id"`
	CategoriesID int    `json:"categories_id" binding:"required,min=1"`
	LogosID      int    `json:"logos_id" binding:"required,min=1"`
	Title        string `json:"title" binding:"required,min=3,max=100"`
	Public       bool   `json:"public"`
}

type PostUpdate struct {
	CategoriesID *int    `json:"categories_id" binding:"omitempty,min=1"`
	LogosID      *int    `json:"logos_id" binding:"omitempty,min=1"`
	Title        *string `json:"title" binding:"omitempty,min=3,max=100"`
	Public       *bool   `json:"public" binding:"omitempty"`
}

func (p *PostUpdate) IsEmpty() bool {
	return p.CategoriesID == nil && p.LogosID == nil && p.Title == nil && p.Public == nil
}

type PostResponse struct {
	ID     int    `json:"id,omitempty"`
	Title  string `json:"title"`
	Public bool   `json:"public"`
	LogoID int    `json:"logo_id"`
}

package entity

type Category struct {
	ID    int    `json:"id,omitempty"`
	Title string `json:"title" binding:"required,min=3,max=50"`
}

type CategoryResponse struct {
	CategoryID int     `json:"id"`
	Posts      []*PostResponse `json:"posts"`
}


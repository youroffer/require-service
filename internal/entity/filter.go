package entity

type Filter struct {
	ID   int    `json:"id,omitempty"`
	Word string `json:"word" binding:"required,min=1"`
}

type FilterResp struct {
	Filters []*Filter `json:"filters"`
	Total   int       `json:"total"`
}

package entity

type Category struct {
	ID    int
	Title string
}

type CategoryResp struct {
	CategoryID int
	Posts      []*PostResponse
}

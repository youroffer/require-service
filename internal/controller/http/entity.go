package httpctrl

type errorResponse struct {
	Message string `json:"message"`
}

type categoryURI struct {
	Category string `uri:"category" binding:"required,min=3,max=100"`
}

type updateCategoryQuery struct {
	Title string `form:"title" binding:"required,min=3,max=100"`
}

type filterURI struct {
	Filter string `uri:"filter" binding:"required,min=1,max=100"`
}

type filterQuery struct {
	Filter string `form:"filter" binding:"required,min=1,max=100"`
}

type PaginationQuery struct {
	Limit  int `form:"limit,default=20" binding:"omitempty,min=1"`
	Offset int `form:"offset,default=0" binding:"omitempty,min=0"`
}

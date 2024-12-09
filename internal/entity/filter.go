package entity

import (
	api "github.com/himmel520/uoffer/require/api/oas"
	"github.com/himmel520/uoffer/require/internal/lib/convert"
)

type Filter struct {
	ID   int
	Word string
}

func ConvertFilterToApi(f *Filter) *api.Filter {
	return &api.Filter{
		ID:   f.ID,
		Word: f.Word,
	}
}

type FiltersResp struct {
	Data    []*Filter
	Page    uint64
	Pages   uint64
	PerPage uint64
}

func (c *FiltersResp) ToApi() *api.FiltersResp {
	return &api.FiltersResp{
		Data:    convert.ApplyPointerToSlice(c.Data, ConvertFilterToApi),
		Page:    int(c.Page),
		Pages:   int(c.Pages),
		PerPage: int(c.PerPage),
	}
}

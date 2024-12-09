package entity

import api "github.com/himmel520/uoffer/require/api/oas"

type Filter struct {
	ID   int
	Word string
}

func (f *Filter) ConvertFilterToApi() *api.Filter {
	return &api.Filter{
		ID:   f.ID,
		Word: f.Word,
	}
}

type FiltersResp struct {
	Filters []*Filter
	Total   int
}

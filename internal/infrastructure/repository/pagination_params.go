package repository

import "github.com/himmel520/uoffer/require/internal/entity"

type PaginationParams struct {
	Limit  entity.Optional[uint64]
	Offset entity.Optional[uint64]
	IDs    entity.Optional[[]int]
}

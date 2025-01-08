package positions

import (
	"context"

	"github.com/himmel520/uoffer/require/internal/infrastructure/repository"
)

func (r *PositionsRepo) Count(ctx context.Context, qe repository.Querier) (int, error) {
	var count int
	err := qe.QueryRow(ctx, `SELECT COUNT(*) FROM posts`).Scan(&count)
	return count, err
}

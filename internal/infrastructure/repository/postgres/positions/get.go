package positions

import (
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"
	"github.com/jackc/pgx/v5"
)

func (r *PositionsRepo) GetByID(ctx context.Context, qe repository.Querier, id int) (*entity.PositionResp, error) {
	query, args, err := squirrel.Select(
		"p.id",
		"c.id",
		"c.title",
		"c.public",
		"p.logo_id",
		"p.title",
		"p.public",
	).
		From("posts AS p").
		Join("categories AS c ON c.id = p.categories_id").
		Where(squirrel.Eq{"p.id": id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	positionResp := &entity.PositionResp{}
	if err = qe.QueryRow(ctx, query, args...).Scan(
		&positionResp.ID,
		&positionResp.Category.ID,
		&positionResp.Category.Title,
		&positionResp.Category.Public,
		&positionResp.LogoID,
		&positionResp.Title,
		&positionResp.Public,
	); err != nil {
		return nil, err
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, repoerr.ErrPostNotFound
	}

	fmt.Println("positionResp", positionResp)

	return positionResp, err
}

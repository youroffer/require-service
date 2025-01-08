package positions

import (
	"context"
	"errors"

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

	return positionResp, err
}

func (r *PositionsRepo) Get(ctx context.Context, qe repository.Querier, params repository.PaginationParams) ([]*entity.PositionResp, error) {
	builder := squirrel.Select(
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
		PlaceholderFormat(squirrel.Dollar)

	if params.Limit.Set {
		builder = builder.Limit(params.Limit.Value)
	}

	if params.Offset.Set {
		builder = builder.Offset(params.Offset.Value)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := qe.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	positionResp := []*entity.PositionResp{}
	for rows.Next() {
		post := &entity.PositionResp{}
		if err = rows.Scan(
			&post.ID,
			&post.Category.ID,
			&post.Category.Title,
			&post.Category.Public,
			&post.LogoID,
			&post.Title,
			&post.Public,
		); err != nil {
			return nil, err
		}

		positionResp = append(positionResp, post)
	}

	if len(positionResp) == 0 {
		return nil, repoerr.ErrPostNotFound
	}

	return positionResp, err
}

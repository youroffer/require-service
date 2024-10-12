package service

import (
	"context"

	"github.com/himmel520/uoffer/require/models"
)

func (s *Service) AddFilter(ctx context.Context, filter string) (*models.Filter, error) {
	return s.repo.AddFilter(ctx, filter)
}

func (s *Service) DeleteFilter(ctx context.Context, filter string) error {
	return s.repo.DeleteFilter(ctx, filter)
}

func (s *Service) GetFilters(ctx context.Context, limit, offset int) (*models.FilterResp, error) {
	filters, err := s.repo.GetFiltersWithPagination(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	count, err := s.repo.GetFiltersCount(ctx)
	if err != nil {
		return nil, err
	}

	return &models.FilterResp{
		Filters: filters,
		Total:   count,
	}, nil
}

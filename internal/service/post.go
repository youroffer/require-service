package service

import (
	"context"

	"github.com/himmel520/uoffer/require/models"
)

func (s *Service) AddPost(ctx context.Context, post *models.Post) (*models.PostResponse, error) {
	return s.repo.AddPost(ctx, post)
}

func (s *Service) UpdatePost(ctx context.Context, id int, post *models.PostUpdate) (*models.PostResponse, error) {
	return s.repo.UpdatePost(ctx, id, post)
}

func (s *Service) DeletePost(ctx context.Context, id int) error {
	return s.repo.DeletePost(ctx, id)
}

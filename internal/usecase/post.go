package usecase

// import (
// 	"context"

// 	"github.com/himmel520/uoffer/require/internal/entity"
// 	"github.com/himmel520/uoffer/require/internal/infrastructure/repository"

// 	"github.com/sirupsen/logrus"
// )

// type PostUsecase struct {
// 	repo repository.PostRepo
// 	log  *logrus.Logger
// }

// func NewPostUsecase(repo repository.PostRepo, log *logrus.Logger) *PostUsecase {
// 	return &PostUsecase{repo: repo, log: log}
// }

// func (uc *PostUsecase) Add(ctx context.Context, post *entity.Post) (*entity.PostResponse, error) {
// 	return uc.repo.Add(ctx, post)
// }

// func (uc *PostUsecase) Update(ctx context.Context, id int, post *entity.PostUpdate) (*entity.PostResponse, error) {
// 	return uc.repo.Update(ctx, id, post)
// }

// func (uc *PostUsecase) Delete(ctx context.Context, id int) error {
// 	return uc.repo.Delete(ctx, id)
// }

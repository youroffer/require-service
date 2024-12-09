package usecase

// import (
// 	"context"

// 	"github.com/himmel520/uoffer/require/internal/entity"
// 	"github.com/himmel520/uoffer/require/internal/infrastructure/repository"

// 	"github.com/sirupsen/logrus"
// )

// type FilterUsecase struct {
// 	repo repository.FilterRepo
// 	log  *logrus.Logger
// }

// func NewFilterUsecase(repo repository.FilterRepo, log *logrus.Logger) *FilterUsecase {
// 	return &FilterUsecase{repo: repo, log: log}
// }

// func (uc *FilterUsecase) Add(ctx context.Context, filter string) (*entity.Filter, error) {
// 	return uc.repo.Add(ctx, filter)
// }

// func (uc *FilterUsecase) Delete(ctx context.Context, filter string) error {
// 	return uc.repo.Delete(ctx, filter)
// }

// func (uc *FilterUsecase) GetWithPagination(ctx context.Context, limit, offset int) (*entity.FilterResp, error) {
// 	filters, err := uc.repo.GetWithPagination(ctx, limit, offset)
// 	if err != nil {
// 		return nil, err
// 	}

// 	count, err := uc.repo.GetCount(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &entity.FilterResp{
// 		Filters: filters,
// 		Total:   count,
// 	}, nil
// }

package usecase

// import (
// 	"context"
// 	"errors"

// 	"github.com/himmel520/uoffer/require/internal/entity"
// 	"github.com/himmel520/uoffer/require/internal/infrastructure/cache"
// 	"github.com/himmel520/uoffer/require/internal/infrastructure/repository"
// 	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"
// 	"github.com/sirupsen/logrus"
// )

// type AnalyticUsecase struct {
// 	repo  repository.AnalyticRepo
// 	cache cache.AnalyticCache
// 	log   *logrus.Logger
// }

// func NewAnalyticUsecase(repo repository.AnalyticRepo, cache cache.AnalyticCache, log *logrus.Logger) *AnalyticUsecase {
// 	return &AnalyticUsecase{repo: repo, cache: cache, log: log}
// }

// func (uc *AnalyticUsecase) clearCache(ctx context.Context, postID int) {
// 	if err := uc.cache.DeleteWithWords(ctx, postID); err != nil {
// 		uc.log.Error(err)
// 	}
// }

// func (uc *AnalyticUsecase) Add(ctx context.Context, analytic *entity.Analytic) (*entity.Analytic, error) {
// 	newAnalytic, err := uc.repo.Add(ctx, analytic)
// 	if err != nil {
// 		return nil, err
// 	}

// 	uc.clearCache(context.Background(), newAnalytic.PostID)

// 	return newAnalytic, nil
// }

// func (uc *AnalyticUsecase) Update(ctx context.Context, id int, analytic *entity.AnalyticUpdate) (*entity.Analytic, error) {
// 	newAnalytic, err := uc.repo.Update(ctx, id, analytic)
// 	if err != nil {
// 		return nil, err
// 	}

// 	uc.clearCache(context.Background(), newAnalytic.PostID)

// 	return newAnalytic, nil
// }

// func (uc *AnalyticUsecase) Delete(ctx context.Context, id int) error {
// 	postID, err := uc.repo.GetPostID(ctx, id)
// 	if err != nil {
// 		return err
// 	}

// 	err = uc.repo.Delete(ctx, id)
// 	if err != nil {
// 		return err
// 	}

// 	uc.clearCache(context.Background(), postID)

// 	return nil
// }

// func (uc *AnalyticUsecase) GetWithWords(ctx context.Context, postID int, role string) (*entity.AnalyticWithWords, error) {
// 	var limit bool
// 	// if role == entity.RoleAnonym {
// 	// 	limit = true
// 	// }

// 	analytic, err := uc.cache.GetWithWords(ctx, postID)
// 	if err != nil {
// 		if !errors.Is(err, repoerr.ErrAnalyticNotFound) {
// 			uc.log.Error(err)
// 		}

// 		analytic, err := uc.repo.Get(ctx, postID)
// 		if err != nil {
// 			return nil, err
// 		}

// 		analyticWords := &entity.AnalyticWithWords{
// 			Analytic: analytic,
// 			Skills:   []*entity.TopWords{},
// 			Keywords: []*entity.TopWords{},
// 		}

// 		if err := uc.cache.SetWithWords(context.Background(), analyticWords, postID); err != nil {
// 			uc.log.Error(err)
// 		}

// 		return analyticWords, nil
// 	}

// 	// проверка len во избежания паники(а вдруг)
// 	if limit && len(analytic.Keywords) > 8 && len(analytic.Keywords) > 8 {
// 		// ограничиваем количество слов для анонимов
// 		analytic.Keywords = analytic.Keywords[:8]
// 		analytic.Skills = analytic.Skills[:8]
// 	}

// 	return analytic, nil
// }

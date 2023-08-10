package service

import (
	"context"
	"mall.com/store"
	"mall.com/store/models"
)

// CategorySrv defines functions used to handle Category request.
type CategorySrv interface {
	Create(ctx context.Context, category *models.WebCategoryCreateParam) uint64
	Update(ctx context.Context, category *models.WebCategoryUpdateParam) int64
	Delete(ctx context.Context, category *models.WebCategoryDeleteParam) int64
	//DeleteCollection(ctx context.Context, username string, CategoryIDs []string, opts metav1.DeleteOptions) error
	Get(ctx context.Context) []models.WebCategoryOption
	List(ctx context.Context, category *models.WebCategoryQueryParam) ([]models.WebCategoryList, int64)
}

type categoryService struct {
	store store.Factory
}

var _ CategorySrv = (*categoryService)(nil)

func newCategorys(srv *service) *categoryService {
	return &categoryService{store: srv.store}
}

func (s *categoryService) Create(ctx context.Context, category *models.WebCategoryCreateParam) uint64 {

	count := s.store.Categorys().Create(ctx, category)
	return count
}

func (s *categoryService) Update(ctx context.Context, category *models.WebCategoryUpdateParam) int64 {

	count := s.store.Categorys().Update(ctx, category)
	return count
}

func (s *categoryService) Delete(ctx context.Context, category *models.WebCategoryDeleteParam) int64 {

	count := s.store.Categorys().Delete(ctx, category)
	return count
}

func (s *categoryService) List(ctx context.Context, category *models.WebCategoryQueryParam) ([]models.WebCategoryList, int64) {

	cateList, rows := s.store.Categorys().List(ctx, category)
	return cateList, rows
}

func (s *categoryService) Get(ctx context.Context) []models.WebCategoryOption {

	selectList := s.store.Categorys().Get(ctx)
	return selectList
}

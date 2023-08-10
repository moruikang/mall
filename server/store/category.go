package store

import (
	"context"
	"mall.com/store/models"
)

// CategoryStore defines the Category storage interface.
type CategoryStore interface {
	Create(ctx context.Context, CategoryParm *models.WebCategoryCreateParam) uint64
	Update(ctx context.Context, CategoryParm *models.WebCategoryUpdateParam) int64
	Delete(ctx context.Context, CategoryParm *models.WebCategoryDeleteParam) int64
	//DeleteCollection(ctx context.Context, username string, CategoryIDs []string, opts metav1.DeleteOptions) error
	Get(ctx context.Context) (option []models.WebCategoryOption)
	List(ctx context.Context, CategoryParm *models.WebCategoryQueryParam) ([]models.WebCategoryList, int64)
}

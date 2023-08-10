package store

import (
	"context"
	"mall.com/store/models"
)

// CategoryStore defines the Category storage interface.
type ProductStore interface {
	Create(ctx context.Context, productParm *models.WebProductCreateParam) int64
	Update(ctx context.Context, productParm *models.WebProductUpdateParam) int64
	UpdateStatus(ctx context.Context, productParm *models.WebProductStatusUpdateParam) int64
	Delete(ctx context.Context, productParm *models.WebProductDeleteParam) int64
	GetInfo(ctx context.Context, productParm *models.WebProductInfoParam) models.WebProductInfo
	//DeleteCollection(ctx context.Context, username string, CategoryIDs []string, opts metav1.DeleteOptions) error
	//Get(ctx context.Context, username, CategoryID string, opts metav1.GetOptions) (*v1.Category, error)
	List(ctx context.Context, productParm *models.WebProductListParam) ([]models.WebProductList, int64)
}

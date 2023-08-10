package store

import (
	"context"
	"mall.com/store/models"
)

type OrderStore interface {
	//Create(ctx context.Context, pParm *models.WebOrderList ) int64
	Update(ctx context.Context, orderParam *models.WebOrderUpdateParam) int64
	//UpdateStatus(ctx context.Context, productParm *models.WebProductStatusUpdateParam ) int64
	Delete(ctx context.Context, orderParam *models.WebOrderDeleteParam) int64
	//GetInfo(ctx context.Context, productParm *models.WebProductInfoParam ) models.WebProductInfo
	//DeleteCollection(ctx context.Context, username string, CategoryIDs []string, opts metav1.DeleteOptions) error
	//Get(ctx context.Context, username, CategoryID string, opts metav1.GetOptions) (*v1.Category, error)
	List(ctx context.Context, orderParam *models.WebOrderListParam) ([]models.WebOrderList, int64)
	Detail(ctx context.Context, orderParam *models.WebOrderDetailParam) models.WebOrderDetail
}

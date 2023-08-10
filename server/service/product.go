package service

import (
	"context"
	"mall.com/store"
	"mall.com/store/models"
)

// ProductSrv defines functions used to handle Product request.
type ProductSrv interface {
	Create(ctx context.Context, product *models.WebProductCreateParam) int64
	Update(ctx context.Context, product *models.WebProductUpdateParam) int64
	UpdateStatus(ctx context.Context, product *models.WebProductStatusUpdateParam) int64
	Delete(ctx context.Context, product *models.WebProductDeleteParam) int64
	GetInfo(ctx context.Context, product *models.WebProductInfoParam) models.WebProductInfo
	//DeleteCollection(ctx context.Context, username string, ProductIDs []string, opts metav1.DeleteOptions) error
	//Get(ctx context.Context, username, ProductID string, opts metav1.GetOptions) (*v1.Product, error)
	List(ctx context.Context, product *models.WebProductListParam) ([]models.WebProductList, int64)
}

type productService struct {
	store store.Factory
}

var _ ProductSrv = (*productService)(nil)

func newProducts(srv *service) *productService {
	return &productService{store: srv.store}
}

func (s *productService) Create(ctx context.Context, product *models.WebProductCreateParam) int64 {

	rows := s.store.Products().Create(ctx, product)
	return rows
}

func (s *productService) Delete(ctx context.Context, product *models.WebProductDeleteParam) int64 {

	rows := s.store.Products().Delete(ctx, product)
	return rows
}

func (s *productService) Update(ctx context.Context, product *models.WebProductUpdateParam) int64 {

	rows := s.store.Products().Update(ctx, product)
	return rows
}

func (s *productService) UpdateStatus(ctx context.Context, product *models.WebProductStatusUpdateParam) int64 {

	rows := s.store.Products().UpdateStatus(ctx, product)
	return rows
}

func (s *productService) GetInfo(ctx context.Context, product *models.WebProductInfoParam) models.WebProductInfo {

	info := s.store.Products().GetInfo(ctx, product)
	return info
}

func (s *productService) List(ctx context.Context, product *models.WebProductListParam) ([]models.WebProductList, int64) {
	productList, rows := s.store.Products().List(ctx, product)
	return productList, rows
}

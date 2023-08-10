package service

import (
	"context"
	"mall.com/store"
	"mall.com/store/models"
)

// OrderSrv defines functions used to handle Order request.
type OrderSrv interface {
	//Create(ctx context.Context, Order *models.WebOrderCreateParam) int64
	Update(ctx context.Context, order *models.WebOrderUpdateParam) int64
	//UpdateStatus(ctx context.Context, Order *models.WebOrderStatusUpdateParam) int64
	Delete(ctx context.Context, order *models.WebOrderDeleteParam) int64
	//GetInfo(ctx context.Context, Order *models.WebOrderInfoParam) models.WebOrderInfo
	//DeleteCollection(ctx context.Context, username string, OrderIDs []string, opts metav1.DeleteOptions) error
	//Get(ctx context.Context, username, OrderID string, opts metav1.GetOptions) (*v1.Order, error)
	List(ctx context.Context, order *models.WebOrderListParam) ([]models.WebOrderList, int64)
	Detail(ctx context.Context, order *models.WebOrderDetailParam) models.WebOrderDetail
}

type orderService struct {
	store store.Factory
}

var _ OrderSrv = (*orderService)(nil)

func newOrders(srv *service) *orderService {
	return &orderService{store: srv.store}
}

func (s *orderService) Delete(ctx context.Context, orderparam *models.WebOrderDeleteParam) int64 {

	rows := s.store.Orders().Delete(ctx, orderparam)
	return rows
}

func (s *orderService) Update(ctx context.Context, orderparam *models.WebOrderUpdateParam) int64 {

	rows := s.store.Orders().Update(ctx, orderparam)
	return rows
}

func (s *orderService) List(ctx context.Context, order *models.WebOrderListParam) ([]models.WebOrderList, int64) {

	orderList, rows := s.store.Orders().List(ctx, order)
	return orderList, rows
}

func (s *orderService) Detail(ctx context.Context, order *models.WebOrderDetailParam) models.WebOrderDetail {

	orderDetail := s.store.Orders().Detail(ctx, order)
	return orderDetail
}

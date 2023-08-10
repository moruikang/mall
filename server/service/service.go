package service

import "mall.com/store"

// Service defines functions used to return resource interface.

type Service interface {
	Categorys() CategorySrv
	Products() ProductSrv
	Orders() OrderSrv
}

type service struct {
	store store.Factory
}

// NewService returns Service interface.
func NewService(store store.Factory) Service {
	return &service{
		store: store,
	}
}

func (s *service) Categorys() CategorySrv {
	return newCategorys(s)
}

func (s *service) Products() ProductSrv {
	return newProducts(s)
}

func (s *service) Orders() OrderSrv {
	return newOrders(s)
}

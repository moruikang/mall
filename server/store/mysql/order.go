package mysql

import (
	"context"
	"gorm.io/gorm"
	"mall.com/pkg/common"
	"mall.com/store/models"
	"strconv"
	"strings"
)

type orders struct {
	db *gorm.DB
}

func newOrders(ds *datastore) *orders {
	return &orders{ds.db}
}

func (s *orders) Delete(ctx context.Context, param *models.WebOrderDeleteParam) int64 {

	return s.db.Delete(&models.Order{}, param.Id).RowsAffected

}

func (s *orders) Update(ctx context.Context, param *models.WebOrderUpdateParam) int64 {

	order := models.Order{
		Id:      param.Id,
		Status:  param.Status,
		Updated: common.NowTime(),
	}
	return s.db.Model(&order).Updates(order).RowsAffected

}

func (s *orders) List(ctx context.Context, param *models.WebOrderListParam) ([]models.WebOrderList, int64) {

	orderList := make([]models.WebOrderList, 0)
	query := &models.Order{
		Id:     param.Id,
		Status: param.Status,
	}
	rows := common.RestPage(param.Page, "order", query, &orderList, &[]models.Order{})
	return orderList, rows

}

func (s *orders) Detail(ctx context.Context, param *models.WebOrderDetailParam) models.WebOrderDetail {

	var order models.Order
	var address models.Address
	var productItem []models.WebProductItem

	// 查询订单信息与地址信息
	s.db.First(&order, param.Id)
	s.db.First(&address, order.AddressId)

	// 查询订单中包含的商品信息
	idList := strings.Split(order.ProductItem, ",")
	productIdList := make([]uint64, 0)
	for _, id := range idList {
		pid, _ := strconv.Atoi(id)
		if pid != 0 {
			productIdList = append(productIdList, uint64(pid))
		}
	}
	s.db.Table("product").Find(&productItem, productIdList)
	orderDetail := models.WebOrderDetail{
		Id:              order.Id,
		Created:         order.Created,
		NickName:        order.NickName,
		Status:          order.Status,
		TotalPrice:      order.TotalPrice,
		Name:            address.Name,
		Mobile:          address.Mobile,
		PostalCode:      address.PostalCode,
		Province:        address.Province,
		City:            address.City,
		District:        address.District,
		DetailedAddress: address.DetailedAddress,
		ProductItem:     productItem,
	}
	return orderDetail

}

package mysql

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"mall.com/config/global"
	"mall.com/pkg/common"
	"mall.com/store/models"
	"strconv"
)

type products struct {
	db *gorm.DB
}

func newProducts(ds *datastore) *products {
	return &products{ds.db}
}

// Create creates a new product.
func (s *products) Create(ctx context.Context, param *models.WebProductCreateParam) int64 {

	product := models.Product{
		CategoryId:        param.CategoryId,
		Title:             param.Title,
		Description:       param.Description,
		Price:             param.Price,
		Amount:            param.Amount,
		MainImage:         param.MainImage,
		Delivery:          param.Delivery,
		Assurance:         param.Assurance,
		Name:              param.Name,
		Weight:            param.Weight,
		Brand:             param.Brand,
		Origin:            param.Origin,
		ShelfLife:         param.ShelfLife,
		NetWeight:         param.NetWeight,
		UseWay:            param.UseWay,
		PackingWay:        param.PackingWay,
		StorageConditions: param.StorageConditions,
		DetailImage:       param.DetailImage,
		Status:            param.Status,
		Created:           common.NowTime(),
	}
	rows := s.db.Create(&product).RowsAffected
	records := s.db.First(&product, product.Id).RowsAffected
	if records > 0 {
		id := strconv.FormatUint(product.Id, 10)
		result, err := global.Es.Index().Index("product").Id(id).BodyJson(product).Do(ctx)
		if err != nil {
			fmt.Println(err)
		}
		return result.PrimaryTerm
	}
	return rows

}

func (s *products) Delete(ctx context.Context, param *models.WebProductDeleteParam) int64 {

	rows := s.db.Delete(&models.Product{}, param.Id).RowsAffected
	if rows > 0 {
		id := strconv.FormatUint(param.Id, 10)
		_, err := global.Es.Delete().Index("product").Id(id).Do(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	return rows

}

func (s *products) Update(ctx context.Context, param *models.WebProductUpdateParam) int64 {

	product := models.Product{
		Id:                param.Id,
		CategoryId:        param.CategoryId,
		Title:             param.Title,
		Description:       param.Description,
		Price:             param.Price,
		Amount:            param.Amount,
		MainImage:         param.MainImage,
		Delivery:          param.Delivery,
		Assurance:         param.Assurance,
		Name:              param.Name,
		Weight:            param.Weight,
		Brand:             param.Brand,
		Origin:            param.Origin,
		ShelfLife:         param.ShelfLife,
		NetWeight:         param.NetWeight,
		UseWay:            param.UseWay,
		PackingWay:        param.PackingWay,
		StorageConditions: param.StorageConditions,
		DetailImage:       param.DetailImage,
		Status:            param.Status,
		Updated:           common.NowTime(),
	}
	rows := s.db.Model(&product).Updates(product).RowsAffected
	records := s.db.First(&product, product.Id).RowsAffected
	if records > 0 {
		id := strconv.FormatUint(param.Id, 10)
		_, err := global.Es.Update().Index("product").Id(id).Doc(product).Do(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	return rows

}

func (s *products) UpdateStatus(ctx context.Context, param *models.WebProductStatusUpdateParam) int64 {

	product := models.Product{
		Id:     param.Id,
		Status: param.Status,
	}
	rows := s.db.Model(&product).Update("status", product.Status).RowsAffected
	records := s.db.First(&product, product.Id).RowsAffected
	if records > 0 {
		id := strconv.FormatUint(param.Id, 10)
		_, err := global.Es.Update().Index("product").Id(id).Doc(product).Do(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	return rows

}

func (s *products) GetInfo(ctx context.Context, param *models.WebProductInfoParam) models.WebProductInfo {

	var product models.WebProductInfo
	s.db.Table("product").First(&product, param.Id)
	return product

}

func (s *products) List(ctx context.Context, param *models.WebProductListParam) ([]models.WebProductList, int64) {

	query := &models.Product{
		Id:         param.Id,
		CategoryId: param.CategoryId,
		Title:      param.Title,
		Status:     param.Status,
	}
	productList := make([]models.WebProductList, 0)
	rows := common.RestPage(param.Page, "product", query, &productList, &[]models.Product{})
	return productList, rows

}

package mysql

import (
	"context"
	"gorm.io/gorm"
	"mall.com/pkg/common"
	"mall.com/store/models"
)

type categorys struct {
	db *gorm.DB
}

func newCategorys(ds *datastore) *categorys {
	return &categorys{ds.db}
}

// Create creates a new category.
func (s *categorys) Create(ctx context.Context, param *models.WebCategoryCreateParam) uint64 {

	var category models.Category
	result := s.db.Where("name = ?", param.Name).First(&category)
	if result.RowsAffected > 0 {
		return category.Id
	}
	category = models.Category{
		Name:     param.Name,
		ParentId: param.ParentId,
		Level:    param.Level,
		Sort:     param.Sort,
		Created:  common.NowTime(),
	}
	s.db.Create(&category)
	return category.Id

}

// Delete creates a new category.
func (s *categorys) Delete(ctx context.Context, param *models.WebCategoryDeleteParam) int64 {

	var pid2, pid3 models.Category
	s.db.Where("parent_id = ?", param.Id).First(&pid2)
	s.db.Where("parent_id = ?", pid2.Id).First(&pid3)
	return s.db.Delete(&models.Category{}, []uint64{param.Id, pid2.Id, pid3.Id}).RowsAffected

}

func (s *categorys) Update(ctx context.Context, param *models.WebCategoryUpdateParam) int64 {

	category := models.Category{
		Id:      param.Id,
		Name:    param.Name,
		Sort:    param.Sort,
		Updated: common.NowTime(),
	}
	return s.db.Model(&category).Updates(category).RowsAffected

}

func (s *categorys) List(ctx context.Context, param *models.WebCategoryQueryParam) ([]models.WebCategoryList, int64) {

	categoryList := make([]models.WebCategoryList, 0)
	query := &models.Category{
		Id:       param.Id,
		Name:     param.Name,
		Level:    param.Level,
		ParentId: param.ParentId,
	}
	rows := common.RestPage(param.Page, "category", query, &categoryList, &[]models.Category{})
	return categoryList, rows

}

func (s *categorys) Get(ctx context.Context) (option []models.WebCategoryOption) {

	selectList := make([]models.WebCategoryList, 0)
	s.db.Table("category").Find(&selectList)
	return getTreeOptions(1, selectList)

}

// 获取树形结构的选项
func getTreeOptions(id uint64, cateList []models.WebCategoryList) (option []models.WebCategoryOption) {
	optionList := make([]models.WebCategoryOption, 0)
	for _, opt := range cateList {
		if opt.ParentId == id && (opt.Level == 1 || opt.Level == 2) {
			option := models.WebCategoryOption{
				Value:    opt.Id,
				Label:    opt.Name,
				Children: getTreeOptions(opt.Id, cateList),
			}
			if opt.Level == 2 {
				option.Children = nil
			}
			optionList = append(optionList, option)
		}
	}
	return optionList
}

package product

import (
	"github.com/gin-gonic/gin"
	"mall.com/pkg/response"
	srvv1 "mall.com/service"
	"mall.com/store"
	"mall.com/store/models"
)

// ProductController create a Product handler used to handle request for Product resource.
type ProductController struct {
	srv srvv1.Service
}

// NewProductController creates a Product handler.
func NewProductController(store store.Factory) *ProductController {
	return &ProductController{
		srv: srvv1.NewService(store),
	}
}

// Create add new Product key pairs to the storage.
func (s *ProductController) Create(c *gin.Context) {

	var r models.WebProductCreateParam

	if err := c.ShouldBindJSON(&r); err != nil {
		response.Failed("请求参数无效", c)
		return

	}

	if rows := s.srv.Products().Create(c, &r); rows > 0 {
		response.Success("创建成功", rows, c)
		return
	}
	response.Failed("创建失败", c)
}

func (s *ProductController) Delete(c *gin.Context) {

	var r models.WebProductDeleteParam

	if err := c.ShouldBindJSON(&r); err != nil {
		response.Failed("请求参数无效", c)
		return

	}

	if rows := s.srv.Products().Delete(c, &r); rows > 0 {
		response.Success("删除成功", rows, c)
		return
	}
	response.Failed("删除失败", c)
}

func (s *ProductController) Update(c *gin.Context) {

	var r models.WebProductUpdateParam

	if err := c.ShouldBindJSON(&r); err != nil {
		response.Failed("请求参数无效", c)
		return

	}

	if rows := s.srv.Products().Update(c, &r); rows > 0 {
		response.Success("更新成功", rows, c)
		return
	}
	response.Failed("更新失败", c)
}

func (s *ProductController) UpdateStatus(c *gin.Context) {

	var r models.WebProductStatusUpdateParam

	if err := c.ShouldBindJSON(&r); err != nil {
		response.Failed("请求参数无效", c)
		return

	}

	if rows := s.srv.Products().UpdateStatus(c, &r); rows > 0 {
		response.Success("更新成功", rows, c)
		return
	}
	response.Failed("更新失败", c)
}

func (s *ProductController) GetInfo(c *gin.Context) {

	var r models.WebProductInfoParam

	if err := c.ShouldBindJSON(&r); err != nil {
		response.Failed("请求参数无效", c)
		return

	}

	productInfo := s.srv.Products().GetInfo(c, &r)
	response.Success("查询成功", productInfo, c)
	return
}

func (s *ProductController) List(c *gin.Context) {

	var r models.WebProductListParam

	if err := c.ShouldBind(&r); err != nil {
		response.Failed("请求参数无效", c)
		return

	}

	productList, rows := s.srv.Products().List(c, &r)
	response.SuccessPage("查询成功", productList, rows, c)
}

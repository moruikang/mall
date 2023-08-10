package order

import (
	"github.com/gin-gonic/gin"
	"mall.com/pkg/response"
	srvv1 "mall.com/service"
	"mall.com/store"
	"mall.com/store/models"
)

type OrderController struct {
	srv srvv1.Service
}

// NewOrderController creates a Order handler.
func NewOrderController(store store.Factory) *OrderController {
	return &OrderController{
		srv: srvv1.NewService(store),
	}
}

// Create add new Order key pairs to the storage.
func (s *OrderController) Delete(c *gin.Context) {

	var r models.WebOrderDeleteParam

	if err := c.ShouldBindJSON(&r); err != nil {
		response.Failed("请求参数无效", c)
		return

	}

	if rows := s.srv.Orders().Delete(c, &r); rows > 0 {
		response.Success("创建成功", rows, c)
		return
	}
	response.Failed("创建失败", c)
}

func (s *OrderController) Update(c *gin.Context) {

	var r models.WebOrderUpdateParam

	if err := c.ShouldBindJSON(&r); err != nil {
		response.Failed("请求参数无效", c)
		return

	}

	if rows := s.srv.Orders().Update(c, &r); rows > 0 {
		response.Success("创建成功", rows, c)
		return
	}
	response.Failed("创建失败", c)
}

func (s *OrderController) List(c *gin.Context) {

	var r models.WebOrderListParam

	if err := c.ShouldBindJSON(&r); err != nil {
		response.Failed("请求参数无效", c)
		return

	}

	productList, rows := s.srv.Orders().List(c, &r)
	response.SuccessPage("查询成功", productList, rows, c)
}

func (s *OrderController) Detail(c *gin.Context) {

	var r models.WebOrderDetailParam

	if err := c.ShouldBindJSON(&r); err != nil {
		response.Failed("请求参数无效", c)
		return

	}

	productDetail := s.srv.Orders().Detail(c, &r)
	response.Success("查询成功", productDetail, c)
}

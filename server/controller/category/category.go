package category

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mall.com/pkg/response"
	srvv1 "mall.com/service"
	"mall.com/store"
	"mall.com/store/models"
)

// CategoryController create a Category handler used to handle request for Category resource.
type CategoryController struct {
	srv srvv1.Service
}

// NewCategoryController creates a Category handler.
func NewCategoryController(store store.Factory) *CategoryController {
	return &CategoryController{
		srv: srvv1.NewService(store),
	}
}

// Create add new Category key pairs to the storage.
func (s *CategoryController) Create(c *gin.Context) {

	var r models.WebCategoryCreateParam

	if err := c.ShouldBindJSON(&r); err != nil {
		response.Failed("请求参数无效", c)
		return

	}

	fmt.Println(r)
	if count := s.srv.Categorys().Create(c, &r); count > 0 {
		response.Success("创建成功", count, c)
		return
	}
	response.Failed("创建失败", c)
}

func (s *CategoryController) Delete(c *gin.Context) {

	var r models.WebCategoryDeleteParam

	if err := c.ShouldBindJSON(&r); err != nil {
		response.Failed("请求参数无效", c)
		return

	}

	if count := s.srv.Categorys().Delete(c, &r); count > 0 {
		response.Success("删除成功", count, c)
		return
	}
	response.Failed("删除失败", c)
}

func (s *CategoryController) Update(c *gin.Context) {

	var r models.WebCategoryUpdateParam

	if err := c.ShouldBindJSON(&r); err != nil {
		response.Failed("请求参数无效", c)
		return

	}

	if count := s.srv.Categorys().Update(c, &r); count > 0 {
		response.Success("更新成功", count, c)
		return
	}
	response.Failed("更新失败", c)
}

func (s *CategoryController) List(c *gin.Context) {

	var param models.WebCategoryQueryParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	cateList, rows := s.srv.Categorys().List(c, &param)
	response.SuccessPage("查询成功", cateList, rows, c)
}

func (s *CategoryController) Get(c *gin.Context) {

	option := s.srv.Categorys().Get(c)
	response.Success("查询成功", option, c)
}

package routers

import (
	"github.com/gin-gonic/gin"
	"log"
	"mall.com/config/global"
	"mall.com/controller/api"
	"mall.com/controller/category"
	"mall.com/controller/order"
	"mall.com/controller/product"
	"mall.com/pkg/middleware"
	"mall.com/store/mysql"
)

func Router() *gin.Engine {

	engine := gin.Default()

	// 开启跨域
	engine.Use(middleware.Cors())

	// 静态资源请求映射
	engine.Static("/image", global.Config.Upload.SavePath)
	//测试
	//engine.GET("/ping", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "pong",
	//	})
	//})

	storeIns, _ := mysql.GetMySQLFactoryOr()

	authMiddleware, err := middleware.NewGinJwtMiddlewares()
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	log.Println(authMiddleware)
	// 后台管理员前端接口
	web := engine.Group("/web")

	{

		web.GET("/captcha", api.WebGetCaptcha)
		// 用户登录API

		//web.POST("/login", api.WebUserLogin)
		//// 开启JWT认证,以下接口需要认证成功才能访问
		//web.Use(middleware.JwtAuth())

		//gin-jwt
		web.POST("/login", authMiddleware.LoginHandler)
		web.Use(authMiddleware.MiddlewareFunc())

		// 文件上传API
		web.POST("/upload", api.WebFileUpload)

		// 类目管理API
		webcategory := web.Group("/category")
		{
			categoryController := category.NewCategoryController(storeIns)

			webcategory.POST("/create", categoryController.Create)
			webcategory.DELETE("/delete", categoryController.Delete)
			webcategory.PUT("/update", categoryController.Update)
			webcategory.GET("/list", categoryController.List)
			webcategory.GET("/option", categoryController.Get)
		}

		webproduct := web.Group("/product")
		{
			productController := product.NewProductController(storeIns)

			webproduct.POST("/create", productController.Create)
			webproduct.DELETE("/delete", productController.Delete)
			webproduct.PUT("/update", productController.Update)
			webproduct.PUT("/status/update", productController.UpdateStatus)
			webproduct.GET("/info", productController.GetInfo)
			webproduct.GET("/list", productController.List)
		}

		weborder := web.Group("/order")
		{
			orderController := order.NewOrderController(storeIns)

			weborder.DELETE("/delete", orderController.Delete)
			web.PUT("/update", orderController.Update)
			web.GET("/list", orderController.List)
			web.GET("/detail", orderController.Detail)
		}

		// 数据统计API
		web.GET("/data/overview/info", api.WebGetDataOverviewInfo)
		web.GET("/today/order/data/info", api.WebGetTodayOrderDataInfo)
		web.GET("/week/data/info", api.WebGetWeekDataInfo)

	}

	// 启动、监听端口
	//post := fmt.Sprintf(":%s", global.Config.Server.Post)
	//if err := engine.Run(post); err != nil {
	//	fmt.Printf("server start error: %s", err)
	//}

	return engine

}

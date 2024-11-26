package router

import (
	"ChainMaker_Backend_demo/src/ctrl"
	//_ "ChainMaker_Backend_demo/src/docs"
	"github.com/gin-gonic/gin"

	"net/http"
)

func HttpServe() {
	// 启动Web服务(默认Debug级别)
	gin.SetMode(gin.ReleaseMode)
	// 生成route
	ginRouter := gin.Default()
	ginRouter.Use(CorsMiddleware())

	//ginRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 初始化路由配置
	InitRouter(ginRouter)
	// 启动Http服务
	err := ginRouter.Run(":8087")
	if err != nil {
		panic(err)
	}
}

func InitRouter(router *gin.Engine) {
	group := router.Group("/")
	initControllers(group) // 定义接口
}

func initControllers(routeGroup *gin.RouterGroup) {

	routeGroup.POST("chainmaker", ctrl.Dispatcher)

}

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 允许跨域请求的来源，可以设置为具体的域名或通配符 "*" 表示允许所有来源
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		// 允许的方法，如GET、POST、PUT、DELETE等
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		// 允许的头部字段，如Content-Type、Authorization等
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// 允许预检请求的有效期，单位为秒
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")

		// 如果是OPTIONS请求，直接返回，不继续处理后续中间件或处理程序
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		} else {
			c.Next() // 继续执行后续中间件或处理程序
		}
	}
}

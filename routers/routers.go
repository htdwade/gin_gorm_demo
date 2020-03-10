package routers

import (
	"gin_gorm_demo/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 静态资源映射
	r.Static("/static", "./static")

	// 加载模板文件
	r.LoadHTMLGlob("./templates/*")

	r.GET("/", controller.IndexHandler)

	// v1路由组
	v1Group := r.Group("/v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", controller.CreateTodo)
		// 查看所有的待办事项
		v1Group.GET("/todo", controller.GetTodoList)
		// 查看某一个待办事项
		v1Group.GET("/todo:id", controller.GetTodo)
		// 修改某一个待办事项
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		// 删除某一个待办事项
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "page not found",
		})
	})
	return r
}

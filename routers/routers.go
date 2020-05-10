package routers

import (
	"github.com/gin-gonic/gin"
	"k-board/controllers"
)

func Setup(api *gin.Engine) {
	namespaces := api.Group("/namespaces")
	{
		namespaces.GET("", controllers.ListNamespaces)
	}
	pods := api.Group("/pods")
	{
		pods.GET("", controllers.ListPods)
	}
}

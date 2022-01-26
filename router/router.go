package router

import (
	"WebVisitor/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadRouter(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// 加载中间件
	g.Use(mw...)

	// 404
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	api := g.Group("/api")
	apiV1 := api.Group("/v1")
	{
		apiV1.GET("/", controllers.TestV1)
	}
	g.GET("/visitor", controllers.GetVisitorInfo)
	return g
}

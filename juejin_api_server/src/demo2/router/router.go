package router

import (
	"demo2/handler/sd"
	"demo2/router/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrent API route")
	})

	scvd := g.Group("/sd")
	{
		scvd.GET("/health", sd.HealCheck)
		scvd.GET("/disk", sd.DiskCheck)
		scvd.GET("/cpu", sd.CPUCheck)
		scvd.GET("/ram", sd.RAMCheck)
	}

	return g
}

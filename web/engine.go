package web

import (
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"gopkg.in/gin-gonic/gin.v1"
)

func GetMainEngine(ol officeLocatorInterface) *gin.Engine {
	router := gin.Default()
	addCORSIfEnabled(router)

	lh := locationHandler{locator: ol}
	v1 := router.Group("/v1")
	{
		v1.GET("/nearest", lh.handleNearest)
	}

	return router
}

func addCORSIfEnabled(router *gin.Engine) {
	rawOrigins := os.Getenv("CORS_ALLOWED_ORIGINS")
	if len(rawOrigins) == 0 {
		return
	}

	origins := strings.Split(rawOrigins, ",")
	config := cors.DefaultConfig()
	config.AllowOrigins = origins
	router.Use(cors.New(config))
}

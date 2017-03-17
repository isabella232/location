package web

import (
	"net/url"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"gopkg.in/gin-gonic/gin.v1"
)

func GetMainEngine(ol officeLocatorInterface) *gin.Engine {
	router := gin.Default()
	addCORSIfEnabled(router)

	lh := locationHandler{
		locator:       ol,
		thoughtbotURL: thoughtbotURL(),
	}

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

func thoughtbotURL() url.URL {
	rawURL := os.Getenv("THOUGHTBOT_URL")
	if len(rawURL) == 0 {
		panic("Expected THOUGHTBOT_URL env variable required")
	}

	u, _ := url.Parse(rawURL)
	return *u
}

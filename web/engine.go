package web

import "github.com/gin-gonic/gin"

func GetMainEngine(ol officeLocatorInterface) *gin.Engine {
	router := gin.Default()

	lh := locationHandler{locator: ol}
	v1 := router.Group("/v1")
	{
		v1.GET("/nearest", lh.handleNearest)
	}

	return router
}

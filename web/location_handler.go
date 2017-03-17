package web

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

//go:generate counterfeiter . officeLocatorInterface
type officeLocatorInterface interface {
	Nearest(ipAddress string) (slug string, distanceKm float64, err error)
}

type locationHandler struct {
	locator officeLocatorInterface
}

func (h *locationHandler) handleNearest(c *gin.Context) {
	location, distanceKm, err := h.locator.Nearest(c.ClientIP())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"slug": location,
		"meta": gin.H{
			"distanceKmToUser": distanceKm,
		},
	})
}

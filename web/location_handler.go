package web

import (
	"net/http"
	"net/url"

	"github.com/thoughtbot/location/locator"

	"gopkg.in/gin-gonic/gin.v1"
)

//go:generate counterfeiter . officeLocatorInterface
type officeLocatorInterface interface {
	Nearest(ipAddress string) (nearestOffice locator.Office, distanceKm float64, err error)
}

type locationHandler struct {
	locator       officeLocatorInterface
	thoughtbotURL url.URL
}

func (h *locationHandler) handleNearest(c *gin.Context) {
	o, distanceKm, err := h.locator.Nearest(c.ClientIP())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	officeURL := o.URL(h.thoughtbotURL)

	c.JSON(http.StatusOK, gin.H{
		"slug": o.Slug,
		"name": o.Name,
		"url":  officeURL.String(),
		"meta": gin.H{
			"distanceKmToUser": distanceKm,
		},
	})
}

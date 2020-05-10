package controllers

import (
	"github.com/gin-gonic/gin"
	"k-board/services"
	"k-board/services/namespace"
	"net/http"
)

func ListNamespaces(c *gin.Context) {
	namespaceSvc := c.MustGet(services.Namespace).(namespace.Service)
	ns, err := namespaceSvc.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"response": ns,
	})
}

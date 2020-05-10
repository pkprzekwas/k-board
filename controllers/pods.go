package controllers

import (
	"github.com/gin-gonic/gin"
	"k-board/services"
	"k-board/services/pod"
	"net/http"
)

func ListPods(c *gin.Context) {
	podSvc := c.MustGet(services.Pod).(pod.Service)
	namespace := c.DefaultQuery("namespace", "default")

	pods, err := podSvc.List(namespace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"response": pods,
	})
}

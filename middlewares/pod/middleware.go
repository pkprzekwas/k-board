package pod

import (
	"github.com/gin-gonic/gin"
	"k-board/services"
	"k-board/services/pod"
	"k8s.io/client-go/kubernetes"
)

func Middleware(client *kubernetes.Clientset) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(services.Pod, pod.Service{
			Client: client.CoreV1(),
		})
	}
}

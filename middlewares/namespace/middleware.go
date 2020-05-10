package namespace

import (
	"github.com/gin-gonic/gin"
	"k-board/services"
	"k-board/services/namespace"
	"k8s.io/client-go/kubernetes"
)

// Injects namespace service to the gin context
func Middleware(client *kubernetes.Clientset) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(services.Namespace, namespace.Service{
			Client: client.CoreV1().Namespaces(),
		})
	}
}

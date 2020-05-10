package main

import (
	"github.com/gin-gonic/gin"
	"k-board/middlewares/namespace"
	"k-board/middlewares/pod"
	"k-board/routers"
	"k-board/utils"
	"log"
)

func main() {
	api := gin.New()

	kubeclient, err := utils.CreateKubeClient()
	if err != nil {
		log.Fatal("couldn't load kubeconfig")
	}

	api.Use(namespace.Middleware(kubeclient))
	api.Use(pod.Middleware(kubeclient))

	routers.Setup(api)

	if err := api.Run(":8080"); err != nil {
		panic("couldn't start the server.")
	}
}

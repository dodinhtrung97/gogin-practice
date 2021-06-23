package main

import (
	"gogin-practice/controller"
	"gogin-practice/service"

	"github.com/gin-gonic/gin"
)

var (
	animalService    service.AnimalService       = service.NewAnimalService()
	animalController controller.AnimalController = controller.NewAnimalController(animalService)
)

func main() {
	v1 := gin.Default()

	v1.POST("/post", func(ctx *gin.Context) {
		ctx.JSON(200, animalController.Save(ctx))
		return
	})

	v1.GET("/get", func(ctx *gin.Context) {
		ctx.JSON(200, animalController.FindAll())
		return
	})

	v1.Run(":8080")
}

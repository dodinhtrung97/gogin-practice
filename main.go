package main

import (
	"gogin-practice/controller"
	"gogin-practice/repository"
	"gogin-practice/service"

	"github.com/gin-gonic/gin"
)

var (
	animalRepository repository.AnimalRepository = repository.NewAnimalRepository()
	animalService    service.AnimalService       = service.NewAnimalService(animalRepository)
	animalController controller.AnimalController = controller.NewAnimalController(animalService)
)

func main() {
	v1 := gin.Default()

	defer animalRepository.Close()

	v1.POST("/post", func(ctx *gin.Context) {
		ctx.JSON(200, animalController.Save(ctx))
	})

	v1.PUT("/put", func(ctx *gin.Context) {
		ctx.JSON(200, animalController.Update(ctx))
	})

	v1.DELETE("/delete", func(ctx *gin.Context) {
		ctx.JSON(200, animalController.Delete(ctx))
	})

	v1.GET("/get/:id", func(ctx *gin.Context) {
		result, err := animalController.FindById(ctx)

		if err != nil {
			ctx.JSON(400, err.Error)
		}
		ctx.JSON(200, result)
	})

	v1.GET("/get/", func(ctx *gin.Context) {
		ctx.JSON(200, animalController.FindAll())
	})
	v1.Run(":8080")
}

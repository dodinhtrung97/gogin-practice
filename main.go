package main

import (
	"gogin-practice/config"
	"gogin-practice/controller"
	"gogin-practice/repository"
	"gogin-practice/service"

	"github.com/gin-gonic/gin"
)

var (
	mysqlConnection  *config.Database            = config.NewMySqlConnection()
	animalRepository repository.AnimalRepository = repository.NewAnimalRepository(mysqlConnection)
	animalService    service.AnimalService       = service.NewAnimalService(animalRepository)
	animalController controller.AnimalController = controller.NewAnimalController(animalService)
)

func main() {
	v1 := gin.Default()

	defer mysqlConnection.Close()

	v1.POST("/post", animalController.Save)
	v1.PUT("/put", animalController.Update)
	v1.DELETE("/delete", animalController.Delete)
	v1.GET("/get/:id", animalController.FindById)
	v1.GET("/get/", animalController.FindAll)

	v1.Run(":8080")
}

package controller

import (
	"fmt"
	"gogin-practice/entity"
	"gogin-practice/service"

	"github.com/gin-gonic/gin"
)

type AnimalController interface {
	Save(ctx *gin.Context) entity.Animal
	FindAll() []entity.Animal
}

type animalController struct {
	service service.AnimalService
}

func NewAnimalController(service service.AnimalService) AnimalController {
	return &animalController{
		service: service,
	}
}

func (controller *animalController) Save(ctx *gin.Context) entity.Animal {
	var animal entity.Animal
	err := ctx.BindJSON(&animal)

	if err != nil {
		fmt.Println("Can't bind request body to Animal")
		return entity.Animal{}
	}
	return controller.service.Save(animal)
}

func (controller *animalController) FindAll() []entity.Animal {
	return controller.service.FindAll()
}

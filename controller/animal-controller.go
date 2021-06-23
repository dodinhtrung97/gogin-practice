package controller

import (
	"errors"
	"fmt"
	"gogin-practice/entity"
	"gogin-practice/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AnimalController interface {
	Save(ctx *gin.Context) entity.Animal
	Update(ctx *gin.Context) entity.Animal
	Delete(ctx *gin.Context) entity.Animal
	FindById(ctx *gin.Context) (entity.Animal, error)
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

func (controller *animalController) Update(ctx *gin.Context) entity.Animal {
	var animal entity.Animal
	err := ctx.BindJSON(&animal)

	if err != nil {
		fmt.Println("Can't bind request body to Animal")
		return entity.Animal{}
	}
	return controller.service.Update(animal)
}

func (controller *animalController) Delete(ctx *gin.Context) entity.Animal {
	var animal entity.Animal
	err := ctx.BindJSON(&animal)

	if err != nil {
		fmt.Println("Can't bind request body to Animal")
		return entity.Animal{}
	}
	return controller.service.Delete(animal)
}

func (controller *animalController) FindById(ctx *gin.Context) (entity.Animal, error) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		return entity.Animal{}, errors.New("Failed to parse input to id type")
	}
	return controller.service.FindById(id), nil
}

func (controller *animalController) FindAll() []entity.Animal {
	return controller.service.FindAll()
}

package controller

import (
	"gogin-practice/entity"
	"gogin-practice/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AnimalController interface {
	Save(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindById(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}

type animalController struct {
	service service.AnimalService
}

func NewAnimalController(service service.AnimalService) AnimalController {
	return &animalController{
		service: service,
	}
}

func (controller *animalController) Save(ctx *gin.Context) {
	var animal entity.Animal
	err := ctx.BindJSON(&animal)

	if err != nil {
		ctx.JSON(400, "Failed to bind request body to Animal")
		return
	}

	ctx.JSON(201, controller.service.Save(animal))
}

func (controller *animalController) Update(ctx *gin.Context) {
	var animal entity.Animal
	err := ctx.BindJSON(&animal)

	if err != nil {
		ctx.JSON(400, "Failed to bind request body to Animal")
		return
	}
	ctx.JSON(200, controller.service.Update(animal))
}

func (controller *animalController) Delete(ctx *gin.Context) {
	var animal entity.Animal
	err := ctx.BindJSON(&animal)

	if err != nil {
		ctx.JSON(400, "Failed to bind request body to Animal")
		return
	}
	ctx.JSON(200, controller.service.Delete(animal))
}

func (controller *animalController) FindById(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(400, "Failed to parse input to id type")
		return
	}
	ctx.JSON(200, controller.service.FindById(id))
}

func (controller *animalController) FindAll(ctx *gin.Context) {
	ctx.JSON(200, controller.service.FindAll())
}

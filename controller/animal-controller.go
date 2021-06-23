package controller

import (
	"gogin-practice/entity"
	"gogin-practice/error"
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

	if err := ctx.BindJSON(&animal); err != nil {
		ctx.JSON(400, "Failed to bind request body to Animal")
		return
	}

	result, err := controller.service.Save(animal)
	if err != nil {
		httpErr := error.NewBadRequestError("Failed to save new animal")
		ctx.JSON(httpErr.Status, httpErr)
		return
	}

	ctx.JSON(201, result)
}

func (controller *animalController) Update(ctx *gin.Context) {
	var animal entity.Animal

	if err := ctx.BindJSON(&animal); err != nil {
		ctx.JSON(400, "Failed to bind request body to Animal")
		return
	}

	result, err := controller.service.Update(animal)

	if err != nil {
		httpErr := error.NewNotFoundError("No existing entry to update")
		ctx.JSON(httpErr.Status, httpErr)
		return
	}
	ctx.JSON(200, result)
}

func (controller *animalController) Delete(ctx *gin.Context) {
	var animal entity.Animal

	if err := ctx.BindJSON(&animal); err != nil {
		ctx.JSON(400, "Failed to bind request body to Animal")
		return
	}

	result, err := controller.service.Delete(animal)

	if err != nil {
		httpErr := error.NewNotFoundError("No entry to delete")
		ctx.JSON(httpErr.Status, httpErr)
		return
	}
	ctx.JSON(200, result)
}

func (controller *animalController) FindById(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(400, "Failed to parse input to id type")
		return
	}

	animal, err := controller.service.FindById(id)
	if err != nil {
		httpErr := error.NewNotFoundError("Can't find entry with id " + strconv.FormatUint(id, 10))
		ctx.JSON(httpErr.Status, httpErr)
		return
	}
	ctx.JSON(200, animal)
}

func (controller *animalController) FindAll(ctx *gin.Context) {
	ctx.JSON(200, controller.service.FindAll())
}

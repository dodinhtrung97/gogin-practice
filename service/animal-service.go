package service

import (
	entity "gogin-practice/entity"
)

type AnimalService interface {
	Save(entity.Animal) entity.Animal
	FindAll() []entity.Animal
}

type animalService struct {
	animals []entity.Animal
}

func NewAnimalService() AnimalService {
	return &animalService{}
}

func (service *animalService) Save(animal entity.Animal) entity.Animal {
	service.animals = append(service.animals, animal)
	return animal
}

func (service *animalService) FindAll() []entity.Animal {
	return service.animals
}

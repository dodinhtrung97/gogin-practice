package service

import (
	entity "gogin-practice/entity"
	repository "gogin-practice/repository"
)

type AnimalService interface {
	Save(entity.Animal) (entity.Animal, error)
	Update(entity.Animal) (entity.Animal, error)
	Delete(entity.Animal) (entity.Animal, error)
	FindById(id uint64) (entity.Animal, error)
	FindAll() []entity.Animal
}

type animalService struct {
	animalRepository repository.AnimalRepository
}

func NewAnimalService(animalRepository repository.AnimalRepository) AnimalService {
	return &animalService{
		animalRepository: animalRepository,
	}
}

func (service *animalService) Save(animal entity.Animal) (entity.Animal, error) {
	return service.animalRepository.Save(animal)
}

func (service *animalService) Update(animal entity.Animal) (entity.Animal, error) {
	return service.animalRepository.Update(animal)
}

func (service *animalService) Delete(animal entity.Animal) (entity.Animal, error) {
	return service.animalRepository.Delete(animal)
}

func (service *animalService) FindById(id uint64) (entity.Animal, error) {
	return service.animalRepository.FindById(id)
}

func (service *animalService) FindAll() []entity.Animal {
	return service.animalRepository.FindAll()
}

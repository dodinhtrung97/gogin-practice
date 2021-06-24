package repository

import (
	"gogin-practice/config"
	"gogin-practice/entity"
)

type AnimalRepository interface {
	Close()
	Save(animal entity.Animal) (entity.Animal, error)
	Update(animal entity.Animal) (entity.Animal, error)
	Delete(animal entity.Animal) (entity.Animal, error)
	FindById(id uint64) (entity.Animal, error)
	FindAll() []entity.Animal
}

type animalRepository struct {
	db config.Database
}

func NewAnimalRepository(db *config.Database) AnimalRepository {
	return &animalRepository{
		db: *db,
	}
}

func (animalRepository *animalRepository) Close() {
	sqlDB, err := animalRepository.db.Connection.DB()
	if err != nil {
		panic("Can't access database instance")
	}
	sqlDB.Close()
}

func (animalRepository *animalRepository) Save(animal entity.Animal) (entity.Animal, error) {
	if err := animalRepository.db.Connection.Save(&animal).Error; err != nil {
		return entity.Animal{}, err
	}
	return animal, nil
}

func (animalRepository *animalRepository) Update(animal entity.Animal) (entity.Animal, error) {
	if err := animalRepository.db.Connection.Save(&animal).Error; err != nil {
		return entity.Animal{}, err
	}
	return animal, nil
}

func (animalRepository *animalRepository) Delete(animal entity.Animal) (entity.Animal, error) {
	if err := animalRepository.db.Connection.Delete(&animal).Error; err != nil {
		return entity.Animal{}, err
	}
	return animal, nil
}

func (animalRepository *animalRepository) FindById(id uint64) (entity.Animal, error) {
	var animal entity.Animal
	if err := animalRepository.db.Connection.Preload("Household").First(&animal, id).Error; err != nil {
		return entity.Animal{}, err
	}

	return animal, nil
}

func (animalRepository *animalRepository) FindAll() []entity.Animal {
	var animals []entity.Animal
	animalRepository.db.Connection.Preload("Household").Find(&animals)

	return animals
}

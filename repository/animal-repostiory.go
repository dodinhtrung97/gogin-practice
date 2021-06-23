package repository

import (
	"gogin-practice/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type AnimalRepository interface {
	Save(animal entity.Animal) entity.Animal
	Update(animal entity.Animal) entity.Animal
	Delete(animal entity.Animal) entity.Animal
	FindById(id uint64) entity.Animal
	FindAll() []entity.Animal
}

type database struct {
	connection *gorm.DB
}

func NewAnimalRepository() AnimalRepository {
	mysql_endpoint := "root:root@tcp(127.0.0.1:3306)/gogin_practice?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(mysql_endpoint), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to db @ " + mysql_endpoint)
	}

	db.AutoMigrate(&entity.Animal{})
	return &database{
		connection: db,
	}
}

func (db *database) Close() {
	sqlDB, err := db.connection.DB()
	if err != nil {
		panic("Can't access database instance")
	}
	sqlDB.Close()
}

func (db *database) Save(animal entity.Animal) entity.Animal {
	db.connection.Save(&animal)
	return animal
}

func (db *database) Update(animal entity.Animal) entity.Animal {
	db.connection.Save(&animal)
	return animal
}

func (db *database) Delete(animal entity.Animal) entity.Animal {
	db.connection.Delete(&animal)
	return animal
}

func (db *database) FindById(id uint64) entity.Animal {
	var animal entity.Animal
	db.connection.First(&animal, id)

	return animal
}

func (db *database) FindAll() []entity.Animal {
	var animals []entity.Animal
	db.connection.Find(&animals)

	return animals
}

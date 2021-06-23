package repository

import (
	"fmt"
	"gogin-practice/entity"
	"path"
	"runtime"
	"strings"

	"path/filepath"

	"github.com/tkanos/gonfig"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type AnimalRepository interface {
	Close()
	Save(animal entity.Animal) entity.Animal
	Update(animal entity.Animal) entity.Animal
	Delete(animal entity.Animal) entity.Animal
	FindById(id uint64) entity.Animal
	FindAll() []entity.Animal
}

type databaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type database struct {
	connection *gorm.DB
}

func NewAnimalRepository() AnimalRepository {
	var db_config databaseConfig
	load_config_err := gonfig.GetConf(getDbConfigFilePath(), &db_config)

	if load_config_err != nil {
		panic(load_config_err.Error())
	}

	mysql_endpoint := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		db_config.User,
		db_config.Password,
		db_config.Host,
		db_config.Port,
		db_config.Database)
	db, err := gorm.Open(mysql.Open(mysql_endpoint), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to db @ " + mysql_endpoint + " with error: " + err.Error())
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

func getDbConfigFilePath() string {
	filename := []string{"config/", "db_config", ".json"}
	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), "..", strings.Join(filename, ""))

	return filepath.FromSlash(filePath)
}
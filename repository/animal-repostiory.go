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
	Save(animal entity.Animal) (entity.Animal, error)
	Update(animal entity.Animal) (entity.Animal, error)
	Delete(animal entity.Animal) (entity.Animal, error)
	FindById(id uint64) (entity.Animal, error)
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
	dbConfig := getDbConfig()
	mysqlEndpoint := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Database)

	db, connectionErr := gorm.Open(mysql.Open(mysqlEndpoint), &gorm.Config{})
	if connectionErr != nil {
		panic("Failed to connect to db @ " + mysqlEndpoint + " with error: " + connectionErr.Error())
	}

	db.AutoMigrate(&entity.Animal{}, &entity.Household{})
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

func (db *database) Save(animal entity.Animal) (entity.Animal, error) {
	if err := db.connection.Save(&animal).Error; err != nil {
		return entity.Animal{}, err
	}
	return animal, nil
}

func (db *database) Update(animal entity.Animal) (entity.Animal, error) {
	if err := db.connection.Save(&animal).Error; err != nil {
		return entity.Animal{}, err
	}
	return animal, nil
}

func (db *database) Delete(animal entity.Animal) (entity.Animal, error) {
	if err := db.connection.Delete(&animal.Household).Error; err != nil {
		return entity.Animal{}, err
	}
	if err := db.connection.Delete(&animal).Error; err != nil {
		return entity.Animal{}, err
	}
	return animal, nil
}

func (db *database) FindById(id uint64) (entity.Animal, error) {
	var animal entity.Animal
	if err := db.connection.Preload("Household").First(&animal, id).Error; err != nil {
		return entity.Animal{}, err
	}

	return animal, nil
}

func (db *database) FindAll() []entity.Animal {
	var animals []entity.Animal
	db.connection.Preload("Household").Find(&animals)

	return animals
}

func getDbConfig() databaseConfig {
	var dbConfig databaseConfig
	if loadConfigErr := gonfig.GetConf(getDbConfigFilePath(), &dbConfig); loadConfigErr != nil {
		panic(loadConfigErr.Error())
	}

	return dbConfig
}

func getDbConfigFilePath() string {
	filename := []string{"config/", "db_config", ".json"}
	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), "..", strings.Join(filename, ""))

	return filepath.FromSlash(filePath)
}

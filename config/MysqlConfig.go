package config

import (
	"fmt"
	"gogin-practice/entity"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/tkanos/gonfig"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type databaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type Database struct {
	Connection *gorm.DB
}

func NewMySqlConnection() *Database {
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

	db.Migrator().DropTable(&entity.Animal{})
	db.Migrator().DropTable(&entity.Household{})
	db.AutoMigrate(&entity.Animal{}, &entity.Household{})
	return &Database{
		Connection: db,
	}
}

func (db *Database) Close() {
	sqlDB, err := db.Connection.DB()
	if err != nil {
		panic("Can't access database instance")
	}
	sqlDB.Close()
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

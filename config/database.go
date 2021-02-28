package config

import (
	"fmt"
	"log"

	"github.com/raphaelkieling/go-vote/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDatabase() *gorm.DB {
	newConfig := NewConfig()

	dbHost := newConfig.Database.Host
	dbName := newConfig.Database.Database
	dbPassword := newConfig.Database.Password
	dbUser := newConfig.Database.User

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Banco conectado com sucesso")

	err = db.AutoMigrate(&models.Campaign{}, &models.Vote{}, &models.Audit{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}

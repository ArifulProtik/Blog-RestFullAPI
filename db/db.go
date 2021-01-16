package db

import (
	"fmt"
	"forum-api/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type manager struct {
	db *gorm.DB
}

// Dbprovider Serves all the Db operation.
var Dbprovider Manager

func init() {
	dsn := "host=127.0.0.1 user=postgres password=112233 dbname=blog port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to init db:", err)
	} else {
		log.Println("Database Connected")
	}
	db.Debug().AutoMigrate(model.User{}, model.Profile{}, model.TokenString{}, model.Comment{}, model.Post{}, model.Like{})
	Dbprovider = &manager{db: db}

}
func (c *manager) PingDB() {
	sql, err := c.db.DB()
	if err != nil {
		log.Fatal("Something went wrong", err)
	}
	err = sql.Ping()
	if err != nil {
		log.Fatal("Couldn't Ping DB", err)
	}
	fmt.Println("DB Pinged!!!!!!!")

}

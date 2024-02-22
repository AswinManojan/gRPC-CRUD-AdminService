package db

import (
	"fmt"
	"log"

	"github.com/grpc/gobus/admin/config"
	"github.com/grpc/gobus/admin/pkg/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Db_connect(cfg *config.Config) *gorm.DB {
	host := cfg.HOST
	dbname := cfg.DBNAME
	password := cfg.PASSWORD
	uname := cfg.USERNAME
	port := cfg.PORT
	sslmode := cfg.SSLMODE
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, uname, password, dbname, port, sslmode)
	Db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to DB")
	}
	Db.AutoMigrate(&model.Admin{})
	return Db
}

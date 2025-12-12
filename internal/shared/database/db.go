package database

import (
	"log"

	"github.com/NIROOZbx/project-server/internal/shared/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDb() {

	var err error

	dsn := config.GetConfig().DSN

	db,err=gorm.Open(postgres.Open(dsn),&gorm.Config{})

	if err!=nil{
		log.Fatalf("❌ Failed to connect: %v", err)
	}

	log.Println("✅ Database connected successfuly")

}

func DB() *gorm.DB{
	return db
}
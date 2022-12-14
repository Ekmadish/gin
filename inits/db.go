package inits

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func ConnecToDB() {
	d, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = d

}

func GetDB() *gorm.DB {
	return DB
}

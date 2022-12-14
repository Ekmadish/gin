package main

import (
	"gin-test.com/test/inits"
	"gin-test.com/test/models"
)

func init() {
	inits.LoadEnvVar()
	inits.ConnecToDB()
}
func main() {
	inits.DB.AutoMigrate(&models.Post{})
}

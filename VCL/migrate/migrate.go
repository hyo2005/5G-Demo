package main

import (
	"github.com/hyo2005/demo-5g/initializer"
	"github.com/hyo2005/demo-5g/models"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
}

func main() {
	initializer.DB.AutoMigrate(&models.Subscriber{})

}

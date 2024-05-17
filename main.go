package main

import (
	_ "github.com/robbymaul/golang-swagger.git/config"
	"github.com/robbymaul/golang-swagger.git/router"
)

func main() {
	router := router.InitRoute()

	router.Run("localhost:8005")
}

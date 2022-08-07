package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jsakash/jwt-check/initializers"
	"github.com/jsakash/jwt-check/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()

}

func main() {

	router := gin.Default()
	router.Use(gin.Logger())
	routes.AuthRoutes(router)

	router.Run()

}

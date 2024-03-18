package main

import (
	"fmt"
	"net/http"

	"goland-crud/config"
	"goland-crud/controller"
	"goland-crud/helper"
	"goland-crud/repository"
	"goland-crud/router"
	"goland-crud/service"
)

func main() {
	fmt.Printf("the server is flowing\n")
	//database
	db := config.DatabaseConnection()

	//repository
	GrowBoxRepository := repository.NewGrowBoxRepository(db)

	//service
	GrowBoxService := service.NewGrowBoxServiceImpl(GrowBoxRepository)

	//controller
	GrowBoxController := controller.NewGrowBoxController(GrowBoxService)

	//router
	routes := router.NewRouter(GrowBoxController)

	server := http.Server{Addr: "localhost:8888", Handler: routes}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

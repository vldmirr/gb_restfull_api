package router

import (
	"fmt"
	"goland-crud/controller"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(GrowBoxController *controller.GrowBoxController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "You in safety place")
	})

	router.GET("/api/GrowBox", GrowBoxController.FindAll)
	router.GET("/api/GrowBox/:GrowBoxId", GrowBoxController.FindById)
	router.POST("/api/GrowBox", GrowBoxController.Create)
	router.PATCH("/api/GrowBox/:GrowBoxId", GrowBoxController.Update)
	router.DELETE("/api/GrowBox/:GrowBoxId", GrowBoxController.Delete)

	return router
}

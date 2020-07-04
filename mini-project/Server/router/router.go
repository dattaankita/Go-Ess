package router

import (
	"../middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router{
	router := mux.NewRouter()
	router.HandleFunc("/", middleware.IndexPage)
	router.HandleFunc("/add", middleware.AddTasks)
	router.HandleFunc("/edit", middleware.EditTasks)
	router.HandleFunc("/delete", middleware.DeleteTasks)
	router.HandleFunc("/getall", middleware.GetAllTasks)
	
	return router
}


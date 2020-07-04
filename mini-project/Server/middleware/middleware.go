package middleware

import (
	"context"
	"log"
	"fmt"
	"net/http"
	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	
)
func IndexPage(res http.ResponseWriter, req *http.Request){
	res.Write([]byte("MyResponse"))
}

func AddTasks(res http.ResponseWriter, req *http.Request){
	res.Write([]byte("add"))
}

func DeleteTasks(res http.ResponseWriter, req *http.Request){
	res.Write([]byte("delete"))
}

func EditTasks(res http.ResponseWriter, req *http.Request){
	res.Write([]byte("Edit"))
}
func GetAllTasks(res http.ResponseWriter, req *http.Request){
	res.Write([]byte("get all"))
}

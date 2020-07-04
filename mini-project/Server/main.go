package main

import (
	
	"fmt"
	"log"
	"net/http"
	"./router"
	//"golang.org/x/net/websocket"
)


func main() {
	r := router.Router()
	fmt.Println("Server running on 5000")
	log.Fatal(http.ListenAndServe(":5000",r))
}


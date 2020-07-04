package main

import(
	"github.com/gorilla/mux"
	"encoding/json"
	"fmt"
	"net/http"
	"log"
)

type Book struct(){
}


func getBooks(res http)
func indexPage(res http.ResponseWriter, req *http.Request){
	fmt.Println("I got the request")
	fmt.Fprintf(res, "Welcoje")
	fmt.Println("I am the one")
}
func handleRequests(){
	http.HandleFunc("/", indexPage)
	fmt.Println("Server running in 3000")
	log.Fatal(http.ListenAndServe(":3000",nil))
}

func handleRoutes(){
	myRoutes := mux.NewRouter().StrictSlash()
	myRoutes.HandleFunc("/",indexPage).Methods("GET")
	myRoutes.HandleFunc("/getBooks",indexPage).Methods("POST")
}

func main(){
	Books = []Book{
		Book{Author: "Ankita", 
			 Title: "Go Language",
			 Content: "This is awesome"},
		Book{Author: "Ankita", 
			Title: "Go Language", 
			Content: "This is awesome"}
		
	}
	handleRequests()


}
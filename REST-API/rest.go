package main

import(
	"fmt"
	"net/http"
	"log"
	"encoding/json"
)

type Book struct{
	Title string 'json:Title'
	Author string 'json:Author'
	Content string 'json:Content'
}

func getBooks
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
	http.HandleFunc("/", indexPage)
	http.HandleFunc("/getBooks", getBooks)
}
func main(){
	Books = []Book{
		Book{Author: "Ankita", Title: "Go Language", Content: "This is awesome"},
		Book{Author: "Ankita", Title: "Go Language", Content: "This is awesome"}
		
	}
	handleRequests()


}
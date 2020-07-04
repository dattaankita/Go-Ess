package main

import "fmt"
import "time"

func sendAPIng(c chan string){
	for i := 0; i<10; i++{
		c <- "ping"
	}
}

func printForAPIng(c chan string){
	for{
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	} 
}

func main(){
	mychan := make(chan string)
	go sendAPIng(mychan)
	go printForAPIng(mychan)

	var input string
	fmt.Scanln(&input)
}
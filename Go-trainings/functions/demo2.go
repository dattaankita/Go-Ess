package main

import "fmt"

func main(){
	x,y := cal(10, 20)
	fmt.Println(x)
	fmt.Println(y)
}

// func main(){
// 	_,y := cal(10, 20)
// 	fmt.Println(y)
// }

func cal(a int, b int) (add int, mul int){
	add = a+b
	mul= a*b
	return
}
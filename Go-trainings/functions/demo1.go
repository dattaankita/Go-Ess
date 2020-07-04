package main

import "fmt"

func main(){
	x := cal(10, 20)
	fmt.Println(x)
}
func cal(a int, b int)int{
	return a+b
}
func cal(a int, b int) z int{
	z = a+b
	return
}
package main

import (
	"fmt"
)

func main(){
	var s []int
	//fmt.Println(s==nil)
	var a = [...]int{1,2,3,4,5,6,7,8,9,10}
	s = a[2:6]
	fmt.Println(s)
	newSlice := append(s,1000,200,10000,5000)
	fmt.Println(newSlice)
	fmt.Println(cap(newSlice))
}
	
//slice is an dynamic array data structure
//it reflects the changes done in main array
//array is of fixed size

	// var s []int
	// var a = []int{1,2,3,4,5,6,7,8,9,10}
	// s = a[0:5]
	// fmt.Println(cap(s))       - 10
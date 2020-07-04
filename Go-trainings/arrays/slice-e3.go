package main

import "fmt"


func main(){
	var s []int
	var a = []int{1,2,3,4,5,6,7,8,9,10}

	s = a[0:5]
	fmt.Println(cap(s))

	var c= make([]int,3)
	copy(c,a[2:5])
	fmt.Println(cap(c))
	c = append(c[:1],c[2])
	fmt.Println(c)
}
	

//delete method is not there
//we have to create another slice containing which data we needed
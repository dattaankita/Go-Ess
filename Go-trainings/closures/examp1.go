package main

import "fmt"
func intSequence() func innerfunc() int{
	a := 0
	return func innerfunc() int{
		a++
		return a
	}
}

func main(){
	getnext := intSequence()
	second := intSequence()
	
	fmt.Println(getnext())
	fmt.Println(getnext())
	fmt.Println(getnext())

	fmt.Println(second())
	fmt.Println(second())
	fmt.Println(second())
}

//use in networking for isolation purpose
//inner function can access outer function variables

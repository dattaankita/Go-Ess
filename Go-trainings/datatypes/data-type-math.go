package main

import(
	"fmt";
	"math"

)

func main(){
	var cost float64 =900
	sqrt := math.Sqrt(cost)

	fmt.Printf("Type: %T value: %v\n",cost,cost)
	fmt.Printf("Type: %T value: %v\n",sqrt,sqrt)
}
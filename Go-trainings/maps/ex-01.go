package main

import "fmt"


func main(){
	

	var m= make(map[string]string)
	m["emp0"] = "kiran"
	m["emp1"] = "kiran"
	for key, value :=range m{
		fmt.Println(key,value)
	}
	delete(m, "emp0")
	for key, value :=range m{
		fmt.Println(key,value)
	}
	i,t := m["emp0"]
	fmt.Println(i, t)
}
	

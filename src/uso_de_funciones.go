package main

import "fmt"

func normalFunction(message string){
	fmt.Println(message)
}
func tripeArgunment(a, b int,c string){
	fmt.Println(a,b,c)
}
func returnValue(a int) int{
	return a * 2
}

func doubleReturn(a int)(c,d int){
	return a, a*2
}

func uso_de_funciones(){
	normalFunction("hola mundo")
	tripeArgunment(1,2,"hola")

	value := returnValue(2)
	fmt.Println("value:",value)

	value1,value2 := doubleReturn(2)
	fmt.Println("value1 y value 2:",value1,value2)
}
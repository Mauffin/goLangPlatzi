package main

import "fmt"

func main() {
	//declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi", pi)
	fmt.Println("pi", pi2)

	// declaracion de variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	//Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area Cuadrado:", areaCuadrado)

	x := 10
	y := 50
	//Suma
	result := x + y
	fmt.Println("Suma:", result)

	//Resta
	result = y - x
	fmt.Println("Resta:", result)

	//Multipicacion
	result = y * x
	fmt.Println("Multiplicacion:", result)

	//Dividir
	result = y / x
	fmt.Println("Division", result)

	//Modulo / porcentaje
	result = y % x
	fmt.Println("Modulo / porcentaje", result)

	//Incremental
	x++
	fmt.Println("Incremental", x)

	//Decremental
	x--
	fmt.Println("Decemental", x)

	//Retos
	//-Rectangulo,Trapecio y de un circulo

	var areaRectangulo int
	largo := 5
	ancho := 6
	areaRectangulo = largo * ancho
	fmt.Println("area de un rectangulo", areaRectangulo)

	var areaTrapecio int
	largoT := 5
	anchoT := 6
	areaTrapecio = largoT * anchoT / 2
	fmt.Println("area de un trapecio", areaTrapecio)

	var areaCirculo float64
	var pii float64
  var radio float64

	radio = 5
	pii = 3.14
	areaCirculo =  pii* radio * radio
	fmt.Println("area de un circulo", areaCirculo)

}

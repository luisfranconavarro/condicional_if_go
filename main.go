package main

import "fmt"

func main()  {
	var dia, mes int

	fmt.Print()
	fmt.Scan(&dia)
	fmt.Print()
	fmt.Scan(&mes)

	if (dia >= 21 && dia <= 31 && mes == 1) || (dia >= 1 && dia <= 18 && mes == 2) {
		fmt.Println("Acuario")
	} else if (dia >= 19 && dia <= 29 && mes == 2) || (dia >= 1 && dia <= 20 && mes == 3) {
		fmt.Println("Picis")
	} else if (dia >= 21 && dia <= 31 && mes == 3) || (dia >= 1 && dia <= 20 && mes == 4) {
		fmt.Println("Aries")
	} else if (dia >= 21 && dia <= 30 && mes == 4) || (dia >= 1 && dia <= 20 && mes == 5) {
		fmt.Println("Tauro")
	} else if (dia >= 21 && dia <= 31 && mes == 5) || (dia >= 1 && dia <= 21 && mes == 6) {
		fmt.Println("Geminis")
	} else if (dia >= 22 && dia <= 30 && mes == 6) || (dia >= 1 && dia <= 22 && mes == 7) {
		fmt.Println("Cancer")
	} else if (dia >= 23 && dia <= 31 && mes == 7) || (dia >= 1 && dia <= 22 && mes == 8) {
		fmt.Println("Leo")
	} else if (dia >= 23 && dia <= 31 && mes == 8) || (dia >= 1 && dia <= 22 && mes == 9) {
		fmt.Println("Virgo")
	} else if (dia >= 23 && dia <= 30 && mes == 9) || (dia >= 1 && dia <= 22 && mes == 10) {
		fmt.Println("Libra")
	} else if (dia >= 23 && dia <= 31 && mes == 10) || (dia >= 1 && dia <= 22 && mes == 11) {
		fmt.Println("Escorpion")
	} else if (dia >= 23 && dia <= 30 && mes == 11) || (dia >= 1 && dia <= 21 && mes == 12) {
		fmt.Println("Sagitario")
	} else if (dia >= 22 && dia <= 31 && mes == 12) || (dia >= 1 && dia <= 20 && mes == 1) {
		fmt.Println("Capricornio")
	} else {
		fmt.Println("Dia o mes incorrecto")
	}
}
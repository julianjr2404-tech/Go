package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	var dificultad int
	var numeroSecreto int
	var maxIntentos int

	fmt.Println("===================================")
	fmt.Println("       ADIVINA EL NÚMERO")
	fmt.Println("===================================")
	fmt.Println("1. Fácil (1-100)")
	fmt.Println("2. Difícil (1-1000)")
	fmt.Print("Seleccione dificultad: ")

	fmt.Scan(&dificultad)

	if dificultad == 2 {
		numeroSecreto = rand.Intn(1000) + 1
		maxIntentos = 10
		fmt.Println("")
		fmt.Println("Modo difícil activado")
		fmt.Println("Adivina el número del 1 al 1000")
	} else {
		numeroSecreto = rand.Intn(100) + 1
		maxIntentos = 7
		fmt.Println("")
		fmt.Println("Modo fácil activado")
		fmt.Println("Adivina el número del 1 al 100")
	}

	intentos := 0

	for intentos < maxIntentos {

		var numero int

		fmt.Printf("Intento %d: ", intentos+1)
		fmt.Scan(&numero)

		intentos++

		if numero == numeroSecreto {
			fmt.Println("")
			fmt.Println("¡GANASTE!")
			fmt.Println("Número correcto")
			break
		}

		if numero < numeroSecreto {
			fmt.Println("El número secreto es MAYOR")
		} else {
			fmt.Println("El número secreto es MENOR")
		}

		fmt.Println("")
	}

	if intentos == maxIntentos {
		fmt.Println("")
		fmt.Println("PERDISTE")
		fmt.Println("El número secreto era:", numeroSecreto)
	}
}
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	opciones := []string{"piedra", "papel", "tijera"}

	var jugador string

	fmt.Println("===================================")
	fmt.Println("     PIEDRA PAPEL O TIJERA")
	fmt.Println("===================================")
	fmt.Println("Escribe:")
	fmt.Println("- piedra")
	fmt.Println("- papel")
	fmt.Println("- tijera")
	fmt.Println("")

	fmt.Print("Tu elección: ")
	fmt.Scan(&jugador)

	maquina := opciones[rand.Intn(len(opciones))]

	fmt.Println("")
	fmt.Println("Tú elegiste:", jugador)
	fmt.Println("La máquina eligió:", maquina)
	fmt.Println("")

	if jugador == maquina {
		fmt.Println("EMPATE")
	} else if jugador == "piedra" && maquina == "tijera" {
		fmt.Println("GANASTE")
	} else if jugador == "papel" && maquina == "piedra" {
		fmt.Println("GANASTE")
	} else if jugador == "tijera" && maquina == "papel" {
		fmt.Println("GANASTE")
	} else {
		fmt.Println("PERDISTE")
	}
}
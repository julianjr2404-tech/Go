package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Personaje struct {
	Nombre  string
	Vida    int
	Ataque  int
	Defensa int
}

type Habilidad struct {
	Nombre string
	Bono   int
}

func (p *Personaje) EstaVivo() bool {
	return p.Vida > 0
}

func (p *Personaje) Atacar(objetivo *Personaje) int {

	damage := p.Ataque - objetivo.Defensa

	variacion := rand.Intn(11) - 5

	damage += variacion

	if damage < 1 {
		damage = 1
	}

	objetivo.Vida -= damage

	return damage
}

func elegirHabilidad(habilidades []Habilidad) Habilidad {

	fmt.Println("")
	fmt.Println("Selecciona una habilidad:")

	for i, h := range habilidades {
		fmt.Printf("%d. %s (+%d ataque)\n", i+1, h.Nombre, h.Bono)
	}

	var opcion int
	fmt.Print("Opción: ")
	fmt.Scan(&opcion)

	if opcion < 1 || opcion > len(habilidades) {
		fmt.Println("Opción inválida. Se usará la primera habilidad.")
		return habilidades[0]
	}

	return habilidades[opcion-1]
}

func main() {

	rand.Seed(time.Now().UnixNano())

	heroe := Personaje{
		Nombre:  "Ingeniero",
		Vida:    100,
		Ataque:  20,
		Defensa: 5,
	}

	dragon := Personaje{
		Nombre:  "Dragón de Datos",
		Vida:    120,
		Ataque:  18,
		Defensa: 4,
	}

	habilidadesHeroe := []Habilidad{
		{"Golpe Básico", 0},
		{"Espadazo", 5},
		{"Bola de Fuego", 10},
		{"Ataque Supremo", 15},
	}

	habilidadesDragon := []Habilidad{
		{"Zarpazo", 0},
		{"Mordida", 4},
		{"Llamarada", 8},
		{"Rugido Mortal", 12},
	}

	fmt.Println("===================================")
	fmt.Println("   LA BATALLA POR LA CUMBRE")
	fmt.Println("===================================")

	for heroe.EstaVivo() && dragon.EstaVivo() {

		fmt.Println("")
		fmt.Println("===================================")
		fmt.Printf("%s Vida: %d\n", heroe.Nombre, heroe.Vida)
		fmt.Printf("%s Vida: %d\n", dragon.Nombre, dragon.Vida)
		fmt.Println("===================================")

		habilidadJugador := elegirHabilidad(habilidadesHeroe)

		heroe.Ataque += habilidadJugador.Bono

		damage := heroe.Atacar(&dragon)

		heroe.Ataque -= habilidadJugador.Bono

		fmt.Println("")
		fmt.Printf("%s usa %s\n", heroe.Nombre, habilidadJugador.Nombre)
		fmt.Printf("Hace %d de daño\n", damage)

		if !dragon.EstaVivo() {
			break
		}

		habilidadDragon := habilidadesDragon[rand.Intn(len(habilidadesDragon))]

		dragon.Ataque += habilidadDragon.Bono

		damageDragon := dragon.Atacar(&heroe)

		dragon.Ataque -= habilidadDragon.Bono

		fmt.Println("")
		fmt.Printf("%s usa %s\n", dragon.Nombre, habilidadDragon.Nombre)
		fmt.Printf("Hace %d de daño\n", damageDragon)
	}

	fmt.Println("")
	fmt.Println("===================================")

	if heroe.EstaVivo() {
		fmt.Println("¡GANASTE LA BATALLA!")
	} else {
		fmt.Println("EL DRAGÓN TE DERROTÓ")
	}

	fmt.Println("===================================")
}
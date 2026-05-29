package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var opciones = []string{"piedra", "papel", "tijera"}

func elegirMaquina() string {
	rand.Seed(time.Now().UnixNano())
	return opciones[rand.Intn(len(opciones))]
}

func decidirGanador(jugador string, maquina string) string {

	if jugador == maquina {
		return "EMPATE"
	}

	reglas := map[string]string{
		"piedra": "tijera",
		"papel":  "piedra",
		"tijera": "papel",
	}

	if reglas[jugador] == maquina {
		return "¡GANASTE!"
	}

	return "PERDISTE"
}

func manejarInicio(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func manejarJugar(w http.ResponseWriter, r *http.Request) {

	jugador := r.FormValue("opcion")

	if jugador != "piedra" && jugador != "papel" && jugador != "tijera" {
		fmt.Fprintf(w, "<h1>Opción inválida</h1>")
		return
	}

	maquina := elegirMaquina()

	resultado := decidirGanador(jugador, maquina)

	fmt.Fprintf(w, `
	<html>
	<head>
	<title>Resultado</title>
	<style>
	body{
		background:#111;
		color:white;
		font-family:Arial;
		text-align:center;
		padding-top:100px;
	}
	a{
		color:cyan;
		font-size:20px;
	}
	</style>
	</head>
	<body>

	<h1>Resultado</h1>

	<h2>Tú elegiste: %s</h2>
	<h2>La máquina eligió: %s</h2>

	<h1>%s</h1>

	<br>

	<a href="/">Jugar otra vez</a>

	</body>
	</html>
	`, jugador, maquina, resultado)
}

func main() {

	http.HandleFunc("/", manejarInicio)

	http.HandleFunc("/jugar", manejarJugar)

	fmt.Println("Servidor corriendo en puerto 8080")

	http.ListenAndServe(":8080", nil)
}
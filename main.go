package main

import (
	"log"

	"github.com/davidrr98/api-apoyo/bd"
	"github.com/davidrr98/api-apoyo/handlers"
)

func main() {
	if !bd.ProbarConexion() {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()

}

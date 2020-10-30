package main

import (
	"log"

	"curso/handlers"
	"curso/bd"

)
func main(){
	if !bd.ProbarConexion(){
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
	
}
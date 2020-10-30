package handlers

import (
	"log"
	"net/http"
	"os"

	"curso/middlew"
	"curso/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores es una dependencia*/
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/periodo", middlew.ChequeoBD(routers.RegistroPeriodo)).Methods("POST")
	router.HandleFunc("/periodos", middlew.ChequeoBD(routers.ConsultarPeriodos)).Methods("GET")
	router.HandleFunc("/periodo", middlew.ChequeoBD(routers.ConsultarPeriodo)).Methods("GET")
	router.HandleFunc("/periodo/edit", middlew.ChequeoBD(routers.ActualizarPeriodo)).Methods("PUT")
	router.HandleFunc("/periodo", middlew.ChequeoBD(routers.EliminarPeriodo)).Methods("DELETE")
	router.HandleFunc("/periodo", middlew.ChequeoBD(routers.EliminarPeriodo)).Methods("DELETE")
	router.HandleFunc("/subirInscripciones", middlew.ChequeoBD(routers.SubirArchivo)).Methods("POST")

	/* Swagger documentacion*/
	

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}

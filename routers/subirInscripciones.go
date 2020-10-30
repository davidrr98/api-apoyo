package routers

import (
	"curso/bd"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func SubirArchivo(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
	}
	ID_num, err := strconv.Atoi(ID)

	if err != nil {
		http.Error(w, "El id ingresado no es valido", 400)
		return
	}
	periodo, err := bd.ConsultarPeriodo(ID_num)

	file, handler, err := r.FormFile("inscripciones")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/" + periodo.Nombre + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir el archivo"+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar el archivo"+err.Error(), http.StatusBadRequest)
		return
	}
	periodo.Estado = "inscripciones"
	err = bd.ActualizarPeriodo(periodo)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realzar la actualizacion del periodo "+err.Error(), 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

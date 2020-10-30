package routers

import (
	"strconv"
	
	"encoding/json"
	"net/http"
	
        "github.com/davidrr98/api-apoyo/models"
	"github.com/davidrr98/api-apoyo/bd"
	
)

func RegistroPeriodo(w http.ResponseWriter, r *http.Request) {

	var t models.Periodo
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}
	if len(t.Nombre) == 0 {
		http.Error(w, "El nombre del periodo es requerido", 400)
		return
	}

	/* _, encontrado, _ := bd.YaExistePeriodo(t.nombre)
	if encontrado {
		http.Error(w, "El nobmre de periodo ya existe", 400)
		return
	} */

	_ , err = bd.InsertoPeriodo(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realziar el registro del periodo "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

/* Nos devuelte los periodos */
func ConsultarPeriodos(w http.ResponseWriter, r *http.Request) {

	peridos , err := bd.ConsultarPeriodos()
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realziar el registro del periodo "+err.Error(), 400)
		return
	}
	
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(peridos)
}

func ConsultarPeriodo(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1{
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
	}
	ID_num, err := strconv.Atoi(ID)
	if err != nil {
		http.Error(w, "El id ingresado no es valido", 400)
		return
	}
	periodo, err := bd.ConsultarPeriodo(ID_num)
	
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar buscar el registro del periodo "+err.Error(), 400)
		return
	}
	if periodo.Nombre == "" {
		http.Error(w, "Periodo no encontrado", 400)
		return
	}
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	

	json.NewEncoder(w).Encode(periodo)
}

func ActualizarPeriodo(w http.ResponseWriter, r *http.Request) {

	var t models.Periodo
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	err = bd.ActualizarPeriodo(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realzar la actualizacion del periodo "+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func EliminarPeriodo(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1{
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
	}
	ID_num, err := strconv.Atoi(ID)
	if err != nil {
		http.Error(w, "El id ingresado no es valido", 400)
		return
	}
	err = bd.EliminarPeriodo(ID_num)
	
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar borrar el registro del periodo "+err.Error(), 400)
		return
	}
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}

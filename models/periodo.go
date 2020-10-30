package models

type Periodo struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Estado      string `json:"estado"`
	FechaInicio string `json:"fechaInicio"`
	FechaFin    string `json:"fechaFin"`
}

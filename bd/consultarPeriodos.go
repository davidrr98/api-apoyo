package bd

import (
	"context"
	"github.com/davidrr98/api-apoyo/models"
	"time"
	"strconv"
)

func ConsultarPeriodos() ([]models.Periodo, error) {
	q := `
	SELECT id_periodo, nombre, estado, fecha_inicio, fecha_fin
	FROM periodo;
	`
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	rows, err := DBConect.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var peridos []models.Periodo
	for rows.Next() {
		var p models.Periodo
		rows.Scan(&p.ID, &p.Nombre, &p.Estado, &p.FechaInicio,
			&p.FechaFin)
		p.FechaInicio = p.FechaInicio[:10]
		p.FechaFin = p.FechaFin[:10]
		peridos = append(peridos, p)
	}

	return peridos, nil

}

func ConsultarPeriodo(id int) (models.Periodo, error) {
	q := `
	SELECT id_periodo, nombre, estado, fecha_inicio, fecha_fin
	FROM periodo where periodo.id_periodo=$1 LIMIT 1;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var p models.Periodo

	rows, err := DBConect.QueryContext(ctx, q, id)
	if err != nil {
		return p, err
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&p.ID, &p.Nombre, &p.Estado, &p.FechaInicio,
			&p.FechaFin)
		p.FechaInicio = p.FechaInicio[:10]
		p.FechaFin = p.FechaFin[:10]
	}
	return p, nil
}

func ActulizarPeriodo(u models.Periodo) (string, error) {
	sqlStatement := `
INSERT INTO periodo (nombre, estado, fecha_inicio, fecha_fin)
VALUES ($1, $2, $3, $4)
RETURNING id_periodo`

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := DBConect
	id := 0
	err := db.QueryRowContext(ctx, sqlStatement, u.Nombre, u.Estado, u.FechaInicio, u.FechaFin).Scan(&id)

	return strconv.Itoa(id), err
}


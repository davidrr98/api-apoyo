package bd

import (
	"context"
	"github.com/davidrr98/api-apoyo/models"
	"strconv"
	"time"
)

func InsertoPeriodo(u models.Periodo) (string, error) {
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

func ActualizarPeriodo(u models.Periodo) (error) {
	sqlStatement := `UPDATE periodo
	SET nombre=$1, estado=$2, fecha_inicio=$3, fecha_fin=$4
	WHERE id_periodo=$5;`

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := DBConect
	_, err := db.ExecContext(ctx, sqlStatement, u.Nombre, u.Estado, u.FechaInicio, u.FechaFin, u.ID)

	return err
}
func EliminarPeriodo(id int) (error) {

	sqlStatement := `DELETE FROM periodo
	WHERE id_periodo=$1;`
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := DBConect
	_, err := db.ExecContext(ctx, sqlStatement, id)
	return err
}
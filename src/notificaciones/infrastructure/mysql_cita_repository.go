// mysql_cita_repository.go
package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"notificaciones/src/core"
	repositories "notificaciones/src/notificaciones/domain"
	"notificaciones/src/notificaciones/domain/entities"
)

type MySQLCitaRepository struct {
	conn *sql.DB
}

func NewMySQLCitaRepository() repositories.ICita {
	conn := core.GetDB() // Asegúrate de que esta función esté definida en el paquete core
	return &MySQLCitaRepository{conn: conn}
}

func (mysql *MySQLCitaRepository) Save(cita *entities.Cita) error {
	query := `
		INSERT INTO Cita (nombre_paciente, apellido_paciente, numero_contacto, area_cita, fecha, hora, estado)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	_, err := mysql.conn.Exec(
		query,
		cita.NombrePaciente,
		cita.ApellidoPaciente,
		cita.NumeroContacto,
		cita.AreaCita,
		cita.Fecha,
		cita.Hora,
		"pendiente",
	)
	if err != nil {
		log.Println("Error al guardar la cita:", err)
		return err
	}
	return nil
}

func (mysql *MySQLCitaRepository) Update(cita *entities.Cita) error {
	query := `
        UPDATE Cita
        SET estado = ?
        WHERE cita_id = ?
    `
	log.Printf("Ejecutando query: %s con estado %s y ID %d", query, cita.Estado, cita.CitaID)

	result, err := mysql.conn.Exec(
		query,
		cita.Estado,
		cita.CitaID,
	)
	if err != nil {
		log.Printf("Error al ejecutar la query: %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error al obtener filas afectadas: %v", err)
		return err
	}

	if rowsAffected == 0 {
		log.Printf("No se encontró la cita con ID: %d", cita.CitaID)
		return fmt.Errorf("cita con ID %d no encontrada", cita.CitaID)
	}

	log.Printf("Cita con ID %d actualizada correctamente", cita.CitaID)
	return nil
}

func (mysql *MySQLCitaRepository) Delete(id int32) error {
	query := "DELETE FROM Cita WHERE cita_id = ?"
	result, err := mysql.conn.Exec(query, id)
	if err != nil {
		log.Println("Error al eliminar la cita:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		log.Println("No se encontró la cita con ID:", id)
		return fmt.Errorf("cita con ID %d no encontrada", id)
	}

	return nil
}

func (mysql *MySQLCitaRepository) GetById(id int32) (*entities.Cita, error) {
	query := `
		SELECT cita_id, nombre_paciente, apellido_paciente, numero_contacto, area_cita, fecha, hora, estado
		FROM Cita
		WHERE cita_id = ?
	`
	row := mysql.conn.QueryRow(query, id)

	var cita entities.Cita
	err := row.Scan(
		&cita.CitaID,
		&cita.NombrePaciente,
		&cita.ApellidoPaciente,
		&cita.NumeroContacto,
		&cita.AreaCita,
		&cita.Fecha,
		&cita.Hora,
		&cita.Estado,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Cita no encontrada:", err)
			return nil, fmt.Errorf("cita con ID %d no encontrada", id)
		}
		log.Println("Error al buscar la cita por ID:", err)
		return nil, err
	}

	return &cita, nil
}

func (mysql *MySQLCitaRepository) GetAll() ([]entities.Cita, error) {
	query := `
		SELECT cita_id, nombre_paciente, apellido_paciente, numero_contacto, area_cita, fecha, hora, estado
		FROM Cita
	`
	rows, err := mysql.conn.Query(query)
	if err != nil {
		log.Println("Error al obtener todas las citas:", err)
		return nil, err
	}
	defer rows.Close()

	var citas []entities.Cita
	for rows.Next() {
		var cita entities.Cita
		err := rows.Scan(
			&cita.CitaID,
			&cita.NombrePaciente,
			&cita.ApellidoPaciente,
			&cita.NumeroContacto,
			&cita.AreaCita,
			&cita.Fecha,
			&cita.Hora,
			&cita.Estado,
		)
		if err != nil {
			log.Println("Error al escanear la cita:", err)
			return nil, err
		}
		citas = append(citas, cita)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return citas, nil
}

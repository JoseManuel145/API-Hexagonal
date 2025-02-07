package infrastructure

import (
	"fmt"
	"log"
	"proyecto/src/accessories/domain/entities"
	"proyecto/src/core"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

func NewMySQL() *MySQL {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

func (mysql *MySQL) Save(name, description string) error {
	query := "INSERT INTO accesories (name, description) VALUES (?, ?)"

	result, err := mysql.conn.ExecutePreparedQuery(query, name, description)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
	}
	return nil
}

func (mysql *MySQL) Delete(id int) error {
	query := "DELETE FROM accesories WHERE id = ?"

	result, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
        return fmt.Errorf("error al ejecutar la consulta: %w", err)
    }
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
        return fmt.Errorf("no se encontró ninguna mascota con el ID %d", id)
    }
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
	}
	return nil
}

func (mysql *MySQL) ViewAll() ([]entities.Accessory, error) {
	query := "SELECT * FROM accesories"
	rows := mysql.conn.FetchRows(query)
	defer rows.Close()

	var accessories []entities.Accessory
	for rows.Next() {
		var accessory entities.Accessory
		if err := rows.Scan(&accessory.Id, &accessory.Name, &accessory.Description); err != nil {
			return nil, fmt.Errorf("error al escanear la fila: %w", err)
		}
		accessories = append(accessories, accessory)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando sobre las filas: %w", err)
	}
	return accessories, nil
}

func (mysql *MySQL) ViewOne(id int) (*entities.Accessory, error) {
	query := "SELECT * FROM accesories WHERE id = ?"
	rows := mysql.conn.FetchRows(query, id)
	defer rows.Close()

	var accessory entities.Accessory
	if rows.Next() {
		if err := rows.Scan(&accessory.Id, &accessory.Name, &accessory.Description); err != nil {
			return nil, fmt.Errorf("error al escanear la fila: %w", err)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando sobre las filas: %w", err)
	}
	return &accessory, nil
}

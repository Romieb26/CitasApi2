// mySQL.go
package core

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

// InitDB inicializa la conexión a la base de datos.
func InitDB() {
	if err := godotenv.Load(); err != nil {
		log.Println("Advertencia: No se pudo cargar .env")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error al conectar a la BD:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("No se pudo conectar a la BD:", err)
	}

	fmt.Println("Conexión a la BD exitosa")
}

// GetDB retorna la conexión a la base de datos.
func GetDB() *sql.DB {
	return db
}

package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database = func() (db *gorm.DB) {
	errorVariables := godotenv.Load()
	if errorVariables != nil {
		panic(errorVariables)

	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s search_path=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
		os.Getenv("DB_SCHEMA"),
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		fmt.Println("error en conexion")
		panic(err)
	}

	schema := os.Getenv("DB_SCHEMA")

	// 1. Crear el schema si no existe
	db.Exec(fmt.Sprintf(`CREATE SCHEMA IF NOT EXISTS "%s"`, schema))

	// 2. Aplicar search_path a nivel de usuario
	db.Exec(fmt.Sprintf(`ALTER ROLE "%s" SET search_path TO "%s"`, os.Getenv("DB_USER"), schema))

	// 3. Aplicar search_path en la sesión actual
	db.Exec(fmt.Sprintf(`SET search_path TO "%s"`, schema))

	fmt.Println("conexion exitosa")
	return db

}

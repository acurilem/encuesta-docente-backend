package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var Database = func() (db *gorm.DB) {
	// if os.Getenv("GO_REST_ENV") == "dev" {
	// Para que funcione en windows.....
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error cargando el archivo .env EStOY EN DAtABASE Go")
		return
	}
	// }

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbInstance := os.Getenv("DB_INSTANCE")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	dbDatabase := os.Getenv("DB_DB")

	dsn := fmt.Sprintf(
		"sqlserver://%s:%s@%s:%s?instance=%s&database=%s&encrypt=disable&connection+timeout=30",
		dbUser, dbPassword, dbHost, dbPort, dbInstance, dbDatabase,
	)
	if db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println("Error en la Conexion", err)
		panic(err)
	} else {
		fmt.Println("Conexion sqlserver exitosa")
		db.Debug()
		return db
	}
}()

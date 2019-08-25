package connection

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func DBConnect() *sql.DB {
	dbHost := `127.0.0.1`
	dbPort := `8889`
	dbUser := `root`
	dbPass := `root`
	dbName := `go_test`
	connections := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	dsn := fmt.Sprintf("%s?%s", connections)
	db, err := sql.Open(`mysql`, dsn)

	if err != nil {
		panic(err)
	}

	return db
}
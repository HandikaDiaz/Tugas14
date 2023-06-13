package connection

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

var Conn *pgx.Conn

func DatabaseConnection() {
	databaseurl := "postgres://postgres:123@localhost:5432/first"

	var err error
	Conn, err = pgx.Connect(context.Background(), databaseurl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Successfuly connected to database")
}

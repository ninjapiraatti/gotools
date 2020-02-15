package gotools

import (
	"database/sql"

	// Using the driver as described in the documentation, don't know why it's blank
	_ "github.com/go-sql-driver/mysql"
)

// AddCreatureDB writes a creature to database.
func AddCreatureDB(creature string) {
	db, err := sql.Open("mysql", "gotesti_1:Xv7NNK181fdAiBe2@tcp(dedi1450.your-server.de:3306)/gotesti")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO animals(name) VALUES(?)", creature)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}

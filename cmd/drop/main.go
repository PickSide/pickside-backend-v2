package main

import (
	"log"
	"me/pickside/db"
	"me/pickside/db/queries"
)

func main() {
	log.Println("Dropping tables...")

	if err := db.OpenConnectionToDB(); err != nil {
		log.Fatal(err)
	}

	defer db.CloseDB()

	qs := []string{
		queries.DropUserTables,
		queries.DropUserSettingsTable,
	}

	for _, q := range qs {
		_, err := db.GetDB().Exec(q)
		if err != nil {
			panic(err)
		}
	}
}

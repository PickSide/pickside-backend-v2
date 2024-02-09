package main

import (
	"log"
	"pickside/service/db"
	"pickside/service/db/queries"
)

func main() {
	log.Println("Dropping tables...")

	if err := db.OpenConnectionToDB(); err != nil {
		log.Fatal(err)
	}

	defer db.CloseDB()

	qs := []string{
		queries.DropActivityTable,
		queries.DropActivityUserTable,
		queries.DropChatroomParticipantsTable,
		queries.DropChatroomTable,
		queries.DropGameModesTable,
		queries.DropGroupMembersTable,
		queries.DropGroupTable,
		queries.DropLocaleTable,
		queries.DropMessageTable,
		queries.DropNotificationTable,
		queries.DropSportGameModesTable,
		queries.DropSportTable,
		queries.DropTokensTable,
		queries.DropUserSettingsTable,
		queries.DropUserTables,
	}

	for _, q := range qs {
		_, err := db.GetDB().Exec(q)
		if err != nil {
			panic(err)
		}
	}
}

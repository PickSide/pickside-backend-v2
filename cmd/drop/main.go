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
		queries.DropActivities,
		queries.DropChatrooms,
		queries.DropGameModes,
		queries.DropGroups,
		queries.DropLocales,
		queries.DropMessage,
		queries.DropNotifications,
		queries.DropSports,
		queries.DropTokens,
		queries.DropUsers,
		queries.DropActivityUsers,
		queries.DropChatroomUsers,
		queries.DropSportGameModes,
		queries.DropGroupUsers,
	}

	for _, q := range qs {
		_, err := db.GetDB().Exec(q)
		if err != nil {
			panic(err)
		}
	}
}

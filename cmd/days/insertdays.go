package main

import (
	"context"
	"log"
	"os"

	"github.com/rmarken5/lottery-calendar/app"
	"github.com/rmarken5/lottery-calendar/database/postgres"
	"github.com/rmarken5/lottery-calendar/lottery"
)

func main() {

	db, err := app.ConnectDatabase()
	if err != nil {
		log.Fatal("unable to connect to database ", err)
	}

	repo := postgres.New(db, db)

	controller := lottery.New(repo)

	file, err := os.ReadFile("jc-calendar.json")
	if err != nil {
		log.Fatal(err)
	}

	err = controller.BulkInsertCalendarDays(context.Background(), file)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("All days inserted for calendar")
}

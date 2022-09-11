package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	"github.com/dixonwille/wmenu"
	_ "github.com/mattn/go-sqlite3"
	"goTravel2.0/controllers"
	"goTravel2.0/database"
	"goTravel2.0/services"
)

var migrate *string

func init() {
	migrate = flag.String("migrate", "", "Put yes for migrate")
}
func main() {
	flag.Parse()

	db, err := sql.Open("sqlite3", "./travel2.0.db")
	if err != nil {
		log.Fatal(err)
	}

	db.Exec(`PRAGMA foreign_keys = ON;`)

	database := database.NewDatabase(db)

	if *migrate == "yes" {
		fmt.Println("execute migrations")
		database.Seed()
	}
	// defer close
	defer database.Close()

	service := services.NewService(database)
	controller := controllers.NewController(service)

	menu := wmenu.NewMenu("Welcome ,What would you like to do?")

	menu.Action(func(opts []wmenu.Opt) error { controller.HandleFunc(opts); return nil })

	menu.Option("Add a clothes plans for your travel", 0, true, nil)
	menu.Option("Add a new travel", 1, false, nil)
	menu.Option("Find you travel", 2, false, nil)
	menu.Option("Update a travel's information", 3, false, nil)
	menu.Option("Delete a travel by ID", 4, false, nil)
	menu.Option("Delete a clothes by ID", 5, false, nil)
	menu.Option("Quit Application", 6, false, nil)
	menuerr := menu.Run()

	if menuerr != nil {
		log.Fatal(menuerr)
	}
}

//tasks : function Add new clothes and update architecture

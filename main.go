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

	if *migrate == "yes" {
		fmt.Println("execute migrations")
		database.Seed(db)
	}
	// defer close
	defer db.Close()

	menu := wmenu.NewMenu("Welcome ,What would you like to do?")

	menu.Action(func(opts []wmenu.Opt) error { controllers.HandleFunc(db, opts); return nil })

	menu.Option("Add a new travel", 0, true, nil)
	menu.Option("Find you travel", 1, false, nil)
	menu.Option("Update a travel's information", 2, false, nil)
	menu.Option("Delete a travel by ID", 3, false, nil)
	menu.Option("Quit Application", 4, false, nil)
	menu.Option("Add a new clothes", 5, false, nil)
	menuerr := menu.Run()

	if menuerr != nil {
		log.Fatal(menuerr)
	}
}

//tasks : function Add new clothes and update architecture

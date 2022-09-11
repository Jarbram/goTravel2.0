package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	cmd "goTravel2.0/cmd/cli"
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
	// this is only for sqlite3
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
	cli := cmd.NewCLI(controller)

	cli.Init()
}

//tasks : function Add new clothes and update architecture

package database

import (
	"database/sql"
	"log"
)

type Database struct {
	Client *sql.DB
}

func NewDatabase(c *sql.DB) (*Database, error) {
	return &Database{c}, nil
}

func (db *Database) Close() {
	db.Client.Close()
}

func Seed(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS  "travels"  (
		"id"	INTEGER,
		"destination"	TEXT,
		"date"	TEXT,
		"budget"	TEXT,
		"clothes_id" INTEGER,
		PRIMARY KEY("id" AUTOINCREMENT),
		FOREIGN KEY("clothes_id") REFERENCES clothes(id)
	);
	CREATE TABLE IF NOT EXISTS "clothes"(
		"id" INTEGER,
		"pants" INTEGER,
		"shirts" INTEGER,
		"travels_id" INTEGER,
		PRIMARY KEY("id" AUTOINCREMENT),
		FOREIGN KEY("travels_id") REFERENCES travels(id)
	);
	`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("seed fails: %v", err)
	}
}

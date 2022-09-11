package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"goTravel2.0/models"
)

type Database struct {
	Client *sql.DB
}

func NewDatabase(db *sql.DB) *Database {
	return &Database{Client: db}
}

func (db *Database) Close() {
	db.Client.Close()
}

func (d *Database) Seed() {
	query := `
	CREATE TABLE IF NOT EXISTS  "travels"(
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
		"underwear" INTEGER,
		"pants" INTEGER,
		"shirts" INTEGER,
		"tshirts" INTEGER,
		"shoes" INTEGER,
		"travels_id" INTEGER,
		PRIMARY KEY("id" AUTOINCREMENT),
		FOREIGN KEY("travels_id") REFERENCES travels(id)
	);
	`
	_, err := d.Client.Exec(query)
	if err != nil {
		log.Fatalf("seed fails: %v", err)
	}
}

func (d *Database) AddClothes(newClothes *models.Clothes) {
	//We create a new SQL statement, stmt. We use db.Prepare to prepare our insert statement and protect the application from SQL injection.

	stmt, _ := d.Client.Prepare("INSERT INTO clothes (id, underwear, pants, shirts, tshirts, shoes) VALUES (?, ?, ?, ?, ?, ?)")
	//Then we run stmt.Exec with the parameters we want to insert.
	result, err := stmt.Exec(nil, newClothes.Underwear, newClothes.Pants, newClothes.Shirts, newClothes.TShirts, newClothes.Shoes)
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	newLastId := strconv.Itoa(int(lastId))
	clothes := d.GetClothesById(newLastId)
	//Then defer the close method and print our results.
	defer stmt.Close()

	fmt.Printf("Added  New Clothes to %+v\n", clothes)

}

func (d *Database) GetClothesById(ourID string) models.Clothes {

	rows, _ := d.Client.Query("SELECT id,underwear, pants, shirts, tshirts, shoes  FROM clothes WHERE id = '" + ourID + "'")
	defer rows.Close()

	ourClothes := models.Clothes{}
	//We then create a new travel object and iterate through the row, scanning each value to the object. Once completed, we return it.

	for rows.Next() {
		rows.Scan(&ourClothes.ID, &ourClothes.Underwear, &ourClothes.Pants, &ourClothes.Shirts, &ourClothes.TShirts, &ourClothes.Shoes)
	}

	return ourClothes

}

func (d *Database) AddTravel(travel *models.Travel) {
	//We create a new SQL statement, stmt. We use db.Prepare to prepare our insert statement and protect the application from SQL injection.

	stmt, _ := d.Client.Prepare("INSERT INTO travels (id, destination, date, budget, clothes_id) VALUES (?, ?, ?, ?, ?)")
	//Then we run stmt.Exec with the parameters we want to insert.
	stmt.Exec(nil, travel.Destination, travel.Date, travel.Budget, travel.Clothes.ID)
	//Then defer the close method and print our results.
	defer stmt.Close()

	fmt.Printf("Added  New travel to  %v \n", travel.Destination)
}

func (d *Database) SearchForTravel(searchString string) []models.Travel {

	//this function that takes the db object and a search string and returns a slice of travel objects (structures).

	rows, err := d.Client.Query("SELECT id, destination, date, budget  FROM travels WHERE destination like '%" + searchString + "%'")
	//We will execute a SELECT statement to select destination,date,budget based on whether the destination matches our search string.
	defer rows.Close()

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	travel := make([]models.Travel, 0)
	//we create a travel structure and fill it with the resulting data. Then we add it to our segment and return
	for rows.Next() {
		ourTravel := models.Travel{}
		err = rows.Scan(&ourTravel.ID, &ourTravel.Destination, &ourTravel.AuxDate, &ourTravel.Budget)
		if err != nil {
			log.Fatal(err)
		}

		newDate, err := time.Parse("2006-01-02 00:00:00+00:00", ourTravel.AuxDate)
		if err != nil {
			log.Fatalf("We can't convert date: %v", err)
		}

		ourTravel.Date = newDate

		travel = append(travel, ourTravel)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return travel
}

func (d *Database) GetTravelById(ourID string) models.Travel {

	fmt.Printf("Value is %v", ourID)

	rows, _ := d.Client.Query("SELECT id, destination, date, budget  FROM travels WHERE id = '" + ourID + "'")
	defer rows.Close()

	ourTravel := models.Travel{}
	//We then create a new travel object and iterate through the row, scanning each value to the object. Once completed, we return it.

	for rows.Next() {
		rows.Scan(&ourTravel.ID, &ourTravel.Destination, &ourTravel.AuxDate, &ourTravel.Budget)
	}

	newDate, err := time.Parse("2006-01-02 00:00:00+00:00", ourTravel.AuxDate)
	if err != nil {
		log.Fatalf("We can't convert date: %v", err)
	}

	ourTravel.Date = newDate
	ourTravel.AuxDate = ""

	fmt.Println("this user you will edit: ", ourTravel)
	return ourTravel

}

func (d *Database) UpdateTravel(ourTravel models.Travel) int64 {

	stmt, err := d.Client.Prepare("UPDATE travels set destination = ?, date = ?, budget = ? where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(ourTravel.Destination, ourTravel.Date, ourTravel.Budget, ourTravel.ID)
	if err != nil {
		log.Fatal(err)
	}

	result, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	//we use a prepared statement and use the values of the passed travel object to execute an UPDATE on the database. We execute the statement and return the affected rows, which should be one.
	return result
}

func (d *Database) DeleteTravel(idToDelete string) int64 {
	//It takes our db connection and the ID to delete.
	stmt, err := d.Client.Prepare("DELETE FROM travels where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	//It prepares a statement that is DELETE and accepts a parameter for id. That is inserted into stmt.Exec and executed.
	res, err := stmt.Exec(idToDelete)
	if err != nil {
		log.Fatal(err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	return affected

}

func (d *Database) DeleteClothes(idToDelete string) int64 {
	//It takes our db connection and the ID to delete.
	stmt, err := d.Client.Prepare("DELETE FROM Clothes where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	//It prepares a statement that is DELETE and accepts a parameter for id. That is inserted into stmt.Exec and executed.
	res, err := stmt.Exec(idToDelete)
	if err != nil {
		if err.Error() == "FOREIGN KEY constraint failed" {
			log.Fatal(errors.New("clothes can't delete because it relate with a travel"))
		}
		log.Fatal(err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	return affected

}

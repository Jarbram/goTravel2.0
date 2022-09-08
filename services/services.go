package services

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type Travel struct {
	ID          int64
	Destination string
	Date        time.Time
	AuxDate     string
	Budget      float64
	Clothes     Clothes
}

type Clothes struct {
	ID     int64
	Pants  uint8
	Shirts uint8
}

func NewTravel(destination string, date time.Time, budget float64) *Travel {
	return &Travel{
		Destination: destination,
		Date:        date,
		Budget:      budget,
	}
}

func AddTravel(db *sql.DB, newTravel *Travel) {
	//We create a new SQL statement, stmt. We use db.Prepare to prepare our insert statement and protect the application from SQL injection.

	stmt, _ := db.Prepare("INSERT INTO travels (id, destination, date, budget) VALUES (?, ?, ?, ?)")
	//Then we run stmt.Exec with the parameters we want to insert.
	stmt.Exec(nil, newTravel.Destination, newTravel.Date, newTravel.Budget)
	//Then defer the close method and print our results.
	defer stmt.Close()

	fmt.Printf("Added  New travel to  %v \n", newTravel.Destination)
}

func SearchForTravel(db *sql.DB, searchString string) []Travel {

	//this function that takes the db object and a search string and returns a slice of travel objects (structures).

	rows, err := db.Query("SELECT id, destination, date, budget  FROM travels WHERE destination like '%" + searchString + "%'")
	//We will execute a SELECT statement to select destination,date,budget based on whether the destination matches our search string.
	defer rows.Close()

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	travel := make([]Travel, 0)
	//we create a travel structure and fill it with the resulting data. Then we add it to our segment and return
	for rows.Next() {
		ourTravel := Travel{}
		err = rows.Scan(&ourTravel.ID, &ourTravel.Destination, &ourTravel.AuxDate, &ourTravel.Budget)
		if err != nil {
			log.Fatal(err)
		}
		//formattedDate := strings.Split(
		//	ourTravel.AuxDate, " ",
		//)[0]

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

func GetTravelById(db *sql.DB, ourID string) Travel {

	fmt.Printf("Value is %v", ourID)

	rows, _ := db.Query("SELECT id, destination, date, budget  FROM travels WHERE id = '" + ourID + "'")
	defer rows.Close()

	OurTravel := Travel{}
	//We then create a new travel object and iterate through the row, scanning each value to the object. Once completed, we return it.

	for rows.Next() {
		rows.Scan(&OurTravel.ID, &OurTravel.Destination, &OurTravel.Date, &OurTravel.Budget)
	}
	fmt.Println("this user you will edit: ", OurTravel)
	return OurTravel

}

//we need to process the new travel object and update the database. We'll do it with the updateTravel function:

func UpdateTravel(db *sql.DB, OurTravel Travel) int64 {

	stmt, err := db.Prepare("UPDATE travels set destination = ?, date = ?, budget = ? where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(OurTravel.Destination, OurTravel.Date, OurTravel.Budget, OurTravel.ID)
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

func DeleteTravel(db *sql.DB, idToDelete string) int64 {
	//It takes our db connection and the ID to delete.
	stmt, err := db.Prepare("DELETE FROM travels where id = ?")
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

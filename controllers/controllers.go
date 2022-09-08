package controllers

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dixonwille/wmenu"
	"goTravel2.0/services"
)

func HandleFunc(db *sql.DB, opts []wmenu.Opt) {

	switch opts[0].Value {

	case 0:
		//we use a bufio scanner to read a new travels's
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter your destination: ")
		destination, _ := reader.ReadString('\n')
		if destination != "\n" {
			destination = strings.TrimSuffix(destination, "\n")
		}

		fmt.Print("Enter enter the date of your trip: ")
		date, _ := reader.ReadString('\n')
		if date != "\n" {
			date = strings.TrimSuffix(date, "\n")
		}
		newDate, err := time.Parse("2006-01-02", date)
		if err != nil {
			log.Fatalf("We can't convert date: %v", err)
		}

		fmt.Print("Enter your budget: ")
		budget, _ := reader.ReadString('\n')
		if budget != "\n" {
			budget = strings.TrimSuffix(budget, "\n")
		}
		newBudget, err := strconv.ParseFloat(budget, 64)

		newTravel := services.NewTravel(destination, newDate, newBudget)
		//We read those values into a buffer one by one, then pass it to the AddTravel method

		services.AddTravel(db, newTravel)
	case 1:

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter a travel to search for : ")
		searchString, _ := reader.ReadString('\n')
		searchString = strings.TrimSuffix(searchString, "\n")
		//we will create a variable named people to store our results
		travel := services.SearchForTravel(db, searchString)

		//we created another bufio reader. We read the travel you are looking for in the searchString.It is completed with the searchForTravel function

		fmt.Printf("Found %v results\n", len(travel))
		//This function returns a list of people results based on our search string.
		for _, ourTravel := range travel {
			fmt.Printf("\n----\nID: %d\nDestination: %s\nDate: %s\nBudget: %f\n", ourTravel.ID, ourTravel.Destination, ourTravel.Date, ourTravel.Budget)
		}
		//We will then loop through the results and display them on the screen.
	case 2:

		//We create another bufio scanner to read in the ID you want to update.
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter an id to update: ")
		updateid, _ := reader.ReadString('\n')
		//We then look up that ID with getTravelById and store it in currenttarvel.
		currentTravel := services.GetTravelById(db, updateid)
		//We then check each value and display the current value while requesting a new value.
		fmt.Printf("Destination (Currently %s):", currentTravel.Destination)
		destination, _ := reader.ReadString('\n')
		if destination != "\n" {
			currentTravel.Destination = strings.TrimSuffix(destination, "\n")
		}

		//fmt.Printf("Date (Currently %s):", currentTravel.Date)
		//date, _ := reader.ReadString('\n')
		//if date != "\n" {
		//	currentTravel.Date = strings.TrimSuffix(date, "\n")
		//}

		//fmt.Printf("Budget (Currently %s):", currentTravel.Budget)
		//budget, _ := reader.ReadString('\n')
		//if budget != "\n" {
		//	currentTravel.Budget = strings.TrimSuffix(budget, "\n")
		//}

		//If the user presses enter, it will keep the current value. If they write something new, it will be updated in the currentTravel object.

		//We'll create a variable named affected, and call the updatePerson method, and pass the db connect method and the currentTravel object with the new information.
		affected := services.UpdateTravel(db, currentTravel)

		if affected == 1 {
			fmt.Println("One row affected")
		}
		//If the update is successful, we will return a message to you.
	case 3:

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter the ID you want to delete : ")
		searchString, _ := reader.ReadString('\n')

		idToDelete := strings.TrimSuffix(searchString, "\n")

		affected := services.DeleteTravel(db, idToDelete)
		//This method takes our db object and the id we want to remove and returns affected, which should be 1.
		if affected == 1 {
			fmt.Println("Deleted Travel from database")
		}
	case 4:
		fmt.Println("Goodbye!")
		os.Exit(3)
	}
}
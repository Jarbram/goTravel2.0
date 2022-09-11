package controllers

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"goTravel2.0/models"
	"goTravel2.0/services"
)

type Controller struct {
	service *services.Service
}

func NewController(service *services.Service) *Controller {
	return &Controller{service}
}

func (c *Controller) AddClothes() {
	var pants uint8
	var shirts uint8

	fmt.Print("Enter how pants you need for to travel? \n")
	fmt.Scanln(&pants, '\n')

	fmt.Print("Enter how shirts you need for to travel? \n")
	fmt.Scanln(&shirts, '\n')

	newClothes := &models.Clothes{
		Pants:  pants,
		Shirts: shirts,
	}

	c.service.AddClothes(newClothes)
}

func (c *Controller) AddTravel() {
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
	newBudget, _ := strconv.ParseFloat(budget, 64)

	fmt.Print("Enter your Clothes ID:")
	clothes_id, _ := reader.ReadString('\n')
	clothes_id = strings.TrimSuffix(clothes_id, "\n")

	newClothes, err := c.service.GetClothesById(clothes_id)

	newTravel := &models.Travel{
		Destination: destination,
		Date:        newDate,
		Budget:      newBudget,
		Clothes:     *newClothes,
	}
	//We read those values into a buffer one by one, then pass it to the AddTravel method

	c.service.AddTravel(newTravel)
}

func (c *Controller) SearchForTravel() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a travel to search for : ")
	searchString, _ := reader.ReadString('\n')
	searchString = strings.TrimSuffix(searchString, "\n")
	//we will create a variable named people to store our results
	travel, _ := c.service.SearchForTravel(searchString)

	//we created another bufio reader. We read the travel you are looking for in the searchString.It is completed with the searchForTravel function

	fmt.Printf("Found %v results\n", len(travel))
	//This function returns a list of people results based on our search string.
	for _, ourTravel := range travel {
		fmt.Printf("\n----\nID: %d\nDestination: %s\nDate: %s\nBudget: %f\n", ourTravel.ID, ourTravel.Destination, ourTravel.Date, ourTravel.Budget)
	}
}

func (c *Controller) UpdateTravel() {
	//We create another bufio scanner to read in the ID you want to update.
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter an id to update: ")
	updateid, _ := reader.ReadString('\n')
	//We then look up that ID with getTravelById and store it in currenttarvel.
	currentTravel, _ := c.service.GetTravelById(updateid)
	//We then check each value and display the current value while requesting a new value.
	fmt.Printf("Destination (Currently %s):", currentTravel.Destination)
	destination, _ := reader.ReadString('\n')
	if destination != "\n" {
		currentTravel.Destination = strings.TrimSuffix(destination, "\n")
	}

	fmt.Printf("Date (Currently %s):", currentTravel.Date)
	date, _ := reader.ReadString('\n')
	if date != "\n" {
		cleanDate := strings.TrimSuffix(date, "\n")
		newDate, err := time.Parse("2006-01-02", cleanDate)
		if err != nil {
			log.Fatalf("We can't convert date: %v", err)
		}
		currentTravel.Date = newDate
	}

	fmt.Printf("Budget (Currently %f):", currentTravel.Budget)
	budget, _ := reader.ReadString('\n')
	if budget != "\n" {
		cleanBudget := strings.TrimSuffix(budget, "\n")
		newBudget, err := strconv.ParseFloat(cleanBudget, 64)
		if err != nil {
			log.Fatalf("We can`t convert budget: %v", err)
		}

		currentTravel.Budget = newBudget
	}

	//If the user presses enter, it will keep the current value. If they write something new, it will be updated in the currentTravel object.

	//We'll create a variable named affected, and call the updatePerson method, and pass the db connect method and the currentTravel object with the new information.
	affected := c.service.UpdateTravel(*currentTravel)

	if affected == 1 {
		fmt.Println("One row affected")
	}
}

func (c *Controller) DeleteTravel() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the ID you want to delete : ")
	searchString, _ := reader.ReadString('\n')

	idToDelete := strings.TrimSuffix(searchString, "\n")

	affected := c.service.DeleteTravel(idToDelete)
	//This method takes our db object and the id we want to remove and returns affected, which should be 1.
	if affected == 1 {
		fmt.Println("Deleted Travel from database")
	}
}

func (c *Controller) DeleteClothes() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the ID you want to delete : ")
	searchString, _ := reader.ReadString('\n')

	idToDelete := strings.TrimSuffix(searchString, "\n")

	affected := c.service.DeleteClothes(idToDelete)
	//This method takes our db object and the id we want to remove and returns affected, which should be 1.
	if affected == 1 {
		fmt.Println("Deleted Clothes from database")
	}
}

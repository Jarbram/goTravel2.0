package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/dixonwille/wmenu"
	"goTravel2.0/controllers"
)

type CLI struct {
	controllers *controllers.Controller
}

func NewCLI(c *controllers.Controller) *CLI {
	return &CLI{c}
}

func (c *CLI) HandleFunc(opts []wmenu.Opt) {

	switch opts[0].Value {

	case 0:
		c.controllers.AddClothes()

	case 1:

		c.controllers.AddTravel()

	case 2:

		c.controllers.SearchForTravel()

	case 3:

		c.controllers.UpdateTravel()

	case 4:

		c.controllers.DeleteTravel()

	case 5:
		c.controllers.DeleteClothes()

	case 6:
		fmt.Println("Goodbye!")
		os.Exit(3)
	}

}
func (c *CLI) Init() {
	menu := wmenu.NewMenu("Welcome ,What would you like to do?")

	menu.Action(func(opts []wmenu.Opt) error { c.HandleFunc(opts); return nil })

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

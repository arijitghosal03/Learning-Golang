package main

import (
	"fmt"
	"github.com/gofiber/fiber"
)

func setupRoutes(app *fiber.App) {
	app.Get("/getleads", getLeads)
	app.Get("/lead/:id", getLeadByID)
	app.Post("/postleads", newLead)
	app.Delete("/deletelead/:id", deleteLead)
}
func initDatabase() {
    
	fmt.Println("Database initialized")
}
func main(){
	app := fiber.New()
	setupRoutes(app)
	fmt.Println("Server is running on port 8080")
	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}
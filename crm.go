package main

import (
	"fmt"
	"github.com/arijitghosal03/Learning-Golang/database"
	"github.com/arijitghosal03/Learning-Golang/lead"
	"github.com/gofiber/fiber"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/jinzhu/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("api/v1/leads", lead.getLeads)
	app.Get("api/v1/leads/:id", lead.getLeadByID)
	app.Post("api/v1/leads", lead.newLead)
	app.Delete("api/v1/leads/:id", lead.deleteLead)
}
func initDatabase() {
    
	fmt.Println("Database initialized")
	varr err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Database connected")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database migrated")
}
func main(){
	app := fiber.New()
	setupRoutes(app)
	initDatabase()
	fmt.Println("Server is running on port 8080")
	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
	defer database.DBConn.Close()
	fmt.Println("Server stopped")
}
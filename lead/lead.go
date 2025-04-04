package lead
import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/arijitghosal03/Learning-Golang/database"
	
	"github.com/gofiber/fiber"
)
type Lead struct {
	gorm.Model
	Name string  `json:"name" `
	Company string  `json:"company" `
	Email string    `json:"email"`
	Phone string     `json:"phone"`
}

func getLeads(c *fiber.Ctx) {
	db := database.DBConn
	var leads []Lead
	if err := db.Find(&leads).Error; err != nil {
		c.Status(500).SendString("Error retrieving leads")
		return
	}
	c.JSON(leads)
}

func newLead(c *fiber.Ctx) {
	db := database.DBConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&lead)
	c.JSON(lead)
}
func getLeadByID(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	if err := db.Where("id = ?", id).Find(&lead).Error; err != nil {
		c.Status(404).SendString("Lead not found")
		return
	}
	c.JSON(lead)
}		
func deleteLead(c *fiber.Ctx) {	
   id:= c.Params("id")
   db := database.DBConn
   var lead Lead
   db.First(&lead, id)
   if lead.Name == "" {
	   c.Status(500).SendString("Lead not found")
	   return
   }
   db.Delete(&lead)
   c.SendString("Lead successfully deleted")
}
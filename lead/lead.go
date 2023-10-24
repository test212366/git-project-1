package lead

import (
	"github.com/akhil/go-fiber-crm-basic/database"

	_ "github.com/dialects/sqlite"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLeads(c *fiber.Ctx) {
	db := database.DBConn
	var leads []Lead
	db.FInd(&leads)
	c.JSON(leads)

}
func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var lead lead
	db.Find(&lead, id)
	c.JSON(lead)
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
func DeleteLeads(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		c.Status(500).Send("No lead Found")
		return
	}
	db.Delete(&lead)
	c.Send("lead deleted")
}

package main

import (
	"fmt"

	"github.com/akhil/go-fiber-crm-basic/database"
	"github.com/akhil/go-fiber-crm-basic/lead"
	_ "github.com/dialects/sqlite"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func setupRoutes(app *fiber.App) {

	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLeads)

}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("fail")
	}

	fmt.Printf("conn")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("database mig")

}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DBConn.Close()

}

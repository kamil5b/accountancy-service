package main

import (
	"fmt"

	"accountancy_service/database"
	"accountancy_service/models"
	"accountancy_service/routes"
	"accountancy_service/utilities"

	"github.com/gofiber/fiber/v2"
)

func SetupTemplate(server_url, db_url, user, password, protocol, db string) {

	database.Connect(
		db_url, user, password, protocol, db,
		&models.User{},
		&models.Account{},
		&models.Store{},
		&models.Item{},
		&models.InvoiceHistory{},
		&models.Inventory{},
		&models.CreditInvoice{},
		&models.DebitInvoice{},
		&models.Transaction{},
	)
	app := fiber.New()
	/*
		origin := utilities.GoDotEnvVariable("VIEW_URL") //ganti view url ini di .env
		app.Use(cors.New(cors.Config{
			AllowCredentials: true,
			AllowOrigins:     []string{origin},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		}))
	*/
	routes.Setup(app)

	err := app.Listen(server_url)
	if err != nil {
		fmt.Println(err)
		fmt.Scan(&err)
	}
}

func main() {
	var (
		acc_url  = utilities.GoDotEnvVariable("ACCOUNTANCY_URL")
		db_url   = utilities.GoDotEnvVariable("DATABASE_URL")
		user     = utilities.GoDotEnvVariable("DATABASE_USER")
		password = utilities.GoDotEnvVariable("DATABASE_PASSWORD")
		protocol = utilities.GoDotEnvVariable("DSN_PROTOCOL")
		db       = utilities.GoDotEnvVariable("DATABASE_NAME")
	)

	SetupTemplate(acc_url, db_url, user, password, protocol, db)
}

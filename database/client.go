package database

import (
	"log"

	"psbank.com/gocloudob/models"

	"github.com/jinzhu/gorm"
)

//Connector variable used for CRUD operation's
var Connector *gorm.DB

//Connect creates MySQL connection
func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	log.Println("Connection was successfull!")
	return nil
}

//Migrate create/updates database table
func Migrate(table *models.Book) {
	Connector.AutoMigrate(&table)
}

func MigrateCustomer(table *models.Customer) {
	Connector.AutoMigrate(&table)
}

func MigrateProduct(table *models.Product) {
	Connector.AutoMigrate(&table)
}

func MigrateCredentials(credentials *models.Credentials) {
	Connector.AutoMigrate(&credentials)
}

func MigrateProfile(table *models.Profiles) {
	Connector.AutoMigrate(&table)
}

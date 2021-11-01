package entity

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db

}
func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("Pharmacy.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//Migrate the schema
	database.AutoMigrate(&Patient{}, &Pharmacist{}, &Receipt{}, &Medicine{}, &PayMedicine{})

	db = database

	password, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)
	//Pharmacist Data
	db.Model(&Pharmacist{}).Create(&Pharmacist{

		Name:     "ฉัตรพัฒน์",
		Email:    "Chattapat@gmail.com",
		Password: string(password),
	})
}

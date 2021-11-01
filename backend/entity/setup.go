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
  	database, err := gorm.Open(sqlite.Open("sa-64.db"), &gorm.Config{})
 	 if err != nil{
    	panic("fail to connect database")
  }

	database.AutoMigrate(
    	&Cashier{},
		&Bill{},
		&Method{},
		&Receipt{},
  	)

  	db = database

	password, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)

	  //Cashier Data
	Ca1 := Cashier{
		Name: "Somsom",
		Email:	"somsom@gmail.com",
		Password: string(password),
	}
	db.Model(&Cashier{}).Create(&Ca1)

	Ca2 := Cashier{
		Name: "Pongpong",
		Email:	"pongpong@gmail.com",
		Password: string(password),
	}
	db.Model(&Cashier{}).Create(&Ca2)

	Me1 := Method{
		Type: "Cash",
	}
	db.Model(&Method{}).Create(&Me1)

	Me2 := Method{
		Type: "Online paymemt",
	}
	db.Model(&Method{}).Create(&Me2)

	Bi1 := Bill{
		PatientRightID: 1,
		ExaminationID: 1,
		Total: 1300,
		CashierID: 1,
	}
	db.Model(&Bill{}).Create(&Bi1)

	Bi2 := Bill{
		PatientRightID: 2,
		ExaminationID: 2,
		Total: 1300,
		CashierID: 2,
	}
	db.Model(&Bill{}).Create(&Bi2)
}
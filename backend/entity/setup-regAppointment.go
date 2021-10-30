package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db

}

func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("sa-G13.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	// Migrate the schema

	database.AutoMigrate(&Patient{}, &Doctor{}, &Clinic{}, &Nurse{}, &Appointment{})

	db = database

	password, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)

	db.Model(&Nurse{}).Create(&Nurse{
		Firstname: "Wimonrat",
		Lastname:  "Kongdee",
		Email:     "wimonrat@gmail.com",
		Password:  string(password),
	})
	db.Model(&Nurse{}).Create(&Nurse{
		Firstname: "Sirapob",
		Lastname:  "Paina",
		Email:     "sirapob@hotmail.com",
		Password:  string(password),
	})

	var wimonrat Nurse
	var sirapob Nurse
	db.Raw("SELECT * FROM nurses WHERE email = ?", "wimonrat@gmail.com").Scan(&wimonrat)
	db.Raw("SELECT * FROM nurses WHERE email = ?", "sirapob@hotmail.com").Scan(&sirapob)

	//---Patient Data
	patient01 := Patient{
		Id_card:            "1111111111111",
		Firstname:          "AAA",
		Lastname:           "BBB",
		Birthdate:          time.Date(2001, 1, 11, 0, 0, 0, 0, time.UTC),
		Age:                20,
		Allergy:            "-",
		Underlying_Disease: "-",
		Gender:             "Male",
		Recorder:           "Phuwadon",
	}
	db.Model(&Patient{}).Create(&patient01)
	patient02 := Patient{
		Id_card:            "2222222222222",
		Firstname:          "CCC",
		Lastname:           "DDD",
		Birthdate:          time.Date(2000, 10, 15, 0, 0, 0, 0, time.UTC),
		Age:                21,
		Allergy:            "-",
		Underlying_Disease: "-",
		Gender:             "Male",
		Recorder:           "Phuwadon",
	}
	db.Model(&Patient{}).Create(&patient02)

	//---Doctor Data
	doctor01 := Doctor{
		Name:  "พญ.ชฎาพร ไทยคณา",
		Email: "chadaporn@gmail.com",
	}
	db.Model(&Doctor{}).Create(&doctor01)
	doctor02 := Doctor{
		Name:  "นพ.สมศักดิ์ คัมคติ",
		Email: "somsak@hotmail.com",
	}
	db.Model(&Doctor{}).Create(&doctor02)

	//--Clinic Data
	clinic01 := Clinic{
		Name: "อายุรกรรม",
	}
	db.Model(&Clinic{}).Create(&clinic01)
	clinic02 := Clinic{
		Name: "สูตินรีเวช",
	}
	db.Model(&Clinic{}).Create(&clinic02)
	clinic03 := Clinic{
		Name: "กุมารเวชกรรม",
	}
	db.Model(&Clinic{}).Create(&clinic03)

	//Appointment1
	db.Model(&Appointment{}).Create(&Appointment{
		Patient:         patient01,
		Doctor:          doctor01,
		Clinic:          clinic01,
		AppointmentTime: time.Date(2021, 11, 01, 0, 0, 0, 0, time.UTC),
		Nurse:           wimonrat,
		Note:            "นัดตัดไหม",
	})

}

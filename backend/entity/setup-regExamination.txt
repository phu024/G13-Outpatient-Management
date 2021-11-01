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
	database, err := gorm.Open(sqlite.Open("SA-Project.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		&Doctor{},
		&Patient{},
		&Clinic{},
		&Disease{},
		&Medicine{},
		&Examination{},
	)

	db = database

	password1, err := bcrypt.GenerateFromPassword([]byte("pk007"), 14)
	password2, err := bcrypt.GenerateFromPassword([]byte("6226404"), 14)
	password3, err := bcrypt.GenerateFromPassword([]byte("taengmo"), 14)

	//Doctor Data
	Phumchai := Doctor{
		Name:     "นพ.ภูมิชัย ศิริพันธ์พรชนะ",
		Email:    "phumchai@gmail.com",
		Password: string(password2),
	}
	db.Model(&Doctor{}).Create(&Phumchai)

	Wichai := Doctor{
		Name:     "นพ.วิชัย ศรีสุรักษ์",
		Email:    "wichai@gmail.com",
		Password: string(password1),
	}
	db.Model(&Doctor{}).Create(&Wichai)

	Anan := Doctor{
		Name:     "พญ.อนันต์ กระเซ็น",
		Email:    "anan@gmail.com",
		Password: string(password3),
	}
	db.Model(&Doctor{}).Create(&Anan)

	// Patient Data
	HarryPotter := Patient{
		IdCard:            "2984009819054",
		FirstName:         "Harry",
		LastName:          "Potter",
		Birthdate:         time.Date(1980, time.July, 31, 0, 0, 0, 0, time.UTC),
		Age:               41,
		Allergy:           "",
		UnderlyingDisease: "",
		Gender:            "Male",
		Recorder:          "TomRiddle",
	}
	db.Model(&Patient{}).Create(&HarryPotter)

	HermioneGranger := Patient{
		IdCard:            "7431049850012",
		FirstName:         "Hermione",
		LastName:          "Granger",
		Birthdate:         time.Date(1979, time.September, 19, 0, 0, 0, 0, time.UTC),
		Age:               42,
		Allergy:           "",
		UnderlyingDisease: "",
		Gender:            "Female",
		Recorder:          "TomRiddle",
	}
	db.Model(&Patient{}).Create(&HermioneGranger)

	RonaldWeasley := Patient{
		IdCard:            "5948730974531",
		FirstName:         "Ronald",
		LastName:          "Weasley",
		Birthdate:         time.Date(1979, time.March, 1, 0, 0, 0, 0, time.UTC),
		Age:               41,
		Allergy:           "",
		UnderlyingDisease: "",
		Gender:            "Male",
		Recorder:          "TomRiddle",
	}
	db.Model(&Patient{}).Create(&RonaldWeasley)

	// Clinic Data
	InternalMedicine := Clinic{
		Name: "อายุรกรรม",
	}
	db.Model(&Clinic{}).Create(&InternalMedicine)

	Dental := Clinic{
		Name: "ทันตกรรม",
	}
	db.Model(&Clinic{}).Create(&Dental)

	Otolaryngology := Clinic{
		Name: "หู คอ จมูก",
	}
	db.Model(&Clinic{}).Create(&Otolaryngology)

	// Disease Data
	Diabetes := Disease{
		Name: "เบาหวาน",
	}
	db.Model(&Disease{}).Create(&Diabetes)

	Cancer := Disease{
		Name: "มะเร็ง",
	}
	db.Model(&Disease{}).Create(&Cancer)

	Cirrhosis := Disease{
		Name: "ตับแข็ง",
	}
	db.Model(&Disease{}).Create(&Cirrhosis)

	NoneDisease := Disease{
		Name: "-",
	}
	db.Model(&Disease{}).Create(&NoneDisease)

	// Medicine Data
	Vitamin := Medicine{
		MedicineName: "วิตามิน",
	}
	db.Model(&Medicine{}).Create(&Vitamin)

	Paracetamol500mg := Medicine{
		MedicineName: "พาราเซตามอล 500 mg",
	}
	db.Model(&Medicine{}).Create(&Paracetamol500mg)

	Expectorants := Medicine{
		MedicineName: "ยาขับเสมหะ",
	}
	db.Model(&Medicine{}).Create(&Expectorants)

	NoneMedicine := Medicine{
		MedicineName: "-",
	}
	db.Model(&Medicine{}).Create(&NoneMedicine)

	// Examination 1
	db.Model(&Examination{}).Create(&Examination{
		Patient:       HarryPotter,
		Doctor:        Phumchai,
		Clinic:        InternalMedicine,
		Disease:       Diabetes,
		Treatment:     "ให้ยา",
		TreatmentCost: 0,
		Medicine:      Vitamin,
		MedicineCost:  200,
		TreatmentTime: time.Now(),
	})
}

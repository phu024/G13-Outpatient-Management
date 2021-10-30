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
	database, err := gorm.Open(sqlite.Open("sa-64.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//Magrate the schema
	database.AutoMigrate(
		&Recorder{}, &Allergy{}, &Underlying_disease{}, &Gender{}, &Patient{},
	)
	db = database

	password, err := bcrypt.GenerateFromPassword([]byte("1234"), 14)

	db.Model(&Recorder{}).Create(&Recorder{
		FirstName: "ภูวดล",
		LastName:  "เดชารัมย์",
		Email:     "phu@email.com",
		Password: string(password),
	})
	db.Model(&Recorder{}).Create(&Recorder{
		FirstName: "แพรวา",
		LastName:  "เดชารัมย์",
		Email:     "phrae@email.com",
		Password: string(password),
	})
	var phuwadon Recorder
	var phrae Recorder
	db.Raw("SELECT * FROM recorders WHERE email = ?", "phu@email.com").Scan(&phuwadon)
	db.Raw("SELECT * FROM recorders WHERE email = ?", "phrae@email.com").Scan(&phrae)

	// Gender Data
	male := Gender{
		Identity: "ชาย",
	}
	db.Model(&Gender{}).Create(&male)
	female := Gender{
		Identity: "หญิง",
	}
	db.Model(&Gender{}).Create(&female)

	//Allergy Data
	a_none := Allergy{
		Information: "ไม่มี",
	}
	db.Model(&Allergy{}).Create(&a_none)
	a_Aspirin := Allergy{
		Information: "Aspirin",
	}
	db.Model(&Allergy{}).Create(&a_Aspirin)
	a_Insulin := Allergy{
		Information: "Insulin",
	}
	db.Model(&Allergy{}).Create(&a_Insulin)
	a_Iodine := Allergy{
		Information: "Iodine",
	}
	db.Model(&Allergy{}).Create(&a_Iodine)

	//Underlying disease Data
	u_none := Underlying_disease{
		Information: "ไม่มี",
	}
	db.Model(&Underlying_disease{}).Create(&u_none)
	u_cancer := Underlying_disease{
		Information: "โรคโลหิตจาง",
	}
	db.Model(&Underlying_disease{}).Create(&u_cancer)
	u_hypertension := Underlying_disease{
		Information: "โรคความดันโลหิตสูง",
	}
	db.Model(&Underlying_disease{}).Create(&u_hypertension)
	u_diabetes := Underlying_disease{
		Information: "โรคเบาหวาน",
	}
	db.Model(&Underlying_disease{}).Create(&u_diabetes)
	u_heart := Underlying_disease{
		Information: "โรคหัวใจ",
	}
	db.Model(&Underlying_disease{}).Create(&u_heart)

	//Patient 1
	db.Model(&Patient{}).Create(&Patient{
		Id_card:            "1111111111111",
		FirstName:          "ปิติวัฒน์",
		LastName:           "เลิศวิทยา",
		Gender:             male,
		Birthdate:          time.Date(2000, 2, 25, 0, 0, 0, 0, time.UTC),
		Age:                21,
		Allergy:            a_none,
		Underlying_disease: u_hypertension,
		Recorder:           phuwadon,
	})
}

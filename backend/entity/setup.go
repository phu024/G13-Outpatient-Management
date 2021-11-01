package entity

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	//"time"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}
func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("schema.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	database.AutoMigrate(
		&Recorder{},
		&Allergy{},
		&Underlying_disease{},
		&Gender{},
		&Patient{},
		&Doctor{},
		&Clinic{},
		&Disease{},
		&Medicine{},
		&Examination{},
		&Nurse{},
		&Appointment{},
		&PatientRight{},
		&Cashier{},
		&Bill{},
		&Method{},
		&Receipt{},
		&Pharmacist{},
		&PayMedicine{},
	)

	db = database

	password, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)

	//ภูวดล
	db.Model(&Recorder{}).Create(&Recorder{
		FirstName: "ภูวดล",
		LastName:  "เดชารัมย์",
		Email:     "phu@email.com",
		Password:  string(password),
	})
	db.Model(&Recorder{}).Create(&Recorder{
		FirstName: "แพรวา",
		LastName:  "เดชารัมย์",
		Email:     "phrae@email.com",
		Password:  string(password),
	})
	var phuwadon Recorder
	var phrae Recorder
	db.Raw("SELECT * FROM recorders WHERE email = ?", "phu@email.com").Scan(&phuwadon)
	db.Raw("SELECT * FROM recorders WHERE email = ?", "phrae@email.com").Scan(&phrae)

	// Gender Data (ข้อมูลเพศ)
	male := Gender{
		Identity: "ชาย",
	}
	db.Model(&Gender{}).Create(&male)
	female := Gender{
		Identity: "หญิง",
	}
	db.Model(&Gender{}).Create(&female)

	//Allergy Data (ข้อมูลแพ้ยา)
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

	//Underlying disease Data (ข้อมูลโรคประจำตัว)
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
	u_allergy := Underlying_disease{
		Information: "โรคภูมิแพ้",
	}
	db.Model(&Underlying_disease{}).Create(&u_allergy)

	//ภูมิชัย
	//Doctor Data (ข้อมูลแพทย์	)
	Phumchai := Doctor{
		Name:     "นพ.ภูมิชัย ศิริพันธ์พรชนะ",
		Email:    "phumchai@gmail.com",
		Password: string(password),
	}
	db.Model(&Doctor{}).Create(&Phumchai)
	Wichai := Doctor{
		Name:     "นพ.วิชัย ศรีสุรักษ์",
		Email:    "wichai@gmail.com",
		Password: string(password),
	}
	db.Model(&Doctor{}).Create(&Wichai)
	Anan := Doctor{
		Name:     "พญ.อนันต์ กระเซ็น",
		Email:    "anan@gmail.com",
		Password: string(password),
	}
	db.Model(&Doctor{}).Create(&Anan)

	// Disease Data (ข้อมูลโรค)
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

	// Clinic Data (ข้อมูลคลินิก)
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

	// Medicine Data (ข้อมูลยา)
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

	//ชฏาพร
	//Nurse Data (ข้อมูลพยาบาลผู้ทำการนัด)
	db.Model(&Nurse{}).Create(&Nurse{
		Firstname: "Wimonrat",
		Lastname:  "Kongdee",
		Email:     "wimonrat@gmail.com",
		Password:  string(password),
	})
	db.Model(&Nurse{}).Create(&Nurse{
		Firstname: "Nipat",
		Lastname:  "Paina",
		Email:     "nipat@hotmail.com",
		Password:  string(password),
	})

	var wimonrat Nurse
	var nipat Nurse
	db.Raw("SELECT * FROM nurses WHERE email = ?", "wimonrat@gmail.com").Scan(&wimonrat)
	db.Raw("SELECT * FROM nurses WHERE email = ?", "nipat@hotmail.com").Scan(&nipat)

	//อนันต์
	//Crashier Data
	Cra1 := Cashier{
		Name:     "อนันต์ กระเซ็น",
		Email:    "anan1234@gmail.com",
		Password: string(password),
	}
	db.Model(&Cashier{}).Create(&Cra1)

	Cra3 := Cashier{
		Name:     "ภูมิชัย สนสวย",
		Email:    "phumchai123@gmail.com",
		Password: string(password),
	}
	db.Model(&Cashier{}).Create(&Cra3)

	//PatientRight Data
	Pr1 := PatientRight{
		Name:     "สิทธิ์สุขภาพถ้วนหน้า",
		Discount: 80,
	}
	db.Model(&PatientRight{}).Create(&Pr1)

	Pr2 := PatientRight{
		Name:     "สิทธิ์นักศึกษา",
		Discount: 50,
	}
	db.Model(&PatientRight{}).Create(&Pr2)

	//ภาคิน
	//Method Data
	Me1 := Method{
		Type: "Cash",
	}
	db.Model(&Method{}).Create(&Me1)

	Me2 := Method{
		Type: "Online paymemt",
	}
	db.Model(&Method{}).Create(&Me2)

}

package entity

import (
	"time"
	"gorm.io/gorm"
)

type Recorder struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Password string
	// 1 Recorder มีได้หลาย Patient
	Patient []Patient `gorm:"foreignKey:RecorderID"`
}

type Allergy struct {
	gorm.Model
	Information string
	// 1 Allergy มีได้หลาย Patient
	Patient []Patient `gorm:"foreignKey:AllergyID"`
}

type Underlying_disease struct {
	gorm.Model
	Information string
	// 1 Underlying_disease มีได้หลาย Patient
	Patient []Patient `gorm:"foreignKey:Underlying_diseaseID"`
}

type Gender struct {
	gorm.Model
	Identity string
	// 1 Gender มีได้หลาย Patient
	Patient []Patient `gorm:"foreignKey:GenderID"`
}

type Patient struct {
	gorm.Model
	Id_card   string
	FirstName string
	LastName  string
	Birthdate time.Time
	Age       uint
	//RecorderID ทำหน้าที่เป็น ForeignKey
	RecorderID *uint
	Recorder   Recorder `gorm:"references:id"`

	//AllergyID ทำหน้าที่เป็น ForeignKey
	AllergyID *uint
	Allergy   Allergy `gorm:"references:id"`

	//Underlying_diseaseID ทำหน้าที่เป็น ForeignKey
	Underlying_diseaseID *uint
	Underlying_disease   Underlying_disease `gorm:"references:id"`

	//GenderID ทำหน้าที่เป็น ForeignKey
	GenderID *uint
	Gender   Gender `gorm:"references:id"`
}

package entity

import (
	"time"

	"gorm.io/gorm"
)

type Doctor struct {
	gorm.Model
	Name        string
	Email       string `gorm:"uniqueIndex"`
	Password    string
	WatchVideos []Examination `gorm:"foreignKey: DoctorID"`
}

type Patient struct {
	gorm.Model
	IdCard            string `gorm:"uniqueIndex"`
	FirstName         string
	LastName          string
	Birthdate         time.Time
	Age               uint
	Allergy           string
	UnderlyingDisease string
	Gender            string
	Recorder          string
	WatchVideos       []Examination `gorm:"foreignKey: PatientID"`
}

type Clinic struct {
	gorm.Model
	Name        string        `gorm:"uniqueIndex"`
	WatchVideos []Examination `gorm:"foreignKey: ClinicID"`
}

type Disease struct {
	gorm.Model
	Name        string        `gorm:"uniqueIndex"`
	WatchVideos []Examination `gorm:"foreignKey: DiseaseID"`
}

type Medicine struct {
	gorm.Model
	MedicineName string        `gorm:"uniqueIndex"`
	WatchVideos  []Examination `gorm:"foreignKey: MedicineID"`
}

type Examination struct {
	gorm.Model
	Treatment     string
	TreatmentCost uint
	MedicineCost  uint
	TreatmentTime time.Time

	// DoctorID ทำหน้าที่เป็น FK
	DoctorID *uint
	Doctor   Doctor `gorm:"references:id"`

	// PatientID ทำหน้าที่เป็น FK
	PatientID *uint
	Patient   Patient `gorm:"references:id"`

	// ClinicID ทำหน้าที่เป็น FK
	ClinicID *uint
	Clinic   Clinic `gorm:"references:id"`

	// DiseaseID ทำหน้าที่เป็น FK
	DiseaseID *uint
	Disease   Disease `gorm:"references:id"`

	// MedicineID ทำหน้าที่เป็น FK
	MedicineID *uint
	Medicine   Medicine `gorm:"references:id"`
}

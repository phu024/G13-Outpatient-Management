package entity

import (
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	ID_Card   string `gorm:"uniqueIndex"`
	FirstName string
	LastName  string
	Birthdate time.Time
	Age       uint8
	// 1 คนไข้ มีได้หลาย บันทึกการจ่ายยา
	PayMedicine []PayMedicine `gorm:"foreignKey:PatientID"`
}
type Pharmacist struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	// 1 เภสัชกร มีได้หลาย บันทึกการจ่ายยา
	PayMedicine []PayMedicine `gorm:"foreignKey:PharmacistID"`
}
type Receipt struct {
	gorm.Model
	Receipt_DateTime time.Time
	// 1 ใบเสร็จ มีได้หลาย บันทึกการจ่ายยา
	PayMedicine []PayMedicine `gorm:"foreignKey:ReceiptID"`
}
type Medicine struct {
	gorm.Model
	Medicine_name string
	// 1 ยา มีได้หลาย บันทึกการจ่ายยา
	PayMedicine []PayMedicine `gorm:"foreignKey:MedicineID"`
}
type PayMedicine struct {
	gorm.Model
	Pay_DateTime    time.Time
	Medicine_Amount uint8

	PatientID *uint
	Patient   Patient `gorm: "references:id"`

	PharmacistID *uint
	Pharmacist   Pharmacist `gorm: "references:id"`

	ReceiptID *uint
	Receipt   Receipt `gorm: "references:id"`

	MedicineID *uint
	Medicine   Medicine `gorm: "references:id"`
}

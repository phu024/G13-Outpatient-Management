package entity

import (
  	"time"
  	"gorm.io/gorm"
)

type Cashier struct {
  	gorm.Model
  	Name          	string
	Email			string		`gorm: "uniqueIndex"`
  	Password      	string      
 	// 1 cashier เป็นเจ้าข้องได้หลาย bill
 	Bills          []Bill     `gorm: "foreignKey:CashierID"`
  	// 1 cashier เป็นเจ้าของได้หลาย receipt
  	Receipts       []Receipt  `gorm: "foreignKey:CashierID"`  
}

type Bill struct {
  	gorm.Model
  	PatientRightID     uint8
  	ExaminationID      uint8
  	Total         uint
  	// ทำหน้าที่เป็น FK
  	CashierID     uint
	Cashier       Cashier	`gorm:"references:id"`
  	// 1 bill อยู่ได้หลาย receipt 
  	Receipt       []Receipt  `gorm: "foreignKey:BillID"`
 	
}

type Method struct {
 	gorm.Model
  	Type          string
  	// 1 method อยู่ได้หลาย receipt
  	Receipt       []Receipt  `gorm: "foreignKey:MethodID"`
}

type Receipt struct {
  	gorm.Model
  	SavedTime      time.Time
  	// ทำหน้าที่เป็น FK
 	CashierID   *uint
	Cashier       Cashier	`gorm:"references:id"`
  	// ทำหน้าที่เป็น FK
  	BillID        *uint
	Bill          Bill		`gorm:"references:id"`
	// ทำหน้าที่เป็น FK
  	MethodID    *uint
  	Method       Method		`gorm:"references:id"`
}

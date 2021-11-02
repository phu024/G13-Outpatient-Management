package entity

import (
	"time"

	"gorm.io/gorm"
)

//ภูวดล
type Recorder struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Password  string
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

	// 1 Patient มีได้หลาย Examination
	Examinations []Examination `gorm:"foreignKey:PatientID"`
	// 1 Patient ทำได้หลาย Appointment
	Appointments []Appointment `gorm:"foreignKey:PatientID"`
	// 1 Patient ทำได้หลาย PayMedicine
	PayMedicine []PayMedicine `gorm:"foreignKey:PatientID"`
}

//ภูมิชัย
type Doctor struct {
	gorm.Model
	Name         string
	Email        string `gorm:"uniqueIndex"`
	Password     string
	Examinations []Examination `gorm:"foreignKey: DoctorID"`
	Appointments []Appointment `gorm:"foreignKey: DoctorID"`
}

type Clinic struct {
	gorm.Model
	Name         string        `gorm:"uniqueIndex"`
	Examinations []Examination `gorm:"foreignKey: ClinicID"`
	Appointments []Appointment `gorm:"foreignKey: ClinicID"`
}

type Disease struct {
	gorm.Model
	Name         string        `gorm:"uniqueIndex"`
	Examinations []Examination `gorm:"foreignKey: DiseaseID"`
}

type Medicine struct {
	gorm.Model
	MedicineName string        `gorm:"uniqueIndex"`
	Examinations []Examination `gorm:"foreignKey: MedicineID"`
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

	Bills      []Bill `gorm:"foreignKey:ExaminationID"`
	MedicineID *uint
	Medicine   Medicine `gorm:"references:id"`
}

//ชฏาพร
type Nurse struct {
	gorm.Model
	Firstname string
	Lastname  string
	Email     string `gorm:"uniqueTndex"`
	Password  string
	// 1 Nurse ทำได้หลาย Appointment
	Appointments []Appointment `gorm:"foreignKey:NurseID"`
}

type Appointment struct {
	gorm.Model

	AppointmentTime time.Time
	Note            string

	//PatientID ทำหน้าที่เป็น FK
	PatientID *uint
	Patient   Patient

	//DoctorID ทำหน้าที่เป็น FK
	DoctorID *uint
	Doctor   Doctor

	//ClinicID ทำหน้าที่เป็น FK
	ClinicID *uint
	Clinic   Clinic

	//NurseID ทำหน้าที่เป็น FK
	NurseID *uint
	Nurse   Nurse
}

//อนันต์

type Cashier struct {
	gorm.Model

	Name string

	Email string `gorm:"uniqueIndex"`

	Password string

	Bills []Bill `gorm:"foreignKey:CashierID"`

	Receipts []Receipt `gorm:"foreignKey:CashierID"`
}

type PatientRight struct {
	gorm.Model

	Name string

	Discount uint

	Bills []Bill `gorm:"foreignKey:PatientRightID"`
}

type Bill struct {
	gorm.Model
	// ทำหน้าที่เป็น FK
	ExaminationID *uint       `gorm:"uniqueIndex"`
	Examination   Examination `gorm:"references:id"`
	// ทำหน้าที่เป็น FK
	PatientRightID *uint
	PatientRight   PatientRight `gorm:"references:id"`

	BillTime time.Time
	Total    uint
	// ทำหน้าที่เป็น FK
	CashierID *uint
	Cashier   Cashier `gorm:"references:id"`

	Receipts []Receipt `gorm:"foreignKey:BillID"`
}

//ภาคิน
type Method struct {
	gorm.Model
	Type string
	// 1 method อยู่ได้หลาย receipt
	Receipt []Receipt `gorm:"foreignKey:MethodID"`
}

type Receipt struct {
	gorm.Model
	SavedTime time.Time
	// ทำหน้าที่เป็น FK
	CashierID *uint
	Cashier   Cashier `gorm:"references:id"`
	// ทำหน้าที่เป็น FK
	BillID *uint
	Bill   Bill `gorm:"references:id"`
	// ทำหน้าที่เป็น FK
	MethodID *uint
	Method   Method `gorm:"references:id"`
}

//รัชอินทร์

type Pharmacist struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	// 1 เภสัชกร มีได้หลาย บันทึกการจ่ายยา
	PayMedicine []PayMedicine `gorm:"foreignKey:PharmacistID"`
}
type PayMedicine struct {
	gorm.Model
	Pay_DateTime    time.Time
	Medicine_Amount uint8

	PatientID *uint
	Patient   Patient `gorm:"references:id"`

	PharmacistID *uint
	Pharmacist   Pharmacist `gorm:"references:id"`

	ReceiptID *uint
	Receipt   Receipt `gorm:"references:id"`

	MedicineID *uint
	Medicine   Medicine `gorm:"references:id"`
}

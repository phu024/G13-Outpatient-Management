package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phu024/G13-Outpatient-Management/entity"
)

// POST /examinations
func CreateExamiation(c *gin.Context) {

	var clinic entity.Clinic
	var disease entity.Disease
	var doctor entity.Doctor
	var medicine entity.Medicine
	var patient entity.Patient
	var examination entity.Examination

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 10 จะถูก bind เข้าตัวแปร examination
	if err := c.ShouldBindJSON(&examination); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 11: ค้นหา patient ด้วย id
	if tx := entity.DB().Where("id = ?", examination.PatientID).First(&patient); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "patient not found"})
		return
	}

	// 12: ค้นหา doctor ด้วย id
	if tx := entity.DB().Where("id = ?", examination.DoctorID).First(&doctor); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "doctor not found"})
		return
	}

	// 13: ค้นหา clinic ด้วย id
	if tx := entity.DB().Where("id = ?", examination.ClinicID).First(&clinic); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "clinic not found"})
		return
	}

	// 14: ค้นหา disease ด้วย id
	if tx := entity.DB().Where("id = ?", examination.DiseaseID).First(&disease); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "disease not found"})
		return
	}

	// 15: ค้นหา medicine ด้วย id
	if tx := entity.DB().Where("id = ?", examination.MedicineID).First(&medicine); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "medicine not found"})
		return
	}
	// 16: สร้าง Examination
	ex := entity.Examination{
		Patient:       patient,                   // โยงความสัมพันธ์กับ Entity Patient
		Doctor:        doctor,                    // โยงความสัมพันธ์กับ Entity Doctor
		Clinic:        clinic,                    // โยงความสัมพันธ์กับ Entity Clinic
		Disease:       disease,                   // โยงความสัมพันธ์กับ Entity Disease
		Medicine:      medicine,                  // โยงความสัมพันธ์กับ Entity Medicine
		Treatment:     examination.Treatment,     // ตั้งค่าฟิลด์ treatment
		TreatmentCost: examination.TreatmentCost, // ตั้งค่าฟิลด์ treatmentCost
		MedicineCost:  examination.MedicineCost,  // ตั้งค่าฟิลด์ medicineCost
		TreatmentTime: examination.TreatmentTime, // ตั้งค่าฟิลด์ treatmentTime
	}

	// 17: บันทึก
	if err := entity.DB().Create(&ex).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ex})
}

// GET /examination/:id
func GetExamination(c *gin.Context) {
	var examination entity.Examination
	id := c.Param("id")
	if err := entity.DB().Preload("Patient").Preload("Doctor").Preload("Clinic").Preload("Disease").Preload("Medicine").Raw("SELECT * FROM examinations WHERE id = ?", id).Find(&examination).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": examination})
}

// GET /examinations
func ListExaminations(c *gin.Context) {
	var examinations []entity.Examination
	if err := entity.DB().Preload("Patient").Preload("Doctor").Preload("Clinic").Preload("Disease").Preload("Medicine").Raw("SELECT * FROM examinations").Find(&examinations).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": examinations})
}

// DELETE /examinations/:id
func DeleteExamination(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM examinations WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "examination not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /examinations
func UpdateExamination(c *gin.Context) {
	var examination entity.Examination
	if err := c.ShouldBindJSON(&examination); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", examination.ID).First(&examination); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "examination not found"})
		return
	}

	if err := entity.DB().Save(&examination).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": examination})
}

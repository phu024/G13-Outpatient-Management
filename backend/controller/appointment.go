package controller

import (
	"net/http"

	"github.com/phu024/G13-Outpatient-Management/entity"
	"github.com/gin-gonic/gin"
)

// POST /appointment
func CreateAppointment(c *gin.Context) {

	var appointment entity.Appointment
	var nurse entity.Nurse
	var doctor entity.Doctor
	var patient entity.Patient
	var clinic entity.Clinic

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร appointment
	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 11: ค้นหา nurse ด้วย id
	if tx := entity.DB().Where("id = ?", appointment.NurseID).First(&nurse); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nurse not found"})
		return
	}
	// 9: ค้นหา patient ด้วย id
	if tx := entity.DB().Where("id = ?", appointment.PatientID).First(&patient); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "patient not found"})
		return
	}

	// 10: ค้นหา doctor ด้วย id
	if tx := entity.DB().Where("id = ?", appointment.DoctorID).First(&doctor); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "doctor not found"})
		return
	}

	// 11: ค้นหา clinic ด้วย id
	if tx := entity.DB().Where("id = ?", appointment.ClinicID).First(&clinic); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "clinic not found"})
		return
	}
	// 12: สร้าง Appointment
	apm := entity.Appointment{
		Patient:         patient,                     // โยงความสัมพันธ์กับ Entity Patient
		Doctor:          doctor,                      // โยงความสัมพันธ์กับ Entity Doctor
		Clinic:          clinic,                      // โยงความสัมพันธ์กับ Entity Clinic
		Nurse:           nurse,                       // โยงความสัมพันธ์กับ Entity Nurse (AppointmentBY_Nurse)
		AppointmentTime: appointment.AppointmentTime, // ตั้งค่าฟิลด์ appointmentTime
		Note:            appointment.Note,            // ตั้งค่าฟิลด์ note
	}

	// 13: บันทึก
	if err := entity.DB().Create(&apm).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": apm})
}

// GET /appointment/:id
func GetAppointment(c *gin.Context) {
	var appointment entity.Appointment
	id := c.Param("id")
	if err := entity.DB().Preload("Patient").Preload("Doctor").Preload("Clinic").Preload("Nurse").Raw("SELECT * FROM appointments WHERE id = ?", id).Find(&appointment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": appointment})
}

// GET /appointment
func ListAppointments(c *gin.Context) {
	var appointments []entity.Appointment
	if err := entity.DB().Preload("Patient").Preload("Doctor").Preload("Clinic").Preload("Nurse").Raw("SELECT * FROM appointments").Find(&appointments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": appointments})
}

// DELETE /appointments/:id
func DeleteAppointment(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM appointments WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "appointment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /appointments
func UpdateAppointment(c *gin.Context) {
	var appointment entity.Appointment
	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", appointment.ID).First(&appointment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "appointment not found"})
		return
	}

	if err := entity.DB().Save(&appointment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": appointment})
}

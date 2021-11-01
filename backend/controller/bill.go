package controller

import (
	"net/http"

	

	"github.com/pechkr2020/sa-project/entity"
	"github.com/gin-gonic/gin"
)

// POST /bills
func CreateBill(c *gin.Context) {

	var examination entity.Examination
	var patientright entity.PatientRight
	var cashier entity.Cashier
	var bill entity.Bill
	var check_exam entity.Bill

	
	
	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร bill
	if err := c.ShouldBindJSON(&bill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// เช็คการบันทึก bill ผลการรักษาซ้ำ  ถ้ามีบิลผลการรักษาที่ซ้ำกัน ให้ return บิลนั้นออกไป
	if tx := entity.DB().Table("bills").Where("examination_id = ?", bill.ExaminationID).First(&check_exam); tx.RowsAffected != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"billDuplicate": check_exam})
		return
	}

	// 9: ค้นหา examination ด้วย id
	if tx := entity.DB().Where("id = ?", bill.ExaminationID).First(&examination); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "examination not found"})
		return
	}

	// 10: ค้นหา patientright ด้วย id
	if tx := entity.DB().Where("id = ?", bill.PatientRightID).First(&patientright); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "resolution not found"})
		return
	}

	// 11: ค้นหา crashier ด้วย id
	if tx := entity.DB().Where("id = ?", bill.CashierID).First(&cashier); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cashier not found"})
		return
	}

	
		
	// 12: สร้าง Bill
	bl := entity.Bill{
		Examination: examination,             // โยงความสัมพันธ์กับ Entity Examination
		PatientRight:       patientright,                  // โยงความสัมพันธ์กับ Entity PatientRight
		Cashier:    cashier,               // โยงความสัมพันธ์กับ Entity Crashier
		BillTime: bill.BillTime, // ตั้งค่าฟิลด์ BillTime
		Total:(examination.Treatment_cost+examination.Medicine_cost)-patientright.Discount,//ตั้งค่าฟิลด์ Total
	}

	// 13: บันทึก
	if err := entity.DB().Create(&bl).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": bl})
}

// GET /bill/:id
func GetBill(c *gin.Context) {
	var bill entity.Bill
	id := c.Param("id")
	if err := entity.DB().Preload("Examination").Preload("PatientRight").Preload("Cashier").Raw("SELECT * FROM bills WHERE id = ?", id).Find(&bill).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": bill})
}

// GET /bills
func ListBills(c *gin.Context) {
	var bills []entity.Bill
	if err := entity.DB().Preload("Examination").Preload("PatientRight").Preload("Cashier").Raw("SELECT * FROM bills").Find(&bills).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bills})
}

// DELETE /bills/:id
func DeleteBill(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM bills WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bill not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /bills
func UpdateBill(c *gin.Context) {
	var bill entity.Bill
	if err := c.ShouldBindJSON(&bill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", bill.ID).First(&bill); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bill not found"})
		return
	}

	if err := entity.DB().Save(&bill).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bill})
}
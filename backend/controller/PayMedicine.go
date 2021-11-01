package controller

import (
	"net/http"

	"github.com/Chattapat/Pharmacy/entity"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

// POST /PayMedicine
func CreatePayMedicine(c *gin.Context) {

	var paymedicine entity.PayMedicine
	var patient entity.Patient
	var medicine entity.Medicine
	var receipt entity.Receipt

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร PayMedicine
	if err := entity.DB().Create(&paymedicine).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&paymedicine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา patient ด้วย id
	if tx := entity.DB().Where("id = ?", paymedicine.PatientID).First(&patient); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID not found"})
		return
	}

	// 10: ค้นหา medicine ด้วย id
	if tx := entity.DB().Where("id = ?", paymedicine.MedicineID).First(&medicine); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Medicine not found"})
		return
	}

	// 11: ค้นหา receipt ด้วย id
	if tx := entity.DB().Where("id = ?", paymedicine.ReceiptID).First(&receipt); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Receipt not found"})
		return
	}

	// 12: สร้าง PayMedicine
	pm := entity.PayMedicine{
		Model: gorm.Model{},

		Pay_DateTime:    paymedicine.Pay_DateTime,
		Medicine_Amount: paymedicine.Medicine_Amount,
		PatientID:       new(uint),
		PharmacistID:    new(uint),
		ReceiptID:       new(uint),
		MedicineID:      new(uint),
	}
	// 13: บันทึก
	if err := entity.DB().Create(&pm).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	c.JSON(http.StatusOK, gin.H{"data": pm})
}

// GET /paymedicine
// List all paymedicine
func ListPayMedicine(c *gin.Context) {
	var paymedicine []entity.PayMedicine
	if err := entity.DB().Raw("SELECT * FROM paymedicines").Scan(&paymedicine).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": paymedicine})
}

// GET /paymedicine/:id
// Get paymedicine by id
func GetPayMedicine(c *gin.Context) {
	var paymedicine entity.PayMedicine
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM paymedicines WHERE id = ?", id).Scan(&paymedicine).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": paymedicine})
}

// DELETE /paymedicines/:id
func DeletePayMedicine(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM paymedicines WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "paymedicine not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /paymedicine
func UpdatePayMedicine(c *gin.Context) {
	var paymedicine entity.PayMedicine
	if err := c.ShouldBindJSON(&paymedicine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", paymedicine.ID).First(&paymedicine); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "paymedicine not found"})
		return
	}
	if err := entity.DB().Save(&paymedicine).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": paymedicine})
}

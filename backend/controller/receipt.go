package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phu024/G13-Outpatient-Management/entity"
)

// POST /receipts
func CreateReceipt(c *gin.Context) {

	var receipt entity.Receipt
	var cashier entity.Cashier
	var bill entity.Bill
	var method entity.Method

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร receipt
	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา cashier ด้วย id
	if tx := entity.DB().Where("id = ?", receipt.CashierID).First(&cashier); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cashier not found"})
		return
	}

	// 10: ค้นหา bill ด้วย id
	if tx := entity.DB().Where("id = ?", receipt.BillID).First(&bill); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bill not found"})
		return
	}

	// 11: ค้นหา method ด้วย id
	if tx := entity.DB().Where("id = ?", receipt.MethodID).First(&method); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "method not found"})
		return
	}
	// 12: สร้าง Receipt
	rc := entity.Receipt{
		Cashier:   cashier,           // โยงความสัมพันธ์กับ Entity Cashier
		Bill:      bill,              // โยงความสัมพันธ์กับ Entity Bill
		Method:    method,            // โยงความสัมพันธ์กับ Entity Method
		SavedTime: receipt.SavedTime, // ตั้งค่าฟิลด์ savedTime
	}

	// 13: บันทึก
	if err := entity.DB().Create(&rc).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": rc})
}

// GET /receipt/:id
func GetReceipt(c *gin.Context) {
	var receipt entity.Receipt
	id := c.Param("id")
	if err := entity.DB().Preload("Cashier").Preload("Bill").Preload("Method").Raw("SELECT * FROM receipts WHERE id = ?", id).Find(&receipt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": receipt})
}

// GET /receipts
func ListReceipts(c *gin.Context) {
	var receipts []entity.Receipt
	if err := entity.DB().Preload("Bill").Preload("Cashier").Preload("Method").Raw("SELECT * FROM receipts").Find(&receipts).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": receipts})
}

// DELETE /receipts/:id
func DeleteReceipt(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM receipts WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "receipt not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /receipts
func UpdateReceipt(c *gin.Context) {
	var receipt entity.Receipt
	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", receipt.ID).First(&receipt); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "receipt not found"})
		return
	}

	if err := entity.DB().Save(&receipt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": receipt})
}

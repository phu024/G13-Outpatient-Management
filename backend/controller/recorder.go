package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phu024/sa-64/entity"
	"golang.org/x/crypto/bcrypt"
)

// GET /recorder
// List all recorder
func ListRecorder(c *gin.Context) {
	var recorders []entity.Recorder
	if err := entity.DB().Raw("SELECT * FROM recorders").Scan(&recorders).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": recorders})
}

// GET /recorder/:id
// Get recorder by id
func GetRecorder(c *gin.Context) {
	var recorder entity.Recorder
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM recorders WHERE id = ?", id).Scan(&recorder).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": recorder})
}

// POST /recorder
func CreateRecorders(c *gin.Context) {
	var recorder entity.Recorder
	if err := c.ShouldBindJSON(&recorder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// เข้ารหัสลับรหัสผ่านที่ผู้ใช้กรอกก่อนบันทึกลงฐานข้อมูล
	bytes, err := bcrypt.GenerateFromPassword([]byte(recorder.Password), 14)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error hashing password"})
		return
	}
	recorder.Password = string(bytes)

	if err := entity.DB().Create(&recorder).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": recorder})
}

// PATCH /recorder
func UpdateRecorder(c *gin.Context) {
	var recorder entity.Recorder
	if err := c.ShouldBindJSON(&recorder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", recorder.ID).First(&recorder); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	if err := entity.DB().Save(&recorder).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": recorder})
}

// DELETE /recorder/:id
func DeleteRecorder(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM recorder WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}
	/*
		if err := entity.DB().Where("id = ?", id).Delete(&entity.Recorder{}).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}*/

	c.JSON(http.StatusOK, gin.H{"data": id})

}


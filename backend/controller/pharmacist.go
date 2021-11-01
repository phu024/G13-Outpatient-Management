package controller

import (
	"net/http"

	"github.com/Chattapat/Pharmacy/entity"
	"github.com/gin-gonic/gin"
)

// GET /pharmacist
// List all pharcist
func ListPharmacist(c *gin.Context) {
	var pharmacist []entity.Pharmacist
	if err := entity.DB().Raw("SELECT * FROM pharmacist").Scan(&pharmacist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": pharmacist})
}

// GET /pharmacist/:id
// Get pharmacist by id
func GetPharmacist(c *gin.Context) {
	var pharmacist entity.Pharmacist
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM pharmacists WHERE id = ?", id).Scan(&pharmacist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": pharmacist})
}

// POST /users
func CreatePharmacist(c *gin.Context) {
	var pharmacist entity.Pharmacist
	if err := c.ShouldBindJSON(&pharmacist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&pharmacist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pharmacist})
}

// PATCH /pharmacist
func UpdatePharmacist(c *gin.Context) {
	var pharmacist entity.Pharmacist
	if err := c.ShouldBindJSON(&pharmacist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", pharmacist.ID).First(&pharmacist); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "pharmacistID not found"})
		return
	}
	if err := entity.DB().Save(&pharmacist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": pharmacist})
}

// DELETE /users/:id
func DeletePharmacist(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM pharmacists WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "pharmacist not found"})
		return
	}
	/*
		if err := entity.DB().Where("id = ?", id).Delete(&entity.Pharmacist{}).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}*/

	c.JSON(http.StatusOK, gin.H{"data": id})
}

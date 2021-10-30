package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phu024/G13-Outpatient-Management/entity"
)

// POST /allergy
func CreateAllergys(c *gin.Context) {
	var allergy entity.Allergy
	if err := c.ShouldBindJSON(&allergy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&allergy).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": allergy})
}

// GET /allergy/:id
func GetAllergy(c *gin.Context) {
	var allergy entity.Allergy
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM allergys WHERE id = ?", id).Scan(&allergy).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": allergy})
}

// GET /allergy
func ListAllergys(c *gin.Context) {
	var allergys []entity.Allergy
	if err := entity.DB().Raw("SELECT * FROM allergies").Scan(&allergys).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": allergys})
}

// DELETE /allergy/:id
func DeleteAllergy(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM allergies WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "resolution not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /allergy
func UpdateAllergy(c *gin.Context) {
	var allergy entity.Allergy
	if err := c.ShouldBindJSON(&allergy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", allergy.ID).First(&allergy); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "allergy not found"})
		return
	}

	if err := entity.DB().Save(&allergy).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": allergy})
}

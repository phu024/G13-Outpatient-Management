package controller

import (
	"net/http"
	"github.com/phu024/sa-64/entity"
	"github.com/gin-gonic/gin"
)

// POST /underlying_diseases
func CreateUnderlying_diseases(c *gin.Context) {
	var underlying_disease entity.Underlying_disease
	if err := c.ShouldBindJSON(&underlying_disease); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&underlying_disease).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": underlying_disease})
}

// GET /underlying_disease/:id
func GetUnderlying_disease(c *gin.Context) {
	var underlying_disease entity.Underlying_disease
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM underlying_diseases WHERE id = ?", id).Scan(&underlying_disease).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": underlying_disease})
}

// GET /underlying_diseases
func ListUnderlying_diseases(c *gin.Context) {
	var underlying_diseases []entity.Underlying_disease
	if err := entity.DB().Raw("SELECT * FROM underlying_diseases").Scan(&underlying_diseases).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": underlying_diseases})
}

// DELETE /underlying_diseases/:id
func DeleteUnderlying_disease(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM underlying_diseases WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "underlying disease not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /underlying_diseases
func UpdateUnderlying_disease(c *gin.Context) {
	var underlying_disease entity.Underlying_disease
	if err := c.ShouldBindJSON(&underlying_disease); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", underlying_disease.ID).First(&underlying_disease); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "underlying disease not found"})
		return
	}

	if err := entity.DB().Save(&underlying_disease).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": underlying_disease})
}
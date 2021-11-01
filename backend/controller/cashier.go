package controller

import (
	"net/http"

	"github.com/pechkr2020/sa-project/entity"
	"github.com/gin-gonic/gin"
)

// GET /cashier
// List all cashier
func ListCashiers(c *gin.Context) {
	var cashiers []entity.Cashier
	if err := entity.DB().Raw("SELECT * FROM cashiers").Scan(&cashiers).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cashiers})
}

// GET /cashier/:id
// Get cashier by id
func GetCashier(c *gin.Context) {
	var cashier entity.Cashier
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM cashiers WHERE id = ?", id).Scan(&cashier).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cashier})
}

// POST /cashiers
func CreateUser(c *gin.Context) {
	var cashier entity.Cashier
	if err := c.ShouldBindJSON(&cashier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&cashier).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cashier})
}

// PATCH /cashiers
func UpdateUser(c *gin.Context) {
	var cashier entity.Cashier
	if err := c.ShouldBindJSON(&cashier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", cashier.ID).First(&cashier); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cashier not found"})
		return
	}

	if err := entity.DB().Save(&cashier).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cashier})
}

// DELETE /cashiers/:id
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM cashiers WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cashier not found"})
		return
	}
	/*
		if err := entity.DB().Where("id = ?", id).Delete(&entity.User{}).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}*/

	c.JSON(http.StatusOK, gin.H{"data": id})
}
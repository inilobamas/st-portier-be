package controllers

import (
	"net/http"
	"st-portier-be/config"
	"st-portier-be/models"

	"github.com/gin-gonic/gin"
)

// Get all companies
func GetCompanies(c *gin.Context) {
	var companies []models.Company
	config.DB.Find(&companies)
	c.JSON(http.StatusOK, gin.H{"data": companies})
}

// Get single company by ID
func GetCompany(c *gin.Context) {
	var company models.Company
	id := c.Param("id")

	if err := config.DB.First(&company, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": company})
}

// Create new company
func CreateCompany(c *gin.Context) {
	var input struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	company := models.Company{
		Name:        input.Name,
		Description: input.Description,
	}
	config.DB.Create(&company)

	c.JSON(http.StatusOK, gin.H{"data": company})
}

// Update company
func UpdateCompany(c *gin.Context) {
	var company models.Company
	id := c.Param("id")

	if err := config.DB.First(&company, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found!"})
		return
	}

	var input struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	company.Name = input.Name
	company.Description = input.Description

	config.DB.Save(&company)
	c.JSON(http.StatusOK, gin.H{"data": company})
}

// Delete company
func DeleteCompany(c *gin.Context) {
	var company models.Company
	id := c.Param("id")

	if err := config.DB.First(&company, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found!"})
		return
	}

	config.DB.Delete(&company)
	c.JSON(http.StatusOK, gin.H{"message": "Company deleted successfully"})
}

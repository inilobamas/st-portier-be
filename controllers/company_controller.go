package controllers

import (
	"net/http"
	"st-portier-be/models"
	"st-portier-be/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Get all companies
func GetCompanies(c *gin.Context) {
	user, _ := c.Get("user") // Get the currently logged-in user
	roleID := user.(models.User).RoleID
	companyID := user.(models.User).CompanyID
	var companies []models.Company
	var err error

	switch roleID {
	case models.SuperAdminRoleID:
		// Super Admin can access all users across all companies
		companies, err = services.GetAllCompanies()
	case models.AdminRoleID, models.NormalUserRoleID:
		// Admin and Normal User can only access users within their company
		companies, err = services.GetCompaniesByID(companyID)
	default:
		// If no permissions, return access denied
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": companies})
}

// GetCompany allows Admin and Normal User to view their company, Super Admin can view any company
func GetCompany(c *gin.Context) {
	user, _ := c.Get("user")
	companyID, _ := strconv.Atoi(c.Param("id"))

	var company *models.Company
	var err error

	// Super Admin can view any company
	if user.(models.User).RoleID == models.SuperAdminRoleID {
		company, err = services.GetCompanyByID(companyID)
	} else {
		// Admin and Normal User can only view their own company
		company, err = services.GetCompanyByIDAndUserCompany(user.(models.User).CompanyID, companyID)
	}

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": company})
}

// CreateCompany allows Super Admin to create a new company
func CreateCompany(c *gin.Context) {
	user, _ := c.Get("user")

	// Only Super Admin can create companies
	if user.(models.User).RoleID != models.SuperAdminRoleID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	var input models.Company
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the service to create a new company
	if err := services.CreateCompany(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create company"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Company created successfully", "data": input})
}

// UpdateCompany allows Admin to update their company, Super Admin can update any company
func UpdateCompany(c *gin.Context) {
	user, _ := c.Get("user")
	companyID, _ := strconv.Atoi(c.Param("id"))

	var input models.Company
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var err error
	if user.(models.User).RoleID == models.SuperAdminRoleID {
		// Super Admin can update any company
		err = services.UpdateCompany(companyID, &input)
	} else if user.(models.User).RoleID == models.AdminRoleID && user.(models.User).CompanyID == companyID {
		// Admin can update only their own company
		err = services.UpdateCompany(companyID, &input)
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update company"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Company updated successfully", "data": input})
}

// DeleteCompany allows only Super Admin to delete a company
func DeleteCompany(c *gin.Context) {
	user, _ := c.Get("user")
	companyID, _ := strconv.Atoi(c.Param("id"))

	// Only Super Admin can delete companies
	if user.(models.User).RoleID != models.SuperAdminRoleID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	if err := services.DeleteCompany(companyID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete company"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Company deleted successfully"})
}

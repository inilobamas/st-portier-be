package controllers

import (
	"net/http"
	"st-portier-be/models"
	"st-portier-be/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateEmployee allows Admin or Super Admin to create a new employee
func CreateEmployee(c *gin.Context) {
	user, _ := c.Get("user")

	// Only Admin or Super Admin can create employees
	if user.(models.User).RoleID != models.SuperAdminRoleID && user.(models.User).RoleID != models.AdminRoleID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	var input models.Employee
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the company ID of the employee to the logged-in user's company
	input.CompanyID = user.(models.User).CompanyID

	// Call the service to create the employee
	if err := services.CreateEmployee(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create employee"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee created successfully", "data": input})
}

// GetEmployeeByID retrieves an employee by their ID
func GetEmployeeByID(c *gin.Context) {
	user, _ := c.Get("user")
	employeeID, _ := strconv.Atoi(c.Param("id"))

	employee, err := services.GetEmployeeByID(employeeID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	// Only allow access if the employee belongs to the same company or user is Super Admin
	if user.(models.User).RoleID != models.SuperAdminRoleID && employee.CompanyID != user.(models.User).CompanyID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": employee})
}

// GetAllEmployeesByCompany fetches all employees for the current user's company
func GetAllEmployeesByCompany(c *gin.Context) {
	user, _ := c.Get("user")

	// Get all employees for the user's company
	employees, err := services.GetAllEmployeesByCompany(user.(models.User).CompanyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch employees"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": employees})
}

// GetAllEmployees retrieves all employees (Super Admin only)
func GetAllEmployees(c *gin.Context) {
	user, _ := c.Get("user")

	// Only Super Admin can retrieve all employees across all companies
	if user.(models.User).RoleID != models.SuperAdminRoleID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Get all employees
	employees, err := services.GetAllEmployees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch employees"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": employees})
}

// UpdateEmployee allows Admin or Super Admin to update an employee's details
func UpdateEmployee(c *gin.Context) {
	// user, _ := c.Get("user")
	employeeID, _ := strconv.Atoi(c.Param("id"))

	var input models.Employee
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the service to update the employee
	if err := services.UpdateEmployee(employeeID, &input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update employee"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee updated successfully", "data": input})
}

// DeleteEmployee allows Super Admin to delete an employee
func DeleteEmployee(c *gin.Context) {
	user, _ := c.Get("user")
	employeeID, _ := strconv.Atoi(c.Param("id"))

	// Only Super Admin can delete employees
	if user.(models.User).RoleID != models.SuperAdminRoleID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	if err := services.DeleteEmployee(employeeID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete employee"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}

package controllers

import (
	"net/http"
	"st-portier-be/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AssignKeyCopy assigns a key copy to an employee
func AssignKeyCopy(c *gin.Context) {
	employeeID, _ := strconv.Atoi(c.Param("employee_id"))
	keyCopyID, _ := strconv.Atoi(c.Param("key_copy_id"))

	// Call the service to assign the key copy
	if err := services.AssignKeyCopy(uint(employeeID), uint(keyCopyID)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Key copy assigned successfully"})
}

// RevokeKeyCopy revokes a key copy from an employee
func RevokeKeyCopy(c *gin.Context) {
	employeeID, _ := strconv.Atoi(c.Param("employee_id"))
	keyCopyID, _ := strconv.Atoi(c.Param("key_copy_id"))

	// Call the service to revoke the key copy
	if err := services.RevokeKeyCopy(uint(employeeID), uint(keyCopyID)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Key copy revoked successfully"})
}

// GetAssignedKeys fetches all keys assigned to an employee
func GetAssignedKeys(c *gin.Context) {
	employeeID, _ := strconv.Atoi(c.Param("employee_id"))

	// Fetch the keys assigned to the employee
	assignments, err := services.GetAssignedKeysForEmployee(uint(employeeID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch assigned keys"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": assignments})
}

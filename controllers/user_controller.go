package controllers

import (
	"net/http"
	"st-portier-be/config"
	"st-portier-be/models"
	"st-portier-be/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// GetUsers lists users based on the role of the requesting user
func GetUsers(c *gin.Context) {
	user, _ := c.Get("user") // Get the currently logged-in user
	roleID := user.(models.User).RoleID
	companyID := user.(models.User).CompanyID

	var users []models.User
	var err error

	switch roleID {
	case models.SuperAdminRoleID:
		// Super Admin can access all users across all companies
		users, err = services.GetAllUsers()
	case models.AdminRoleID, models.NormalUserRoleID:
		// Admin and Normal User can only access users within their company
		users, err = services.GetUsersByCompany(companyID)
	default:
		// If no permissions, return access denied
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// Get single user by ID
func GetUser(c *gin.Context) {
	user, _ := c.Get("user") // Get the currently logged-in user
	strUserID := c.Param("id")
	roleID := user.(models.User).RoleID

	userID, _ := strconv.Atoi(strUserID)

	// Fetch the user by ID
	fetchedUser, err := services.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Super Admin can access any user
	if roleID == models.SuperAdminRoleID {
		c.JSON(http.StatusOK, gin.H{"data": fetchedUser})
		return
	}

	// Admin and Normal User can only access users within their company
	if (roleID == models.AdminRoleID || roleID == models.NormalUserRoleID) && fetchedUser.CompanyID == user.(models.User).CompanyID {
		c.JSON(http.StatusOK, gin.H{"data": fetchedUser})
		return
	}

	// If no permissions, return access denied
	c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
}

// Create new user
func CreateUser(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Role     string `json:"role" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 8)

	user := models.User{
		Username: input.Username,
		Password: string(hashedPassword),
	}
	config.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// Update user
func UpdateUser(c *gin.Context) {
	var input models.User
	user, _ := c.Get("user") // Get the currently logged-in user
	strUserID := c.Param("id")

	userID, _ := strconv.Atoi(strUserID)

	// Bind the JSON input to the user model
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Fetch the user by ID
	fetchedUser, err := services.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Super Admin can update any user
	if user.(models.User).RoleID == models.SuperAdminRoleID {
		if err := services.UpdateUser(userID, &input); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update user"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
		return
	}

	// Admin can only update users within their company
	if user.(models.User).RoleID == models.AdminRoleID && fetchedUser.CompanyID == user.(models.User).CompanyID {
		if err := services.UpdateUser(userID, &input); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update user"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
		return
	}

	// Normal Users can only update their own profile
	if user.(models.User).RoleID == models.NormalUserRoleID && user.(models.User).ID == fetchedUser.ID {
		if err := services.UpdateUser(userID, &input); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update user"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
		return
	}

	// If no permissions, return access denied
	c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
}

// Delete user
func DeleteUser(c *gin.Context) {
	user, _ := c.Get("user") // Get the currently logged-in user
	strUserID := c.Param("id")

	userID, _ := strconv.Atoi(strUserID)

	// Fetch the user by ID
	fetchedUser, err := services.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Super Admin can delete any user
	if user.(models.User).RoleID == models.SuperAdminRoleID {
		if err := services.DeleteUser(userID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete user"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
		return
	}

	// Admin can only delete users within their company
	if user.(models.User).RoleID == models.AdminRoleID && fetchedUser.CompanyID == user.(models.User).CompanyID {
		if err := services.DeleteUser(userID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete user"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
		return
	}

	// Normal Users cannot delete other users
	c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
}

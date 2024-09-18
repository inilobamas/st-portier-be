package routes

import (
	"st-portier-be/controllers"
	"st-portier-be/middleware"
	"st-portier-be/models"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()

	// Authentication
	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)

	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{

		// User CRUD routes
		authorized.POST("/users", middleware.RequireRole(models.SuperAdminRoleID, models.AdminRoleID), controllers.CreateUser)
		authorized.GET("/users/:id", middleware.RequireRole(models.SuperAdminRoleID, models.AdminRoleID, models.NormalUserRoleID), controllers.GetUser)
		authorized.GET("/users", middleware.RequireRole(models.SuperAdminRoleID, models.AdminRoleID, models.NormalUserRoleID), controllers.GetUsers)
		authorized.PUT("/users/:id", middleware.RequireRole(models.SuperAdminRoleID, models.AdminRoleID, models.NormalUserRoleID), controllers.UpdateUser)
		authorized.DELETE("/users/:id", middleware.RequireRole(models.SuperAdminRoleID, models.AdminRoleID), controllers.DeleteUser)

		// Company CRUD routes
		authorized.POST("/companies", middleware.RequireRole(models.SuperAdminRoleID), controllers.CreateCompany)
		authorized.GET("/companies/:id", middleware.RequireRole(models.SuperAdminRoleID, models.AdminRoleID, models.NormalUserRoleID), controllers.GetCompany)
		authorized.GET("/companies", middleware.RequireRole(models.SuperAdminRoleID), controllers.GetCompanies)
		authorized.PUT("/companies/:id", middleware.RequireRole(models.SuperAdminRoleID, models.AdminRoleID), controllers.UpdateCompany)
		authorized.DELETE("/companies/:id", middleware.RequireRole(models.SuperAdminRoleID), controllers.DeleteCompany)
	}

	return r
}

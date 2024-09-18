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

		// Building CRUD routes
		authorized.POST("/buildings", middleware.RequireRole(models.AdminRoleID, models.SuperAdminRoleID), controllers.CreateBuilding)
		authorized.GET("/buildings/:id", middleware.RequireRole(models.AdminRoleID, models.NormalUserRoleID), controllers.GetBuilding)
		authorized.GET("/buildings", middleware.RequireRole(models.SuperAdminRoleID), controllers.GetBuildings)
		authorized.GET("/buildings/:building_id/floors", middleware.RequireRole(models.AdminRoleID, models.NormalUserRoleID, models.SuperAdminRoleID), controllers.GetFloorsByBuildingID)
		authorized.PUT("/buildings/:id", middleware.RequireRole(models.AdminRoleID, models.SuperAdminRoleID), controllers.UpdateBuilding)
		authorized.DELETE("/buildings/:id", middleware.RequireRole(models.SuperAdminRoleID), controllers.DeleteBuilding)

		// Floors CRUD routes
		authorized.POST("/floors", middleware.RequireRole(models.AdminRoleID, models.SuperAdminRoleID), controllers.CreateFloor)
		authorized.GET("/floors/:id", middleware.RequireRole(models.AdminRoleID, models.NormalUserRoleID), controllers.GetFloor)
		authorized.GET("/floors", middleware.RequireRole(models.SuperAdminRoleID), controllers.GetFloorsByCompany)
		authorized.GET("/floors/:floor_id/rooms", middleware.RequireRole(models.AdminRoleID, models.NormalUserRoleID, models.SuperAdminRoleID), controllers.GetRoomsByFloorID)
		authorized.PUT("/floors/:id", middleware.RequireRole(models.AdminRoleID, models.SuperAdminRoleID), controllers.UpdateFloor)
		authorized.DELETE("/floors/:id", middleware.RequireRole(models.SuperAdminRoleID), controllers.DeleteFloor)

		// Rooms CRUD routes
		authorized.POST("/rooms", middleware.RequireRole(models.AdminRoleID, models.SuperAdminRoleID), controllers.CreateRoom)
		authorized.GET("/rooms/:id", middleware.RequireRole(models.AdminRoleID, models.NormalUserRoleID), controllers.GetRoom)
		authorized.GET("/rooms", middleware.RequireRole(models.SuperAdminRoleID), controllers.GetRoomsByCompany)
		authorized.PUT("/rooms/:id", middleware.RequireRole(models.AdminRoleID, models.SuperAdminRoleID), controllers.UpdateRoom)
		authorized.DELETE("/rooms/:id", middleware.RequireRole(models.SuperAdminRoleID), controllers.DeleteRoom)
	}

	return r
}

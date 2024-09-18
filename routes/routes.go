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
		authorized.GET("/companies/:company_id/buildings", middleware.RequireRole(models.AdminRoleID, models.NormalUserRoleID, models.SuperAdminRoleID), controllers.GetBuildingsByCompany)
		authorized.PUT("/companies/:id", middleware.RequireRole(models.SuperAdminRoleID, models.AdminRoleID), controllers.UpdateCompany)
		authorized.DELETE("/companies/:id", middleware.RequireRole(models.SuperAdminRoleID), controllers.DeleteCompany)

		// Building CRUD routes
		authorized.POST("/buildings", middleware.RequireRole(models.AdminRoleID, models.SuperAdminRoleID), controllers.CreateBuilding)
		authorized.GET("/buildings/:id", middleware.RequireRole(models.AdminRoleID, models.NormalUserRoleID), controllers.GetBuilding)
		authorized.GET("/buildings", middleware.RequireRole(models.SuperAdminRoleID), controllers.GetAllBuildings)
		authorized.GET("/buildings/:building_id/floors", middleware.RequireRole(models.AdminRoleID, models.NormalUserRoleID, models.SuperAdminRoleID), controllers.GetFloorsByBuildingID)
		authorized.PUT("/buildings/:id", middleware.RequireRole(models.AdminRoleID, models.SuperAdminRoleID), controllers.UpdateBuilding)
		authorized.DELETE("/buildings/:id", middleware.RequireRole(models.SuperAdminRoleID), controllers.DeleteBuilding)

		// Floors CRUD routes
		authorized.POST("/floors", middleware.RequireRole(models.AdminRoleID, models.SuperAdminRoleID), controllers.CreateFloor)
		authorized.GET("/floors/:id", middleware.RequireRole(models.AdminRoleID, models.NormalUserRoleID), controllers.GetFloor)
		authorized.GET("/floors", middleware.RequireRole(models.SuperAdminRoleID), controllers.GetAllFloors)
		authorized.GET("/floors/:floor_id/rooms", middleware.RequireRole(models.AdminRoleID, models.NormalUserRoleID, models.SuperAdminRoleID), controllers.GetRoomsByFloorID)
		authorized.PUT("/floors/:id", middleware.RequireRole(models.AdminRoleID, models.SuperAdminRoleID), controllers.UpdateFloor)
		authorized.DELETE("/floors/:id", middleware.RequireRole(models.SuperAdminRoleID), controllers.DeleteFloor)

		// Rooms CRUD routes
		authorized.POST("/rooms", middleware.RequireRole(models.AdminRoleID, models.SuperAdminRoleID), controllers.CreateRoom)
		authorized.GET("/rooms/:id", middleware.RequireRole(models.AdminRoleID, models.NormalUserRoleID), controllers.GetRoom)
		authorized.GET("/rooms", middleware.RequireRole(models.SuperAdminRoleID), controllers.GetAllRooms)
		authorized.GET("/rooms/:room_id/locks", middleware.RequireRole(models.AdminRoleID, models.NormalUserRoleID, models.SuperAdminRoleID), controllers.GetRoomsByFloorID)
		authorized.PUT("/rooms/:id", middleware.RequireRole(models.AdminRoleID, models.SuperAdminRoleID), controllers.UpdateRoom)
		authorized.DELETE("/rooms/:id", middleware.RequireRole(models.SuperAdminRoleID), controllers.DeleteRoom)

		// Locks CRUD routes
		authorized.POST("/locks", middleware.RequireRole(models.AdminRoleID, models.SuperAdminRoleID), controllers.CreateLock)
		authorized.GET("/locks/:id", middleware.RequireRole(models.AdminRoleID, models.NormalUserRoleID), controllers.GetLock)
		authorized.GET("/locks", middleware.RequireRole(models.SuperAdminRoleID), controllers.GetAllLocks)
		authorized.GET("/locks/:lock_id/key_copies", middleware.RequireRole(models.AdminRoleID, models.NormalUserRoleID, models.SuperAdminRoleID), controllers.GetKeyCopiesByLockID)
		authorized.PUT("/locks/:id", middleware.RequireRole(models.AdminRoleID, models.SuperAdminRoleID), controllers.UpdateLock)
		authorized.DELETE("/locks/:id", middleware.RequireRole(models.SuperAdminRoleID), controllers.DeleteLock)

		// Key-Copy CRUD routes
		authorized.POST("/key_copies", middleware.RequireRole(models.AdminRoleID, models.SuperAdminRoleID), controllers.CreateKeyCopy)
		authorized.GET("/key_copies/:id", middleware.RequireRole(models.AdminRoleID, models.NormalUserRoleID), controllers.GetKeyCopy)
		authorized.GET("/key_copies", middleware.RequireRole(models.SuperAdminRoleID), controllers.GetAllKeyCopies)
		authorized.PUT("/key_copies/:id", middleware.RequireRole(models.AdminRoleID, models.SuperAdminRoleID), controllers.UpdateKeyCopy)
		authorized.DELETE("/key_copies/:id", middleware.RequireRole(models.SuperAdminRoleID), controllers.DeleteKeyCopy)
	}

	return r
}

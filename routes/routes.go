package routes

import (
	"st-portier-be/controllers"
	"st-portier-be/middleware"
	"st-portier-be/models"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {

	// Authentication
	r.POST("/login", controllers.Login)

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
		authorized.GET("/buildings", middleware.RequireRole(models.SuperAdminRoleID), controllers.GetAllBuildings)
		authorized.GET("/buildings/company/:company_id", middleware.RequireRole(models.AdminRoleID, models.NormalUserRoleID, models.SuperAdminRoleID), controllers.GetBuildingsByCompany)
		authorized.PUT("/buildings/:id", middleware.RequireRole(models.AdminRoleID, models.SuperAdminRoleID), controllers.UpdateBuilding)
		authorized.DELETE("/buildings/:id", middleware.RequireRole(models.SuperAdminRoleID), controllers.DeleteBuilding)

		// Floors CRUD routes
		authorized.POST("/floors", middleware.RequireRole(models.AdminRoleID, models.SuperAdminRoleID), controllers.CreateFloor)
		authorized.GET("/floors/:id", middleware.RequireRole(models.AdminRoleID, models.NormalUserRoleID), controllers.GetFloor)
		authorized.GET("/floors", middleware.RequireRole(models.SuperAdminRoleID), controllers.GetAllFloors)
		authorized.GET("/floors/building/:building_id", middleware.RequireRole(models.AdminRoleID, models.NormalUserRoleID, models.SuperAdminRoleID), controllers.GetFloorsByBuildingID)
		authorized.PUT("/floors/:id", middleware.RequireRole(models.AdminRoleID, models.SuperAdminRoleID), controllers.UpdateFloor)
		authorized.DELETE("/floors/:id", middleware.RequireRole(models.SuperAdminRoleID), controllers.DeleteFloor)

		// Rooms CRUD routes
		authorized.POST("/rooms", middleware.RequireRole(models.AdminRoleID, models.SuperAdminRoleID), controllers.CreateRoom)
		authorized.GET("/rooms/:id", middleware.RequireRole(models.AdminRoleID, models.NormalUserRoleID), controllers.GetRoom)
		authorized.GET("/rooms", middleware.RequireRole(models.SuperAdminRoleID), controllers.GetAllRooms)
		authorized.GET("/rooms/floor/:floor_id", middleware.RequireRole(models.AdminRoleID, models.NormalUserRoleID, models.SuperAdminRoleID), controllers.GetRoomsByFloorID)
		authorized.PUT("/rooms/:id", middleware.RequireRole(models.AdminRoleID, models.SuperAdminRoleID), controllers.UpdateRoom)
		authorized.DELETE("/rooms/:id", middleware.RequireRole(models.SuperAdminRoleID), controllers.DeleteRoom)

		// Locks CRUD routes
		authorized.POST("/locks", middleware.RequireRole(models.AdminRoleID, models.SuperAdminRoleID), controllers.CreateLock)
		authorized.GET("/locks/:id", middleware.RequireRole(models.AdminRoleID, models.NormalUserRoleID), controllers.GetLock)
		authorized.GET("/locks", middleware.RequireRole(models.SuperAdminRoleID), controllers.GetAllLocks)
		authorized.GET("/locks/room/:room_id", middleware.RequireRole(models.AdminRoleID, models.NormalUserRoleID, models.SuperAdminRoleID), controllers.GetRoomsByFloorID)
		authorized.PUT("/locks/:id", middleware.RequireRole(models.AdminRoleID, models.SuperAdminRoleID), controllers.UpdateLock)
		authorized.DELETE("/locks/:id", middleware.RequireRole(models.SuperAdminRoleID), controllers.DeleteLock)

		// Key-Copy CRUD routes
		authorized.POST("/key_copies", middleware.RequireRole(models.AdminRoleID, models.SuperAdminRoleID), controllers.CreateKeyCopy)
		authorized.GET("/key_copies/:id", middleware.RequireRole(models.AdminRoleID, models.NormalUserRoleID), controllers.GetKeyCopy)
		authorized.GET("/key_copies", middleware.RequireRole(models.SuperAdminRoleID), controllers.GetAllKeyCopies)
		authorized.GET("/key_copies/lock/:lock_id", middleware.RequireRole(models.AdminRoleID, models.NormalUserRoleID, models.SuperAdminRoleID), controllers.GetKeyCopiesByLockID)
		authorized.PUT("/key_copies/:id", middleware.RequireRole(models.AdminRoleID, models.SuperAdminRoleID), controllers.UpdateKeyCopy)
		authorized.DELETE("/key_copies/:id", middleware.RequireRole(models.SuperAdminRoleID), controllers.DeleteKeyCopy)

		// Employee CRUD routes #TODO: Get Employee by Company ID
		authorized.POST("/employees", middleware.RequireRole(models.AdminRoleID, models.SuperAdminRoleID), controllers.CreateEmployee)
		authorized.GET("/employees", middleware.RequireRole(models.SuperAdminRoleID, models.AdminRoleID), controllers.GetAllEmployees)
		authorized.GET("/employees/:id", middleware.RequireRole(models.AdminRoleID, models.SuperAdminRoleID), controllers.GetEmployeeByID)
		authorized.PUT("/employees/:id", middleware.RequireRole(models.AdminRoleID, models.SuperAdminRoleID), controllers.UpdateEmployee)
		authorized.DELETE("/employees/:id", middleware.RequireRole(models.AdminRoleID, models.SuperAdminRoleID), controllers.DeleteEmployee)

		// Key Copy Assignment routes #TODO: by Company ID
		authorized.POST("/employees/assign/:employee_id/:key_copy_id", middleware.RequireRole(models.NormalUserRoleID, models.AdminRoleID, models.SuperAdminRoleID), controllers.AssignKeyCopy)
		authorized.POST("/employees/revoke/:employee_id/:key_copy_id", middleware.RequireRole(models.NormalUserRoleID, models.AdminRoleID, models.SuperAdminRoleID), controllers.RevokeKeyCopy)
		authorized.GET("/employees/key-copies/:employee_id", middleware.RequireRole(models.NormalUserRoleID, models.AdminRoleID, models.SuperAdminRoleID), controllers.GetAssignedKeys)
	}
}

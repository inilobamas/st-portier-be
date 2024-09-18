package routes

import (
	"st-portier-be/controllers"
	"st-portier-be/middleware"

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
		// Users
		authorized.GET("/users", controllers.GetUsers)
		authorized.GET("/users/:id", controllers.GetUser)
		authorized.POST("/users", controllers.CreateUser)
		authorized.PUT("/users/:id", controllers.UpdateUser)
		authorized.DELETE("/users/:id", controllers.DeleteUser)

		// Companies
		authorized.GET("/companies", controllers.GetCompanies)
		authorized.GET("/companies/:id", controllers.GetCompany)
		authorized.POST("/companies", controllers.CreateCompany)
		authorized.PUT("/companies/:id", controllers.UpdateCompany)
		authorized.DELETE("/companies/:id", controllers.DeleteCompany)
	}

	return r
}

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/sharvatic/BookMyHotel/middleware"
	"github.com/sharvatic/BookMyHotel/controllers"
	"time"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"}, // Your frontend origin
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Allow Authorization header
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))
	r.Use(func(c *gin.Context) {
		c.Header("Cross-Origin-Opener-Policy", "same-origin")
		c.Header("Cross-Origin-Embedder-Policy", "require-corp")
		c.Next()
	})
	
	api := r.Group("/api")
	{
		api.POST("/signup", controllers.Signup)
		api.POST("/login", controllers.Login)
		api.POST("/setClaims", middleware.SetClaims)

		api.GET("/menu", controllers.ViewAllMenus)
		api.POST("/menu", middleware.AuthMiddleware("staff"), controllers.CreateMenu)
		api.GET("/menu/:menu_id", controllers.ViewAllMenuItems)
		api.POST("/menu/add", middleware.AuthMiddleware("staff"), controllers.AddMenuItem)
		api.PUT("/menu/:menu_id/:menu_item_id/update", middleware.AuthMiddleware("staff"), controllers.UpdateMenuItems)
		
		//Order routes
		//api.GET("/orders", middleware.AuthMiddleware("staff"), controllers.ViewAllOrders)
		api.POST("/orders", middleware.AuthMiddleware("user"), controllers.PlaceOrder)
		api.GET("/orders", middleware.AuthMiddleware("user"), controllers.ViewMyOrders)

		//Table routes
		api.POST("/tables/create", middleware.AuthMiddleware("staff"), controllers.CreateTable)
		api.POST("/tables/:id/book", middleware.AuthMiddleware("user"), controllers.BookTable)
		api.POST("/tables/:id/cancel", middleware.AuthMiddleware("user"), controllers.CancelTable)
		api.GET("/tables", controllers.GetTables)
	}
	return r
}

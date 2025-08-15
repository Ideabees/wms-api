package routes

import (
	"wms-app/internal/controllers"
	"wms-app/internal/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/ulule/limiter/v3"
	ginlimiter "github.com/ulule/limiter/v3/drivers/middleware/gin"
	memory "github.com/ulule/limiter/v3/drivers/store/memory"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// Rate limiter: 5 requests per second per IP
	rate, _ := limiter.NewRateFromFormatted("5-M")
	store := memory.NewStore()
	instance := limiter.New(store, rate)
	r.Use(ginlimiter.NewMiddleware(instance))

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"} // Replace with your client's origin(s)
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Authorization", "Content-Type"}
	config.AllowCredentials = true
	r.Use(cors.New(config))

	r.POST("/api/register", controllers.Register)

	r.POST("/api/login", controllers.Login)

	// Protected group
	protected := r.Group("/v1/api")
	protected.Use(utils.JWTMiddleware())
	{
		protected.GET("/permissions", controllers.GetPermissions)
		protected.POST("/customer", controllers.CreateCustomer)
		protected.GET("/customers", controllers.GetCustomers)
		protected.DELETE("/customers", controllers.DeleteCustomers)
		protected.POST("/create_bulk_customers", controllers.CreateBulkCustomers)
		protected.POST("/logout", controllers.Logout)
	}

	chater := r.Group("/v1/api/chat")
	chater.Use(utils.JWTMiddleware())
	{
		chater.POST("/create/one-to-one", controllers.CreateChat)
		chater.GET("/one-to-one/:chat_id", controllers.GetChat)
		chater.GET("/users/:user_id/chats", controllers.GetUserChats)
		chater.POST("/one-to-one/messages", controllers.SendMessage)
		chater.GET("/:chat_id/messages", controllers.GetMessages)
		chater.PUT("/messages/:message_id/read", controllers.MarkMessageRead)
		chater.POST("chats/send/message", controllers.SendMessageOneToOne)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

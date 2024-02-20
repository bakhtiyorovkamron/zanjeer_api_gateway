package api

import (
	_ "github.com/Projects/zanjeer_api_gateway/api/docs"
	"github.com/Projects/zanjeer_api_gateway/api/handlers"
	"github.com/Projects/zanjeer_api_gateway/config"
	"github.com/Projects/zanjeer_api_gateway/pkg/logger"
	"github.com/Projects/zanjeer_api_gateway/storage"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func New(cfg config.Config, strg storage.StorageI, log *logger.Logger) *gin.Engine {
	r := gin.Default()

	h := handlers.NewHandlerV1(&handlers.HandlerV1Config{
		Logger:   nil,
		Cfg:      cfg,
		Postgres: strg,
	})

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowBrowserExtensions = true
	corsConfig.AllowMethods = []string{"*"}
	r.Use(cors.New(corsConfig))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	admin := r.Group("/admin")
	admin.POST("/login", h.Login)
	// admin.Use(h.JwtAuthMiddleware())
	// admin.Use(h.SuperAdminCheckType())
	admin.POST("/add/admin", h.CreateAdmin)
	admin.GET("/get/admins", h.GetAdmins)

	user := r.Group("/user")
	user.POST("/register", h.UserRegister)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("swagger/doc.json")))
	return r
}

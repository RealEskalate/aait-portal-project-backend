package swagger

import (
	"github.com/Elizabethyonas/A2SV-Portal-Project/cmd/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupSwaggerRoutes(router *gin.Engine) {
	// Use the generated docs
	docs.SwaggerInfo.Title = "A2SV Portal API"
	docs.SwaggerInfo.Description = "API Server for A2SV Portal Application"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// Set up the Swagger UI endpoint
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

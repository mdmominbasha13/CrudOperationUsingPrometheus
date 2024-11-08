package router

import (
	"github.com/CrudOperationUsingPrometheus/pkg/authentication"
	"github.com/CrudOperationUsingPrometheus/pkg/config"
	"github.com/CrudOperationUsingPrometheus/pkg/controller"
	"github.com/CrudOperationUsingPrometheus/pkg/metrics"
	"github.com/CrudOperationUsingPrometheus/pkg/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/gin-gonic/gin"
)

func Init() {
	router := NewRouter()
	router.Run(config.Appconfig.GetString("server.port"))
}

func NewRouter() *gin.Engine {
	router := gin.New()
	metrics.InitMetric()
	resource := router.Group("/api")
	router.POST("/login", controller.Login)
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	resource.Use(middleware.LogRequestInfo(), authentication.Auth(), middleware.PrometheusMiddleware())
	{
		resource.GET("/book/", controller.GetBook)
		resource.GET("/book/:bookid", controller.GetBookById)
		resource.POST("/book/", controller.CreateBook)
		resource.PUT("/book/:bookid", controller.UpdateBookById)
		resource.DELETE("/book/:bookid", controller.DeleteBookById)

	}
	return router
}

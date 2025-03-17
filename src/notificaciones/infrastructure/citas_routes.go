// citas_routes.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type CitaRouter struct {
	engine *gin.Engine
}

func NewCitaRouter(engine *gin.Engine) *CitaRouter {
	return &CitaRouter{
		engine: engine,
	}
}

func (router *CitaRouter) Run() {
	// Inicializar dependencias
	createController, getByIdController, updateController, deleteController, getAllController := InitCitaDependencies()

	// Grupo de rutas para citas
	citaGroup := router.engine.Group("/notificacion")
	{
		citaGroup.POST("/", createController.Run)
		citaGroup.GET("/:id", getByIdController.Run)
		citaGroup.PUT("/:id", updateController.Run)
		citaGroup.DELETE("/:id", deleteController.Run)
		citaGroup.GET("/", getAllController.Run)
	}
}

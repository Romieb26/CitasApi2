// GetAllCitasController.go
package infrastructure

import (
	"net/http"

	application "notificaciones/src/notificaciones/application"

	"github.com/gin-gonic/gin"
)

type GetAllCitasController struct {
	getAllUseCase *application.GetAllCitasUseCase
}

func NewGetAllCitasController(getAllUseCase *application.GetAllCitasUseCase) *GetAllCitasController {
	return &GetAllCitasController{
		getAllUseCase: getAllUseCase,
	}
}

func (ctrl *GetAllCitasController) Run(c *gin.Context) {
	citas, err := ctrl.getAllUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener las citas",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, citas)
}

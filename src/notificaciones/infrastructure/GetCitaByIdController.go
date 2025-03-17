// GetCitaByIdController.go
package infrastructure

import (
	"net/http"
	"strconv"

	application "notificaciones/src/notificaciones/application"

	"github.com/gin-gonic/gin"
)

type GetCitaByIdController struct {
	getByIdUseCase *application.GetCitaByIdUseCase
}

func NewGetCitaByIdController(getByIdUseCase *application.GetCitaByIdUseCase) *GetCitaByIdController {
	return &GetCitaByIdController{
		getByIdUseCase: getByIdUseCase,
	}
}

func (ctrl *GetCitaByIdController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	cita, err := ctrl.getByIdUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener la cita",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, cita)
}

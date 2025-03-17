// DeleteCitaController.go
package infrastructure

import (
	"net/http"
	"strconv"

	application "notificaciones/src/notificaciones/application"

	"github.com/gin-gonic/gin"
)

type DeleteCitaController struct {
	deleteUseCase *application.DeleteCitaUseCase
}

func NewDeleteCitaController(deleteUseCase *application.DeleteCitaUseCase) *DeleteCitaController {
	return &DeleteCitaController{
		deleteUseCase: deleteUseCase,
	}
}

func (ctrl *DeleteCitaController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	errDelete := ctrl.deleteUseCase.Run(int32(id))
	if errDelete != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo eliminar la cita",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Cita eliminada exitosamente",
	})
}

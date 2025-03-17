// UpdateCitaController.go
package infrastructure

import (
	"log"
	"net/http"
	"strconv"

	"notificaciones/src/notificaciones/application"
	"notificaciones/src/notificaciones/domain/entities"

	"github.com/gin-gonic/gin"
)

type UpdateCitaController struct {
	updateUseCase *application.UpdateCitaUseCase
}

func NewUpdateCitaController(updateUseCase *application.UpdateCitaUseCase) *UpdateCitaController {
	return &UpdateCitaController{
		updateUseCase: updateUseCase,
	}
}

func (ctrl *UpdateCitaController) Run(c *gin.Context) {
	// Obtener el ID de la cita desde los parámetros de la URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	// Definir la estructura de la solicitud JSON
	var citaRequest struct {
		Estado string `json:"estado"`
	}

	// Parsear el JSON de la solicitud
	if err := c.ShouldBindJSON(&citaRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	// Crear una nueva entidad Cita con el ID y el estado proporcionado
	cita := entities.NewCita("", "", "", "", "", "", citaRequest.Estado)
	cita.CitaID = int32(id)

	log.Printf("Actualizando cita con ID %d a estado %s", cita.CitaID, cita.Estado)

	// Llamar al caso de uso para actualizar la cita
	err = ctrl.updateUseCase.Run(cita)
	if err != nil {
		log.Printf("Error al actualizar la cita: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo actualizar la cita",
			"error":   err.Error(),
		})
		return
	}

	// Devolver la cita actualizada
	c.JSON(http.StatusOK, gin.H{
		"message": "Estado de la cita actualizado correctamente",
	})
}

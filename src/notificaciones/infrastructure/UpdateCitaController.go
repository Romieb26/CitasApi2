// UpdateCitaController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"notificaciones/src/notificaciones/application"
	"notificaciones/src/notificaciones/domain/entities" // Importar el paquete entities

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
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	var citaRequest struct {
		NombrePaciente   string `json:"nombrePaciente"`
		ApellidoPaciente string `json:"apellidoPaciente"`
		NumeroContacto   string `json:"numeroContacto"`
		AreaCita         string `json:"areaCita"`
		Fecha            string `json:"fecha"`
		Hora             string `json:"hora"`
	}

	if err := c.ShouldBindJSON(&citaRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	cita := entities.NewCita( // Usar entities.NewCita
		citaRequest.NombrePaciente,
		citaRequest.ApellidoPaciente,
		citaRequest.NumeroContacto,
		citaRequest.AreaCita,
		citaRequest.Fecha,
		citaRequest.Hora,
	)
	cita.CitaID = int32(id)

	updatedCita, err := ctrl.updateUseCase.Run(cita)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo actualizar la cita",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedCita)
}

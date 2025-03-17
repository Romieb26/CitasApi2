// CreateCitaController.go
package infrastructure

import (
	"net/http"

	"notificaciones/src/notificaciones/application"
	"notificaciones/src/notificaciones/domain/entities" // Importar el paquete entities

	"github.com/gin-gonic/gin"
)

type CreateCitaController struct {
	createUseCase *application.CreateCitaUseCase
}

func NewCreateCitaController(createUseCase *application.CreateCitaUseCase) *CreateCitaController {
	return &CreateCitaController{
		createUseCase: createUseCase,
	}
}

func (ctrl *CreateCitaController) Run(c *gin.Context) {
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

	createdCita, err := ctrl.createUseCase.Run(cita)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo crear la cita",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdCita)
}

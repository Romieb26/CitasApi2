package infrastructure

import (
	"notificaciones/src/notificaciones/application"
)

func InitCitaDependencies() (
	*CreateCitaController,
	*GetCitaByIdController,
	*UpdateCitaController,
	*DeleteCitaController,
	*GetAllCitasController,
) {
	repo := NewMySQLCitaRepository()

	createUseCase := application.NewCreateCitaUseCase(repo)
	getByIdUseCase := application.NewGetCitaByIdUseCase(repo)
	updateUseCase := application.NewUpdateCitaUseCase(repo)
	deleteUseCase := application.NewDeleteCitaUseCase(repo)
	getAllUseCase := application.NewGetAllCitasUseCase(repo)

	createController := NewCreateCitaController(createUseCase)
	getByIdController := NewGetCitaByIdController(getByIdUseCase)
	updateController := NewUpdateCitaController(updateUseCase)
	deleteController := NewDeleteCitaController(deleteUseCase)
	getAllController := NewGetAllCitasController(getAllUseCase)

	return createController, getByIdController, updateController, deleteController, getAllController
}

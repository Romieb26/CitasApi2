// GetAllCitasUseCase.go
package application

import (
	repositories "notificaciones/src/notificaciones/domain"
	"notificaciones/src/notificaciones/domain/entities"
)

type GetAllCitasUseCase struct {
	repo repositories.ICita
}

func NewGetAllCitasUseCase(repo repositories.ICita) *GetAllCitasUseCase {
	return &GetAllCitasUseCase{repo: repo}
}

func (uc *GetAllCitasUseCase) Run() ([]entities.Cita, error) {
	citas, err := uc.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return citas, nil
}

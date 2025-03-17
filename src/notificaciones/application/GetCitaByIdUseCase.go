// GetCitaByIdUseCase.go
package application

import (
	repositories "notificaciones/src/notificaciones/domain"
	"notificaciones/src/notificaciones/domain/entities"
)

type GetCitaByIdUseCase struct {
	repo repositories.ICita
}

func NewGetCitaByIdUseCase(repo repositories.ICita) *GetCitaByIdUseCase {
	return &GetCitaByIdUseCase{repo: repo}
}

func (uc *GetCitaByIdUseCase) Run(id int32) (*entities.Cita, error) {
	cita, err := uc.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return cita, nil
}

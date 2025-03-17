// UpdateCitaUseCase.go
package application

import (
	repositories "notificaciones/src/notificaciones/domain"
	"notificaciones/src/notificaciones/domain/entities"
)

type UpdateCitaUseCase struct {
	repo repositories.ICita
}

func NewUpdateCitaUseCase(repo repositories.ICita) *UpdateCitaUseCase {
	return &UpdateCitaUseCase{repo: repo}
}

func (uc *UpdateCitaUseCase) Run(cita *entities.Cita) (*entities.Cita, error) {
	err := uc.repo.Update(cita)
	if err != nil {
		return nil, err
	}
	return cita, nil
}

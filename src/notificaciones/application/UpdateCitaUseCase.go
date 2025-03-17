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

func (uc *UpdateCitaUseCase) Run(cita *entities.Cita) error {
	return uc.repo.Update(cita)
}

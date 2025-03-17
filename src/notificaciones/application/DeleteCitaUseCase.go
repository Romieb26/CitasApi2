// DeleteCitaUseCase.go
package application

import repositories "notificaciones/src/notificaciones/domain"

type DeleteCitaUseCase struct {
	repo repositories.ICita
}

func NewDeleteCitaUseCase(repo repositories.ICita) *DeleteCitaUseCase {
	return &DeleteCitaUseCase{repo: repo}
}

func (uc *DeleteCitaUseCase) Run(id int32) error {
	err := uc.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

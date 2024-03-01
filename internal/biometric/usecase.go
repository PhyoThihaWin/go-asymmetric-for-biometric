package biometric

import (
	"pthw.com/asymmetric-for-biometric/models"
)

type UseCase interface {
	CreateBiometric(data *models.UserBiometric) (*models.UserBiometric, error)
}

type BiometricUseCase struct {
	biometricRepo Repository
}

func NewBiometricUseCase(biometricRepo Repository) *BiometricUseCase {
	return &BiometricUseCase{
		biometricRepo: biometricRepo,
	}
}

func (b BiometricUseCase) CreateBiometric(data *models.UserBiometric) (*models.UserBiometric, error) {
	return b.biometricRepo.CreateBiometric(data)
}

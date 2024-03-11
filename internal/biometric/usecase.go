package biometric

import (
	"pthw.com/asymmetric-for-biometric/models"
)

type UseCase interface {
	CreateBiometric(data *models.UserBiometric) (*models.UserBiometric, error)
	GetChallenge(biometricId string) (*models.CHALLENGE, error)
	ValidateBiometric(biometricId string, sigature string) (string, error)
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

func (b BiometricUseCase) GetChallenge(biometricId string) (*models.CHALLENGE, error) {
	return b.biometricRepo.GetChallenge(biometricId)
}

func (b BiometricUseCase) ValidateBiometric(biometricId string, sigature string) (string, error) {
	return b.biometricRepo.ValidateBiometric(biometricId, sigature)
}

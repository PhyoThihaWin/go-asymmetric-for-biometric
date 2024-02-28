package biometric

import (
	"context"
)

type BiometricUseCase struct {
	biometricRepo Repository
}

func NewBiometricUseCase(biometricRepo Repository) *BiometricUseCase {
	return &BiometricUseCase{
		biometricRepo: biometricRepo,
	}
}

func (b BiometricUseCase) CreateBiometric(ctx context.Context, id string) error {
	return nil
}

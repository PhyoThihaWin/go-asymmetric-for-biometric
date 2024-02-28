package biometric

import (
	"context"

	"gorm.io/gorm"
)

type BiometricUseCase struct {
	db *gorm.DB
}

func NewBiometricUseCase(d *gorm.DB) *BiometricUseCase {
	return &BiometricUseCase{
		db: d,
	}
}

func (b BiometricUseCase) CreateBiometric(ctx context.Context, id string) error {
	return nil
}

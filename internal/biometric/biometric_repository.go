package biometric

import (
	"context"

	"gorm.io/gorm"
)

type BiometricRepository struct {
	db *gorm.DB
}

func NewBiometricRepository(db *gorm.DB) *BiometricRepository {
	return &BiometricRepository{
		db: db,
	}
}

func (b BiometricRepository) CreateBiometric(ctx context.Context, id string) error {
	return nil
}

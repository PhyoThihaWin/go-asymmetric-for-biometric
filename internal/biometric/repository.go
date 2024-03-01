package biometric

import (
	"gorm.io/gorm"
	"pthw.com/asymmetric-for-biometric/models"
)

type Repository interface {
	CreateBiometric(data *models.UserBiometric) (*models.UserBiometric, error)
	// CreateChallenge(data *models.CHALLENGE) (*models.CHALLENGE, error)
}

type BiometricRepository struct {
	db *gorm.DB
}

func NewBiometricRepository(db *gorm.DB) *BiometricRepository {
	return &BiometricRepository{
		db: db,
	}
}

func (b BiometricRepository) CreateBiometric(data *models.UserBiometric) (*models.UserBiometric, error) {
	if b.IsDeviceIdExist(data.DEVICE_ID) {
		result := b.db.Model(&data).Where("device_id = ?", data.DEVICE_ID).Updates(&data)
		return data, result.Error
	} else {
		result := b.db.Create(&data)

		var UserBiometric models.UserBiometric
		b.db.First(UserBiometric, result.RowsAffected)

		// b.CreateChallenge(&UserBiometric)
		return &UserBiometric, result.Error
	}
}

func (b BiometricRepository) CreateChallenge(data *models.CHALLENGE) {
	if !b.IsDeviceIdExistInChallenge(data.DEVICE_ID) {
		b.db.Create(&data)
	}
}

func (b BiometricRepository) IsDeviceIdExist(deviceId string) bool {
	var UserBiometric models.UserBiometric
	result := b.db.First(&UserBiometric, "device_id = ?", deviceId)

	if result.Error == gorm.ErrRecordNotFound {
		// ID exist in the database
		return false
	} else {
		// ID does not exist in the database
		return true
	}
}

func (b BiometricRepository) IsDeviceIdExistInChallenge(deviceId string) bool {
	var Challenge models.CHALLENGE
	result := b.db.First(&Challenge, "device_id = ?", deviceId)

	if result.Error == gorm.ErrRecordNotFound {
		// ID exist in the database
		return false
	} else {
		// ID does not exist in the database
		return true
	}
}

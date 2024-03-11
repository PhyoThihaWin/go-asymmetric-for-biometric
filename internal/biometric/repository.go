package biometric

import (
	"errors"
	"fmt"
	"strconv"

	"gorm.io/gorm"
	"pthw.com/asymmetric-for-biometric/models"
	"pthw.com/asymmetric-for-biometric/utils"
)

type Repository interface {
	CreateBiometric(data *models.UserBiometric) (*models.UserBiometric, error)
	GetChallenge(deviceId string) (*models.CHALLENGE, error)
	ValidateBiometric(biometricId string, sigature string) (string, error)
}

type BiometricRepository struct {
	db *gorm.DB
}

func NewBiometricRepository(db *gorm.DB) *BiometricRepository {
	return &BiometricRepository{
		db: db,
	}
}

// create biometric
func (b BiometricRepository) CreateBiometric(data *models.UserBiometric) (*models.UserBiometric, error) {
	var result *models.UserBiometric
	var err *error
	if b.IsDeviceIdExist(data.DEVICE_ID) {
		raw := b.db.Where("device_id = ?", data.DEVICE_ID).Updates(&data)

		var UserBiometric models.UserBiometric
		raw.First(&UserBiometric, "device_id = ?", data.DEVICE_ID)
		result = &UserBiometric
		err = &raw.Error
	} else {
		raw := b.db.Create(&data)

		b.db.First(&data, &data.BIOMETRIC_ID)
		fmt.Println("Result: " + strconv.FormatUint(uint64(data.ID), 10))

		result = data
		err = &raw.Error
	}

	Challenge := models.CHALLENGE{
		CHALLENGE:    utils.RandRunes(15),
		BIOMETRIC_ID: data.BIOMETRIC_ID,
		DEVICE_ID:    data.DEVICE_ID,
	}
	b.CreateChallenge(&Challenge)

	return result, *err
}

func (b BiometricRepository) CreateChallenge(data *models.CHALLENGE) {
	if b.IsDeviceIdExistInChallenge(data.DEVICE_ID) {
		b.db.Where("device_id = ?", data.DEVICE_ID).Updates(&data)
	} else {
		b.db.Create(&data)
	}
}

func (b BiometricRepository) IsDeviceIdExist(deviceId string) bool {
	var UserBiometric models.UserBiometric
	result := b.db.First(&UserBiometric, "device_id = ?", deviceId)

	if result.Error == gorm.ErrRecordNotFound {
		// ID does not exist in the database
		return false
	} else {
		// ID exist in the database
		return true
	}
}

func (b BiometricRepository) IsDeviceIdExistInChallenge(deviceId string) bool {
	var Challenge models.CHALLENGE
	result := b.db.First(&Challenge, "device_id = ?", deviceId)
	if result.Error == gorm.ErrRecordNotFound {
		// ID does not exist in the database
		return false
	} else {
		// ID exist in the database
		return true
	}
}

// get challenge
func (b BiometricRepository) GetChallenge(biometricId string) (*models.CHALLENGE, error) {
	var Challenge models.CHALLENGE
	result := b.db.First(&Challenge, "biometric_id =?", biometricId)
	return &Challenge, result.Error
}

// verify signature
func (b BiometricRepository) ValidateBiometric(biometricId string, sigature string) (string, error) {
	var UserBiometric models.UserBiometric
	var Challenge models.CHALLENGE
	result := b.db.First(&UserBiometric, "biometric_id =?", biometricId)
	result2 := b.db.First(&Challenge, "device_id =?", UserBiometric.DEVICE_ID)

	signatureValid := false
	if result.Error == nil && result2.Error == nil {
		fmt.Println("Result: " + Challenge.CHALLENGE)
		fmt.Println("Result: " + sigature)
		fmt.Println("Result: " + UserBiometric.PUBLIC_KEY)
		signatureValid = utils.ValidateSignature(
			Challenge.CHALLENGE, sigature, UserBiometric.PUBLIC_KEY,
		)
	} else {
		return "", errors.New("Signature validation failed")
	}

	if signatureValid {
		return "Signature validation success", nil
	} else {
		return "", errors.New("Signature validation failed")
	}

}

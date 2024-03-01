package models

import (
	"pthw.com/asymmetric-for-biometric/utils"
)

type UserBiometric struct {
	utils.CustomModel
	DEVICE_ID    string `json:"device_id"`
	PUBLIC_KEY   string `json:"public_key"`
	BIOMETRIC_ID string `json:"biometric_id"`
}

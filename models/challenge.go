package models

import "pthw.com/asymmetric-for-biometric/utils"

type CHALLENGE struct {
	utils.CustomModel
	CHALLENGE string `json:"challenge"`
	DEVICE_ID string `json:"device_id"`
}

package services

import (
	"wms-app/config"
	"wms-app/internal/models/dbModels"
)

func CreateUser(user *dbModels.User) error {
	return config.DB.Create(user).Error
}

package service

import "gorm.io/gorm"

type WalletDbService struct {
	db *gorm.DB
}

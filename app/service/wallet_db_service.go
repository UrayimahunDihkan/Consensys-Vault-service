package service

import "gorm.io/gorm"

type WalletDbService struct {
	db *gorm.DB
}

func NewWalletDbService(db *gorm.DB) *WalletDbService {
	return &WalletDbService{db: db}
}

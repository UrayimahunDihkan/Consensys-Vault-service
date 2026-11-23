package app

import (
	"consensus-wallet/app/abi"
	"fmt"

	"gorm.io/gorm"
)

var DB *gorm.DB

func Initialize() error {
	if err := LoadConfig(); err != nil {
		return fmt.Errorf("load config error: %w", err)
	}

	if err := abi.LoadABI(); err != nil {
		return fmt.Errorf("abi loading error: %w", err)
	}

	//dsn := "user:password@tcp(127.0.0.1:3306)/your_db?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:Root@123456@tcp(127.0.0.1:3306)/defi_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := InitDB(dsn)
	if err != nil {
		// %v 建议换成 %w
		// 用 %w 可以保留原始错误，上层能用 errors.Is/As 判断根因
		return fmt.Errorf("init db error: %w", err)
	}
	DB = db

	return err
}

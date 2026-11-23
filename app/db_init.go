package app

import (
	"consensus-wallet/app/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err // 连接失败，直接返回
	}
	if err := db.AutoMigrate(&model.PostInfo{}); err != nil { // 连接成功才迁移
		return nil, err
	}
	return db, nil
}

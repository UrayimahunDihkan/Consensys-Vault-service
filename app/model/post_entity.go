package model

type PostInfo struct {
	Id       int64  `gorm:"primaryKey;autoIncrement"`
	Hash     string `gorm:"type:varchar(64);index"`
	PostType int64  `gorm:"index"`
	Content  string `gorm:"type:varchar(1000)"`
}

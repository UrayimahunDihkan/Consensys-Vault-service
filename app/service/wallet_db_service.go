package service

import (
	"consensus-wallet/app/controller/rsp"
	"consensus-wallet/app/model"
	"context"

	"gorm.io/gorm"
)

type WalletDbService struct {
	db *gorm.DB
}

func NewWalletDbService(db *gorm.DB) *WalletDbService {
	return &WalletDbService{db: db}
}

func (s *WalletDbService) Posts(ctx context.Context, page, size int, postType int64) (*rsp.PostListResult, error) {
	// 参数兜底
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 20
	}
	if size > 100 {
		size = 100
	}

	query := s.db.WithContext(ctx).Model(&model.PostInfo{})
	if postType > 0 {
		query = query.Where("post_type = ?", postType)
	}

	// 总数
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	// 列表（按 id 倒序，最新在前）
	var posts []model.PostInfo
	offset := (page - 1) * size
	if err := query.Order("id DESC").Offset(offset).Limit(size).Find(&posts).Error; err != nil {
		return nil, err
	}

	return &PostListResult{
		Total: total,
		Page:  page,
		Size:  size,
		List:  posts,
	}, nil
}

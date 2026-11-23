package rsp

import "consensus-wallet/app/model"

type PostListResult struct {
	Total int64            `json:"total"`
	Page  int              `json:"page"`
	Size  int              `json:"size"`
	List  []model.PostInfo `json:"list"`
}

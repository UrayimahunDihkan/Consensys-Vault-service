package controller

import "consensus-wallet/app/service"

type WalletController struct {
	walletDbService *service.WalletDbService
}

func NewWalletController(_walletDbService *service.WalletDbService) *WalletController {
	return &WalletController{
		walletDbService: _walletDbService,
	}
}

//func (c *WalletController) Posts(c *gin.Context) {
//
//}

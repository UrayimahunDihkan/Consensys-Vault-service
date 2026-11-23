package controller

import (
	"consensus-wallet/app"
	"consensus-wallet/app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WalletController struct {
	walletDbService *service.WalletDbService
}

func NewWalletController() *WalletController {
	return &WalletController{
		walletDbService: service.NewWalletDbService(app.DB),
	}
}

// GET /posts?page=1&size=20&post_type=1
func (c *WalletController) Posts(g *gin.Context) {
	
	page, _ := strconv.Atoi(g.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(g.DefaultQuery("size", "20"))
	postType, _ := strconv.ParseInt(g.DefaultQuery("post_type", "0"), 10, 64)

	// 调 service
	result, err := c.walletDbService.Posts(g.Request.Context(), page, size, postType)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "查询帖子失败"})
		return
	}

	g.JSON(http.StatusOK, result)
}

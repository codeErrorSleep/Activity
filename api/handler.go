package api

import (
	"Activity/constant"
	"Activity/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	gameService GameService
}

func NewHandler(gameService GameService) *Handler {
	return &Handler{
		gameService: gameService,
	}
}

// RegisterRoutes 注册路由
func (h *Handler) RegisterRoutes(r *gin.Engine) {
	game := r.Group(constant.GameBasePath)
	{
		game.POST("/participate", h.ParticipateGame)
		game.GET("/status", h.GetGameStatus)
		game.GET("/prize", h.GetUserPrize)
	}
}

// ParticipateGame 参与玩法
func (h *Handler) ParticipateGame(c *gin.Context) {
	var req ParticipateGameReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, BaseResp{
			Code:    constant.ErrInvalidParam,
			Message: "invalid params",
		})
		return
	}

	// 获取用户信息（实际项目中应该从中间件获取）
	user := models.User{
		Uid: c.GetString("uid"),
	}

	// 根据玩法类型创建对应的动作
	var action models.ActionInterface
	switch req.GameType {
	case "post":
		action = &models.CommunityPostAction{
			PostID: req.PostID,
		}
	case "checkin":
		action = &models.CheckinAction{
			CheckinTime: time.Now(),
		}
	default:
		c.JSON(http.StatusBadRequest, BaseResp{
			Code:    constant.ErrInvalidParam,
			Message: "unsupported game type",
		})
		return
	}

	// 执行玩法逻辑
	result, err := h.gameService.ParticipateGame(c, user, req.ActivityID, req.GameName, action)
	if err != nil {
		c.JSON(http.StatusInternalServerError, BaseResp{
			Code:    constant.ErrSystem,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, BaseResp{
		Code:    0,
		Message: "success",
		Data:    result,
	})
}

// GetGameStatus 获取玩法状态
func (h *Handler) GetGameStatus(c *gin.Context) {
	var req GetGameStatusReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, BaseResp{
			Code:    constant.ErrInvalidParam,
			Message: "invalid params",
		})
		return
	}

	// 获取用户信息
	user := models.User{
		Uid: c.GetString("uid"),
	}

	// 获取状态
	status, err := h.gameService.GetGameStatus(c, user, req.ActivityID, req.GameName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, BaseResp{
			Code:    constant.ErrSystem,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, BaseResp{
		Code:    0,
		Message: "success",
		Data:    status,
	})
}

// GetUserPrize 获取用户奖品
func (h *Handler) GetUserPrize(c *gin.Context) {
	var req GetUserPrizeReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, BaseResp{
			Code:    constant.ErrInvalidParam,
			Message: "invalid params",
		})
		return
	}

	// 获取用户信息
	user := models.User{
		Uid: c.GetString("uid"),
	}

	// 获取奖品信息
	prize, err := h.gameService.GetUserPrize(c, user, req.ActivityID, req.GameName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, BaseResp{
			Code:    constant.ErrSystem,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, BaseResp{
		Code:    0,
		Message: "success",
		Data:    prize,
	})
}

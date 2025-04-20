package api

import (
	"Activity/constant"
	"Activity/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	gameService     GameService
	activityService ActivityService
}

func NewHandler(gameService GameService, activityService ActivityService) *Handler {
	return &Handler{
		gameService:     gameService,
		activityService: activityService,
	}
}

// RegisterRoutes 注册所有API路由
func (h *Handler) RegisterRoutes(r *gin.Engine) {
	// API版本分组
	v1 := r.Group("")
	{
		// 活动相关接口
		activity := v1.Group("/activity")
		{
			activity.POST("", h.CreateActivity)
			activity.PUT("/:id", h.UpdateActivity)
			activity.GET("/:id", h.GetActivity)
			activity.POST("/:id/participate", h.Participate)
			activity.GET("/:id/participation", h.GetParticipation)
		}

		// 游戏相关接口
		game := v1.Group("/game")
		{
			game.POST("/participate", h.ParticipateGame)
			game.GET("/status", h.GetGameStatus)
			game.GET("/prize", h.GetUserPrize)
		}
	}
}

// @Summary		创建活动
// @Description	创建一个新的活动
// @Tags			活动管理
// @Accept			json
// @Produce		json
// @Param			activity	body		CreateActivityRequest	true	"活动信息"
// @Success		200			{object}	BaseResp{data=CreateActivityResponse}
// @Failure		400			{object}	BaseResp
// @Failure		500			{object}	BaseResp
// @Router			/activity [post]
func (h *Handler) CreateActivity(c *gin.Context) {
	var req CreateActivityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, BaseResp{
			Code:    constant.ErrInvalidParam,
			Message: "invalid params",
		})
		return
	}

	resp, err := h.activityService.CreateActivity(c, &req)
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
		Data:    resp,
	})
}

// @Summary		更新活动
// @Description	更新指定ID的活动信息
// @Tags			活动管理
// @Accept			json
// @Produce		json
// @Param			id			path		string					true	"活动ID"
// @Param			activity	body		UpdateActivityRequest	true	"活动信息"
// @Success		200			{object}	BaseResp{data=UpdateActivityResponse}
// @Failure		400			{object}	BaseResp
// @Failure		500			{object}	BaseResp
// @Router			/activity/{id} [put]
func (h *Handler) UpdateActivity(c *gin.Context) {
	var req UpdateActivityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, BaseResp{
			Code:    constant.ErrInvalidParam,
			Message: "invalid params",
		})
		return
	}

	resp, err := h.activityService.UpdateActivity(c, &req)
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
		Data:    resp,
	})
}

// @Summary		获取活动信息
// @Description	获取指定ID的活动详细信息
// @Tags			活动管理
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"活动ID"
// @Success		200	{object}	BaseResp{data=GetActivityResponse}
// @Failure		400	{object}	BaseResp
// @Failure		500	{object}	BaseResp
// @Router			/activity/{id} [get]
func (h *Handler) GetActivity(c *gin.Context) {
	activityID := c.Param("id")
	if activityID == "" {
		c.JSON(http.StatusBadRequest, BaseResp{
			Code:    constant.ErrInvalidParam,
			Message: "activity_id is required",
		})
		return
	}

	// 将字符串ID转换为int64
	id, err := strconv.ParseInt(activityID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, BaseResp{
			Code:    constant.ErrInvalidParam,
			Message: "invalid activity_id",
		})
		return
	}

	resp, err := h.activityService.GetActivity(c, id)
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
		Data:    resp,
	})
}

// @Summary		参与活动
// @Description	用户参与指定活动
// @Tags			活动管理
// @Accept			json
// @Produce		json
// @Param			id				path		string				true	"活动ID"
// @Param			participation	body		ParticipateRequest	true	"参与信息"
// @Success		200				{object}	BaseResp{data=ParticipateResponse}
// @Failure		400				{object}	BaseResp
// @Failure		500				{object}	BaseResp
// @Router			/activity/{id}/participate [post]
func (h *Handler) Participate(c *gin.Context) {
	var req ParticipateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, BaseResp{
			Code:    constant.ErrInvalidParam,
			Message: "invalid params",
		})
		return
	}

	resp, err := h.activityService.Participate(c, &req)
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
		Data:    resp,
	})
}

// @Summary		获取参与记录
// @Description	获取用户在指定活动中的参与记录
// @Tags			活动管理
// @Accept			json
// @Produce		json
// @Param			id		path		string	true	"活动ID"
// @Param			user_id	query		string	true	"用户ID"
// @Success		200		{object}	BaseResp{data=GetParticipationResponse}
// @Failure		400		{object}	BaseResp
// @Failure		500		{object}	BaseResp
// @Router			/activity/{id}/participation [get]
func (h *Handler) GetParticipation(c *gin.Context) {
	activityID := c.Param("id")
	if activityID == "" {
		c.JSON(http.StatusBadRequest, BaseResp{
			Code:    constant.ErrInvalidParam,
			Message: "activity_id is required",
		})
		return
	}

	// 将字符串ID转换为int64
	id, err := strconv.ParseInt(activityID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, BaseResp{
			Code:    constant.ErrInvalidParam,
			Message: "invalid activity_id",
		})
		return
	}

	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, BaseResp{
			Code:    constant.ErrInvalidParam,
			Message: "user_id is required",
		})
		return
	}

	resp, err := h.activityService.GetParticipation(c, id, userID)
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
		Data:    resp,
	})
}

// @Summary		参与玩法
// @Description	用户参与指定玩法
// @Tags			玩法管理
// @Accept			json
// @Produce		json
// @Param			participation	body		ParticipateGameReq	true	"参与信息"
// @Success		200				{object}	BaseResp{data=ParticipateGameResponse}
// @Failure		400				{object}	BaseResp
// @Failure		500				{object}	BaseResp
// @Router			/game/participate [post]
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
	switch req.GameName {
	case "post":
		action = &models.CommunityPostAction{
			// PostID: req.PostID,
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

// @Summary		获取玩法状态
// @Description	获取用户在指定玩法中的状态
// @Tags			玩法管理
// @Accept			json
// @Produce		json
// @Param			activity_id	query		string	true	"活动ID"
// @Param			game_name	query		string	true	"玩法名称"
// @Success		200			{object}	BaseResp{data=GetGameStatusResponse}
// @Failure		400			{object}	BaseResp
// @Failure		500			{object}	BaseResp
// @Router			/game/status [get]
func (h *Handler) GetGameStatus(c *gin.Context) {
	var req GetGameStatusReq
	if err := c.ShouldBindQuery(&req); err != nil {
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

	resp, err := h.gameService.GetGameStatus(c, user, req.ActivityID, req.GameName)
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
		Data:    resp,
	})
}

// @Summary		获取用户奖品
// @Description	获取用户在指定玩法中获得的奖品
// @Tags			玩法管理
// @Accept			json
// @Produce		json
// @Param			activity_id	query		string	true	"活动ID"
// @Param			game_name	query		string	true	"玩法名称"
// @Success		200			{object}	BaseResp{data=GetUserPrizeResponse}
// @Failure		400			{object}	BaseResp
// @Failure		500			{object}	BaseResp
// @Router			/game/prize [get]
func (h *Handler) GetUserPrize(c *gin.Context) {
	var req GetUserPrizeReq
	if err := c.ShouldBindQuery(&req); err != nil {
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

	resp, err := h.gameService.GetUserPrize(c, user, req.ActivityID, req.GameName)
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
		Data:    resp,
	})
}

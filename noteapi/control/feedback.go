package control

import (
	"noteapi/model"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/zxysilent/utils"
)

// FeedbackAdd doc
// @Tags feedback
// @Summary 添加反馈
// @Param body body model.Feedback true "请求数据"
// @Success 200 {object} model.Reply "返回数据"
// @Router /api/feedback/add [post]
func FeedbackAdd(ctx echo.Context) error {
	ipt := &model.Feedback{}
	err := ctx.Bind(ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
	}
	ipt.Ctime = time.Now().Unix()
	err = model.FeedbackAdd(ipt)
	if err != nil {
		return ctx.JSON(utils.Fail("添加失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("succ"))
}

// FeedbackAll doc
// @Tags feedback
// @Summary 所有反馈
// @Success 200 {object} model.Reply{data=[]model.Feedback} "返回数据"
// @Router /api/feedback/all [get]
func FeedbackAll(ctx echo.Context) error {
	mods, err := model.FeedbackAll()
	if err != nil {
		return ctx.JSON(utils.ErrOpt("未查询到数据", err.Error()))
	}
	return ctx.JSON(utils.Succ("succ", mods))
}

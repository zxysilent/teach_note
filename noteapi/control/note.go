package control

import (
	"noteapi/model"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/zxysilent/utils"
)

// NoteGet doc
// @Tags note
// @Summary 通过id获取笔记
// @Param id query int true "id"
// @Success 200 {object} model.Reply{data=model.Note} "返回数据"
// @Router /api/note/get [get]
func NoteGet(ctx echo.Context) error {
	ipt := &model.IptId{}
	err := ctx.Bind(ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
	}
	mod, err := model.NoteGet(ipt.Id)
	if err != nil {
		return ctx.JSON(utils.ErrOpt("未查询到用户信息", err.Error()))
	}
	return ctx.JSON(utils.Succ("succ", mod))
}

// NoteAll doc
// @Tags note
// @Summary 所有笔记
// @Success 200 {object} model.Reply{data=[]model.Note} "返回数据"
// @Router /api/note/all [get]
func NoteAll(ctx echo.Context) error {
	mods, err := model.NoteAll()
	if err != nil {
		return ctx.JSON(utils.ErrOpt("未查询到数据", err.Error()))
	}
	return ctx.JSON(utils.Succ("succ", mods))
}

// NoteAdd doc
// @Tags note
// @Summary 添加笔记
// @Param token query string true "凭证"
// @Param body body model.Note true "请求数据"
// @Success 200 {object} model.Reply "返回数据"
// @Router /adm/note/add [post]
func NoteAdd(ctx echo.Context) error {
	ipt := &model.Note{}
	err := ctx.Bind(ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
	}
	ipt.Ctime = time.Now().Unix()
	ipt.Utime = ipt.Ctime
	err = model.NoteAdd(ipt)
	if err != nil {
		return ctx.JSON(utils.Fail("添加失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("succ"))
}

// NoteEdit doc
// @Tags note
// @Summary 修改笔记
// @Param token query string true "凭证"
// @Param body body model.Note true "请求数据"
// @Success 200 {object} model.Reply "返回数据"
// @Router /adm/note/edit [post]
func NoteEdit(ctx echo.Context) error {
	ipt := &model.Note{}
	err := ctx.Bind(ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
	}
	ipt.Utime = time.Now().Unix()
	err = model.NoteEdit(ipt)
	if err != nil {
		return ctx.JSON(utils.Fail("修改失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("succ"))
}

// NoteDrop doc
// @Tags note
// @Summary 删除笔记
// @Param body body model.IptId true "请求数据"
// @Param token query string true "凭证"
// @Success 200 {object} model.Reply "返回数据"
// @Router /adm/note/drop [post]
func NoteDrop(ctx echo.Context) error {
	ipt := &model.IptId{}
	err := ctx.Bind(ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
	}
	mod, err := model.NoteGet(ipt.Id)
	if err != nil {
		return ctx.JSON(utils.ErrOpt("删除失败", "不存在当前数据"))
	}
	err = model.NoteDrop(mod.Id)
	if err != nil {
		return ctx.JSON(utils.ErrOpt("删除失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("succ"))
}

package router

import (
	"noteapi/control"

	"github.com/labstack/echo/v4"
)

// apiRouter 通用访问
func apiRouter(api *echo.Group) {
	// note
	api.GET("/note/get", control.NoteGet)
	api.GET("/note/all", control.NoteAll)
	api.POST("/note/add", control.NoteAdd)
	api.POST("/note/edit", control.NoteEdit)
	api.POST("/note/drop", control.NoteDrop)

	// feedback
	api.GET("/feedback/all", control.FeedbackAll)
	api.POST("/feedback/add", control.FeedbackAdd)
}

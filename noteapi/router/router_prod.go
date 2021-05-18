// +build prod

package router

import (
	"github.com/labstack/echo/v4"
	"github.com/zxysilent/logs"
)

func init() {
	logs.SetLevel(logs.WARN)
}

// RegDocs 注册文档
// prod[正式] 模式不需要文档
func RegDocs(engine *echo.Echo) {
}

// midLogger 中间件-日志记录
func midLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return next
}

package router

import (
	"noteapi/conf"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// RunApp 入口
func RunApp() {
	engine := echo.New()
	// 恢复 日志记录
	engine.Use(midRecover, midLogger)
	// 跨域设置
	engine.Use(middleware.CORSWithConfig(crosConfig))
	// 不显示横幅
	engine.HideBanner = true
	// 自定义错误处理
	engine.HTTPErrorHandler = HTTPErrorHandler
	engine.Debug = conf.App.IsDev()
	// 静态目录
	engine.Static("/static", "static")
	// ico
	engine.File("/favicon.ico", "favicon.ico")
	// 注册文档
	RegDocs(engine)
	api := engine.Group("/api")
	apiRouter(api)
	engine.Start(conf.App.Addr)
}

/*  正式模式 编译 取消文档
 *  go build -tags=prod -o prod.exe main.go
 *
 *  开发模式 编译 添加文档
 *  go build -o dev.exe main.go
 */

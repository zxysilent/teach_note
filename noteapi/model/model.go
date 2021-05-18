package model

import (
	"noteapi/conf"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/zxysilent/logs"
)

// Db 数据库操作句柄
var Db *sqlx.DB

// Init 初始化数据库连接
func Init() {
	// 初始化数据库操作的 Xorm
	db, err := sqlx.Open("sqlite3", "./feedback.db")
	if err != nil {
		logs.Fatal("数据库 open:", err.Error())
	}
	if err = db.Ping(); err != nil {
		logs.Fatal("数据库 ping:", err.Error())
	}
	// 最大空闲链接
	db.SetMaxIdleConns(conf.App.DbIdle)
	// 最大打开链接
	db.SetMaxOpenConns(conf.App.DbOpen)
	Db = db
	logs.Info("model init")
}

type IptId struct {
	Id int `query:"id" form:"id" json:"id"` //仅包含Id的请求
}

// Reply 生成api文档使用
// 代码里未使用，也不要使用
type Reply struct {
	Code int    `json:"code" example:"200"`
	Msg  string `json:"msg" example:"提示信息"`
}

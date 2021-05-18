package model

import (
	"fmt"
	"log"
	"reflect"
	"testing"

	"github.com/jmoiron/sqlx"
)

func TestTag(t *testing.T) {
	var ptr interface{} = new(Note)
	ref := reflect.TypeOf(ptr)
	if ref.Kind() == reflect.Ptr {
		ref = ref.Elem()
	}
	if ref.Kind() != reflect.Struct {
		panic("Check type error not Struct")
	}
	// do something
	fieldNum := ref.NumField()
	for i := 0; i < fieldNum; i++ {
		tags := ref.Field(i).Tag
		json := tags.Get("json")
		if json != "" && tags.Get("db") != "-" {
			fmt.Printf("%s,", json)
		}
	}
}
func TestDb(t *testing.T) {
	// 初始化数据库操作的 Xorm
	db, err := sqlx.Open("sqlite3", "./feedback.db")
	if err != nil {
		log.Fatalln("数据库 open:", err.Error())
	}
	if err = db.Ping(); err != nil {
		log.Fatalln("数据库 ping:", err.Error())
	}
	// do something
	mod := Feedback{}
	err = db.Get(&mod, "SELECT * FROM f feedback id = ?", 1)
	log.Println(mod, err)
}

package db

import (
	"testing"
)

func TestInitEngine(t *testing.T) {
	err := initEngine()
	if err != nil {
		t.Error("数据库连接失败")
	}
}

func TestStdMasterDB(t *testing.T) {
	db := StdMasterDB()
	if db == nil {
		t.Error("数据库连接失败")
	}
}

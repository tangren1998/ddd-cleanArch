package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"github.com/tangren1998/ddd-cheanArch/user/domain"
)

var db *gorm.DB

func init() {
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	//db.AutoMigrate(&SignCode{})
	// 迁移 schema
	db.AutoMigrate(&domain.User{})
	// 创建一条数据用于测试
	db.Create(&domain.User{ID: "1", Username: "张三"})
}

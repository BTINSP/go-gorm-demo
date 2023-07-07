package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	dsn := "host=localhost user=postgres password=postgres dbname=database port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		panic("数据库连接失败")
	}

	//	设置连接池
	pool, err := database.DB()
	if err != nil {
		panic("连接池创建失败")
	}
	//	最大闲置连接数
	pool.SetMaxIdleConns(4)
	//	最大连接数
	pool.SetMaxOpenConns(20)
	

	db = database
}

func GetDB() *gorm.DB {
	return db;
}
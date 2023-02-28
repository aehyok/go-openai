package model

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// 定义一个自定义的命名策略
type MyNamingStrategy struct {
	schema.NamingStrategy
}

// 实现TableName方法，用于返回表名
func (ns MyNamingStrategy) TableName(table string) string {
	return table
}

// 实现ColumnName方法，用于返回列名
func (ns MyNamingStrategy) ColumnName(table, column string) string {
	// 返回列名
	return column
}

var DB *gorm.DB

func Database() {
	// 如何连接数据库 ? MySQL + Navicat
	// 需要更改的内容：用户名，密码，数据库名称
	dsn := "root:M9y2512!@tcp(175.178.60.76:3306)/meta?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: MyNamingStrategy{},
	})

	fmt.Println("db = ", db)
	fmt.Println("err = ", err)

	// 连接池
	sqlDB, err := db.DB()

	if err != nil {
		panic(err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second) // 10秒钟

	// // 迁移
	// db.AutoMigrate(&store.List{})
	DB = db
}

package dao

import (
	"context"
	"pathpro-go/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	_db *gorm.DB
)

func migrate() error {
	return _db.AutoMigrate(User{})
}

func initDB() {
	dsn := conf.GetString("database.dsn")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)    // SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(100)   // SetMaxOpenConns 设置打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(10) // SetConnMaxLifetime 设置了连接可复用的最大时间

	_db = db
	err = migrate()
	if err != nil {
		panic("failed to migrate database")
	}
}

func Init() {
	// 初始化数据库连接
	initDB()
}

func Close() {
	// 关闭数据库连接
	sqlDB, _ := _db.DB()
	sqlDB.Close()
}

func GetDB() *gorm.DB {
	return _db
}

func GetDBWithCtx(ctx context.Context) *gorm.DB {
	return _db.WithContext(ctx)
}

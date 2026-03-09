package init

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"product-srv/config"
	"product-srv/model"
	"time"
)

func MysqlInit() {
	var err error
	data := config.Config.Mysql
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		data.User,
		data.Password,
		data.Host,
		data.Port,
		data.Database,
	)
	config.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}
	fmt.Println("数据库连接成功")
	err = config.DB.AutoMigrate(
		&model.Spu{},
		&model.Sku{},
		&model.Specs{},
		&model.User{},
	)
	if err != nil {
		panic("数据表迁移失败")
	}
	fmt.Println("数据表迁移成功")
	sqlDB, err := config.DB.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
}

package mysql

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Self *gorm.DB
}

var DB *Database

func (db *Database) InitDB() (err error) {
	config := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=%t&loc=%s",
		viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.addr"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.database"),
		true,
		"Local",
	)
	DB = &Database{}
	DB.Self, err = gorm.Open(mysql.Open(config), &gorm.Config{})
	if err != nil {
		return err
	}
	dbset, _ := DB.Self.DB()
	dbset.SetMaxIdleConns(10) // 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。

	DB.autoMigrate() // 自动迁移
	return nil
}

func (db *Database) autoMigrate() {
	db.Self.AutoMigrate(
		&UserTables{},
		&Visitor{},
	)
}

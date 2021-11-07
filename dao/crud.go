package dao

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"envelope_db_writer/entity"
)

var db *gorm.DB

func InitDB() {
	username := os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	dbname := os.Getenv("MYSQL_DBNAME")

	// MYSQL dns格式： {username}:{password}@tcp({host}:{port})/{Dbname}?charset=utf8&parseTime=True&loc=Local
	// 类似{username}使用花括号包着的名字都是需要替换的参数
	dns := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		username, password, host, port, dbname)

	Db, err := gorm.Open("mysql", dns)
	if err != nil {
		panic("failed to connect mysql, error: " + err.Error())
	}

	sqlDB := Db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	db = Db
}

func UpdateOpenState(envelope *entity.Envelope) {
	db.Model(envelope).
		Where("envelope_id = ?", envelope.EnvelopeID).
		Update("opened", envelope.Opened)
}

func InsertEnvelope(envelope *entity.Envelope) {
	if err := db.Create(envelope).Error; err != nil {
		fmt.Println("insert failed: ", err)
		return
	}
}

package mysql

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var Engine *gorm.DB

func Connect(dsn string, maxIdle, maxConn int) (engine *gorm.DB, err error) {
	out := logrus.StandardLogger().Out

	newLogger := logger.New(
		log.New(out, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			Colorful:      true,
			LogLevel:      logger.Info,
		},
	)

	config := gorm.Config{
		Logger:      newLogger,
		PrepareStmt: true,
		QueryFields: true,
	}

	Engine, err = gorm.Open(mysql.Open(dsn), &config)
	if err != nil {
		return nil, err
	}
	if db, err := Engine.DB(); err != nil {
		logrus.WithError(err).Errorf("Engine DB get failure")
	} else {
		db.SetMaxIdleConns(maxIdle)
		db.SetMaxOpenConns(maxConn)
	}
	Engine.Logger.LogMode(logger.Info)

	return Engine, nil

}

func ConnectMysql() error {
	dsn := "root:Jenoayo12!@tcp(127.0.0.1:3306)/dkmk?charset=utf8mb4&parseTime=True&loc=Local"
	if _, err := Connect(dsn, 10, 10); err != nil {
		return err
	} else {
		return nil
	}
}

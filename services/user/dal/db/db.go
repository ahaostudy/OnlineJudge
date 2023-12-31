package db

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"main/services/user/config"
	"main/services/user/dal/model"
	"main/services/user/pkg/sha256"
)

var DB *gorm.DB

// InitMySQL 初始化MySQL
func InitMySQL() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=Local",
		config.Config.Mysql.Username,
		config.Config.Mysql.Password,
		config.Config.Mysql.Host,
		config.Config.Mysql.Port,
		config.Config.Mysql.Dbname,
		config.Config.Mysql.Charset,
	)

	var log logger.Interface
	if gin.Mode() == "debug" {
		log = logger.Default.LogMode(logger.Info)
	} else {
		log = logger.Default
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: log,
	})
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db

	if err = migration(); err != nil {
		return err
	}

	return initAdmin()
}

func migration() error {
	return DB.AutoMigrate(
		new(model.User),
	)
}

func initAdmin() error {
	_, err := GetUser(1)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = InsertUser(&model.User{
			ID:       1,
			Nickname: "admin",
			Username: "admin",
			Password: sha256.Encrypt("admin"),
			Role:     model.ConstRoleOfAdmin,
		})
	}
	return err
}

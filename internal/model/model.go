package model

import (
	"fmt"
	"time"

	"github.com/camtrik/gin-blog/global"
	"github.com/camtrik/gin-blog/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Model struct {
	Id         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	CreatedOn  uint32 `json:"created_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
	ModifiedBy string `json:"modified_by"`
	ModifiedOn uint32 `json:"modified_on"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	var dsn string
	var dialector gorm.Dialector
	switch databaseSetting.DBType {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
			databaseSetting.Username,
			databaseSetting.Password,
			databaseSetting.Host,
			databaseSetting.DBName,
			databaseSetting.Charset,
			databaseSetting.ParseTime,
		)
		dialector = mysql.Open(dsn)
	default:
		return nil, fmt.Errorf("only support mysql database currently")
	}
	db, err := gorm.Open(dialector, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		db = db.Debug()
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// register callbacks
	db.Callback().Create().Before("gorm:create").Register("update_time_stamp_for_create", updateTimeStampForCreateCallback)
	db.Callback().Update().Before("gorm:update").Register("update_time_stamp_for_update", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Before("gorm:delete").Register("delete", deleteCallback)

	sqlDB.SetMaxIdleConns(databaseSetting.MaxIdleConns)
	sqlDB.SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil
}

// model callback
// func updateTimeStampForCreateCallback(db *gorm.DB) {
// 	if db.Error == nil {
// 		nowTime := time.Now().Unix()

// 		if createTimeField := db.Statement.Schema.LookUpField("CreatedOn"); createTimeField != nil {
// 			value, _ := createTimeField.ValueOf(db.Statement.Context, db.Statement.ReflectValue)
// 			if isZero(value) {
// 				createTimeField.Set(db.Statement.Context, db.Statement.ReflectValue, nowTime)
// 			}
// 		}

// 		if modifyTimeField := db.Statement.Schema.LookUpField("ModifiedOn"); modifyTimeField != nil {
// 			value, _ := modifyTimeField.ValueOf(db.Statement.Context, db.Statement.ReflectValue)
// 			if isZero(value) {
// 				modifyTimeField.Set(db.Statement.Context, db.Statement.ReflectValue, nowTime)
// 			}
// 		}
// 	}
// }

func updateTimeStampForCreateCallback(db *gorm.DB) {
	if db.Error == nil {
		nowTime := time.Now().Unix()

		if createTimeField := db.Statement.Schema.LookUpField("CreatedOn"); createTimeField != nil {
			value, _ := createTimeField.ValueOf(db.Statement.Context, db.Statement.ReflectValue)
			if isZero(value) {
				db.Statement.SetColumn("CreatedOn", nowTime)
			}
		}

		if modifyTimeField := db.Statement.Schema.LookUpField("ModifiedOn"); modifyTimeField != nil {
			value, _ := modifyTimeField.ValueOf(db.Statement.Context, db.Statement.ReflectValue)
			if isZero(value) {
				db.Statement.SetColumn("ModifiedOn", nowTime)
			}
		}
	}
}

func updateTimeStampForUpdateCallback(db *gorm.DB) {
	if db.Error == nil {
		nowTime := time.Now().Unix()
		db.Statement.SetColumn("ModifiedOn", nowTime)
	}
}

func deleteCallback(db *gorm.DB) {
	if db.Error == nil {
		now := time.Now().Unix()
		if db.Statement.Schema.LookUpField("DeletedOn") != nil {
			db.Statement.SetColumn("DeletedOn", now)
		}

		if db.Statement.Schema.LookUpField("IsDel") != nil {
			db.Statement.SetColumn("IsDel", 1)
		}
	}
}

func isZero(value interface{}) bool {
	return value == nil || value == 0 || value == ""
}

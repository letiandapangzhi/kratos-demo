package data

import (
	"kratos-demo/app/user/internal/conf"
	"kratos-demo/app/user/internal/data/dal/query"

	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	client *query.Query
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	db, err := gorm.Open(mysql.Open(c.Database.Source))
	if err != nil {
		return nil, nil, err
	}
	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, err
	} else if sqlDB.Ping() != nil {
		return nil, nil, fmt.Errorf("db ping fail")
	}

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		closeErr := sqlDB.Close()
		if closeErr != nil {
			log.NewHelper(logger).Error("closing the data fail", closeErr)
			return
		}
	}
	return &Data{
		client: query.Use(db),
	}, cleanup, nil
}

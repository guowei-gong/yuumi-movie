package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"yuumi-movie/internal/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(NewData, NewDB, NewGreeterRepo, NewUserRepo)

type Data struct {
	db *gorm.DB
}

func NewData(c *conf.Data, logger log.Logger, db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db}, cleanup, nil
}

func NewDB(c *conf.Data) *gorm.DB {
	// e.g Dsn: user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	db, err := gorm.Open(mysql.Open(c.Database.Dsn), &gorm.Config{})
	if err != nil {
		panic("")
	}
	initDB(db)
	return db
}

func initDB(db *gorm.DB) {
	if err := db.AutoMigrate(&User{}); err != nil {
		panic(err)
	}
}

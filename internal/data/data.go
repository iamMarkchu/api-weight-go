package data

import (
	"api-weight-go/internal/conf"
	"api-weight-go/internal/model/orm"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewWeightRepo)

// Data .
type Data struct {
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	db, err := gorm.Open(mysql.Open(c.Database.Source))
	if err != nil {
		panic(err)
	}
	if c.Database.Debug {
		db.AutoMigrate(orm.WeightModel{})
	}
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		db,
	}, cleanup, nil
}

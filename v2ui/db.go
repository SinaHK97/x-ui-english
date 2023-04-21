package v2ui

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var v2db *gorm.DB

func initDB(dbPath string) error {
	c := &gorm.Config{
		Logger: logger.Discard,
	}
	var err error

	dsn := "host=82.115.16.31 user=postgres password=1 dbname=xui port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	v2db, err := gorm.Open(postgres.Open(dsn), c)

	// v2db, err = gorm.Open(sqlite.Open(dbPath), c)
	if err != nil {
		return err
	}

	return nil
}

func getV2Inbounds() ([]*V2Inbound, error) {
	inbounds := make([]*V2Inbound, 0)
	err := v2db.Model(V2Inbound{}).Find(&inbounds).Error
	return inbounds, err
}

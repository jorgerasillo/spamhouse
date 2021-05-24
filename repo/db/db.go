package db

import (
	"github.com/cenkalti/backoff"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) error {
	return db.AutoMigrate(IPAddress{})
}

func GetDB(dataSourceName string) (db *gorm.DB, err error) {

	getDB := func() error {
		db, err = gorm.Open(sqlite.Open(dataSourceName), &gorm.Config{})
		if err != nil {
			return err
		}
		return nil
	}

	err = backoff.Retry(getDB, backoff.NewExponentialBackOff())
	if err == nil {
		err = migrate(db)
	}
	return
}

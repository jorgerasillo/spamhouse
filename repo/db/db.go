package db

import (
	"github.com/cenkalti/backoff"
	"github.com/jorgerasillo/spamhouse/graph/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) error {
	return db.AutoMigrate(model.IPAddress{})
}

func GetDB(dataSourceName string) (db *gorm.DB, err error) {

	getDB := func() error {
		db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
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

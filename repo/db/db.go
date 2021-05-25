package db

import (
	"github.com/cenkalti/backoff"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB, defaultUser string, defaultPassword string) error {
	if err := db.AutoMigrate(IPAddress{}, User{}); err != nil {
		return err
	}
	u := User{
		UserID:   defaultUser,
		Password: defaultPassword,
	}
	// the create statement below is mainly for bootstrapping a default user
	db.Create(&u)
	return nil
}

func GetDB(dataSourceName string, defaultUser string, defaultPassword string) (db *gorm.DB, err error) {

	getDB := func() error {
		db, err = gorm.Open(sqlite.Open(dataSourceName), &gorm.Config{})
		if err != nil {
			return err
		}
		return nil
	}

	err = backoff.Retry(getDB, backoff.NewExponentialBackOff())
	if err == nil {
		err = migrate(db, defaultUser, defaultPassword)
	}
	return
}

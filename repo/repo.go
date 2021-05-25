package repo

import (
	"errors"

	"github.com/jorgerasillo/spamhouse/repo/db"
	"gorm.io/gorm"
)

type Repository interface {
	GetIP(ip string) (db.IPAddress, error)
	AddIP(ip db.IPAddress) (db.IPAddress, error)
	UpdateIP(ip db.IPAddress) (db.IPAddress, error)
	IsUserValid(userID string, password string) bool
}

type repo struct {
	DB *gorm.DB
}

// New creates new store repository, used to abstract and enable mocking db connection
func New(db *gorm.DB) (Repository, error) {
	return &repo{
		DB: db,
	}, nil
}

func (r repo) AddIP(ip db.IPAddress) (db.IPAddress, error) {
	tx := r.DB.Begin()
	if err := tx.Create(&ip); err.Error != nil {
		tx.Rollback()
		return ip, err.Error
	}
	if err := tx.Commit(); err.Error != nil {
		tx.Rollback()
		return ip, err.Error
	}

	return ip, nil
}

func (r repo) UpdateIP(ip db.IPAddress) (db.IPAddress, error) {
	tx := r.DB.Begin()
	if err := tx.Save(&ip); err.Error != nil {
		tx.Rollback()
		return ip, err.Error
	}
	if err := tx.Commit(); err.Error != nil {
		tx.Rollback()
		return ip, err.Error
	}

	return ip, nil
}

func (r repo) GetIP(IP string) (db.IPAddress, error) {
	var ipAddress db.IPAddress
	if err := r.DB.Where("ip = ?", IP).First(&ipAddress); err.Error != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			return ipAddress, db.ErrRecordNotFound
		}
		return ipAddress, err.Error
	}
	return ipAddress, nil
}

func (r repo) IsUserValid(userID string, password string) bool {
	var u *db.User
	if err := r.DB.Where("user_id = ?", userID).Where("password = ?", password).First(&u); err.Error != nil {
		return false
	}
	if u != nil {
		return true
	}
	return false
}

package repo

import (
	"github.com/jorgerasillo/spamhouse/graph/model"
	"gorm.io/gorm"
)

type Repository interface {
	GetIP(ip string) (model.IPAddress, error)
	AddIP(ip model.IPAddress) (model.IPAddress, error)
	UpdateIP(ip model.IPAddress) (model.IPAddress, error)
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

func (r repo) AddIP(ip model.IPAddress) (model.IPAddress, error) {
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

func (r repo) UpdateIP(ip model.IPAddress) (model.IPAddress, error) {
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

func (r repo) GetIP(IP string) (model.IPAddress, error) {
	var ipAddress model.IPAddress
	if err := r.DB.First(&ipAddress); err.Error != nil {
		return ipAddress, err.Error
	}
	return ipAddress, nil
}

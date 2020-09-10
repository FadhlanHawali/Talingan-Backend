package db

import "gorm.io/gorm"

type ServicesDB struct {
	db *gorm.DB
}

func NewServicesDB(db *gorm.DB) *ServicesDB{
	return &ServicesDB{
		db: db,
	}
}
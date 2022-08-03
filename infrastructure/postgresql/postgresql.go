package postgresql

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB *gorm.DB

func NewDB(dsn string) (DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err
}

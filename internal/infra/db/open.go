package db

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Driver string
	DSN    string
}

func Open(cfg Config) (*gorm.DB, error) {
	var dial gorm.Dialector
	switch cfg.Driver {
	case "mysql":
		dial = mysql.Open(cfg.DSN)
	default:
		return nil, fmt.Errorf("unsupported driver: %s", cfg.Driver)
	}
	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Println("gorm connected ->", cfg.Driver)
	return db, nil
}

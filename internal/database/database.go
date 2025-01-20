package database

import (
	"context"
	"fliqt/internal/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func New() *Database {
	db := &Database{
		config: config.New(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30)
	defer cancel()
	for {
		select {
		case <-ctx.Done():
		default:
		}
		if err := db.connect(); err != nil {
			log.Printf("Failed to connect to MySQL: %v", err)
			<-time.After(1 * time.Second)
			continue
		}
		break
	}

	return db
}

type Database struct {
	*gorm.DB
	config *config.Config
}

func (d *Database) connect() error {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&time_zone=UTC",
		d.config.MySQL.User,
		d.config.MySQL.Password,
		d.config.MySQL.Host,
		d.config.MySQL.Port,
		d.config.MySQL.DBName,
	)
	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	d.DB = db

	return nil
}

package migration

import (
	"context"
	"fliqt/internal/config"
	"fliqt/internal/database"
	"fliqt/internal/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func New() *Migration {
	return &Migration{
		config: config.New(),
	}
}

type Migration struct {
	config *config.Config
}

func (m *Migration) Migrate(ctx context.Context) {
	if m.createDatabase(ctx) {
		db := database.New()

		for _, schema := range m.schema() {
			db.AutoMigrate(schema)
		}

		newSeed().ExecAll(ctx)
	}

}

func (m *Migration) schema() []any {
	return []any{
		&model.Employee{},
	}
}

func (m *Migration) createDatabase(ctx context.Context) bool {
	var (
		db *gorm.DB
	)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	for {
		select {
		case <-ctx.Done():
		default:
		}
		if conn, err := m.connect(); err != nil {
			log.Printf("Failed to connect to MySQL: %v", err)
			<-time.After(1 * time.Second)
			continue
		} else {
			db = conn
			break
		}
	}

	// 創建資料庫
	sql := "CREATE DATABASE IF NOT EXISTS " + m.config.MySQL.DBName
	if err := db.Exec(sql).Error; err != nil {
		log.Panicf("Failed to create database: %v", err)
	}
	log.Printf("Database %s created successfully!", m.config.MySQL.DBName)

	return true
}

func (m *Migration) connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local",
		m.config.MySQL.User,
		m.config.MySQL.Password,
		m.config.MySQL.Host,
		m.config.MySQL.Port,
	)
	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

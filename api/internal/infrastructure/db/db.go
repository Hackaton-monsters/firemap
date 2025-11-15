package db

import (
	"database/sql"
	"firemap/internal/infrastructure/config"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func NewDBForMigrations(config *config.Config) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	db.SetMaxOpenConns(config.MaxOpenConn)
	db.SetMaxIdleConns(config.MaxIdleConn)
	db.SetConnMaxLifetime(time.Duration(config.ConnMaxLifetime) * time.Minute)

	return db
}

func NewDB(config *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	db.SetMaxOpenConns(config.MaxOpenConn)
	db.SetMaxIdleConns(config.MaxIdleConn)
	db.SetConnMaxLifetime(time.Duration(config.ConnMaxLifetime) * time.Minute)

	g, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			NameReplacer: schema.NamingStrategy{}.NameReplacer,
			NoLowerCase:  false,
		},
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	return g
}

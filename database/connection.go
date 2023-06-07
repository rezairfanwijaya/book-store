package database

import (
	"book-store/admin"
	"book-store/helper"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewConnection(envPath string) (*gorm.DB, error) {
	env, err := helper.GetENV(envPath)
	if err != nil {
		return &gorm.DB{}, err
	}

	userName := env["USERNAME"]
	password := env["PASSWORD"]
	host := env["HOST"]
	port := env["PORT"]
	databaseName := env["DATABASE_NAME"]

	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		userName,
		password,
		host,
		port,
		databaseName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, fmt.Errorf("failed connect to database : %v", err.Error())
	}

	// migrasi schema
	if err := db.AutoMigrate(&admin.Admin{}); err != nil {
		return db, fmt.Errorf("error migaration schema : %v", err.Error())
	}

	// migrasi admin
	if err := AdminMigration(db); err != nil {
		return db, err
	}

	return db, nil
}

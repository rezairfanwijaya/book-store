package database

import (
	"book-store/admin"
	"book-store/helper"
	"fmt"

	"gorm.io/gorm"
)

func AdminMigration(db *gorm.DB) error {
	// cek apakah sudah ada data admin
	adminExist, err := isAdminExist(db)
	if err != nil {
		return fmt.Errorf("failed check data admin : %v", err.Error())
	}

	if !adminExist {
		passEncrypt, err := helper.EcryptPassword("admin12345")
		if err != nil {
			return err
		}

		newAdmin := admin.Admin{
			Email:    "admin@gmil.com",
			Password: passEncrypt,
			Role:     "admin",
		}

		if err := db.Create(&newAdmin).Error; err != nil {
			return fmt.Errorf("failed save data admin : %v", err.Error())
		}
	}

	return nil
}

func isAdminExist(db *gorm.DB) (bool, error) {
	var admin admin.Admin
	if err := db.Find(&admin).Error; err != nil {
		return true, err
	}

	if admin.ID == 0 {
		return false, nil
	}

	return true, nil
}

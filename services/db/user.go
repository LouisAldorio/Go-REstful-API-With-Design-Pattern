package services

import (
	"fmt"
	"myapp/models"

	"gorm.io/gorm"
)

func (db *db) UserGetAll() ([]*models.User, error) {

	var users []*models.User
	err := db.Tx.Find(&users).Error
	if err != nil {
		db.Tx.Rollback()
		panic(err)
	}

	return users, nil
}

func (d *db) UserCreate(input models.User) (*models.User, error) {

	err := d.Tx.Model(&models.User{}).Create(&input).Error
	if err != nil {
		d.Tx.Rollback()
		panic(err)
	}

	return &input, nil
}

func (d *db) UserGetByID(id int) (*models.User, error) {

	var user models.User
	err := d.Tx.Model(&models.User{}).Where("id = ?", id).Take(&user).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			d.Tx.Rollback()
			panic(err)
		} else {
			d.Tx.Rollback()
			return nil, fmt.Errorf("user not available")
		}
	}

	return &user, nil
}

func (d *db) UserGetByEmail(email string) (*models.User, error) {

	var user models.User
	err := d.Tx.Model(&models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			d.Tx.Rollback()
			panic(err)
		} else {
			return nil, fmt.Errorf("user not available")
		}
	}

	return &user, nil
}

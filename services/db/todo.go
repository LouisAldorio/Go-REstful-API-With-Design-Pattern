package services

import (
	"fmt"
	"myapp/models"

	"gorm.io/gorm"
)

func (d *db) TodoCreate(input models.Todo) (*models.Todo, error) {

	err := d.Tx.Model(&models.Todo{}).Create(&input).Error
	if err != nil {
		d.Tx.Rollback()
		panic(err)
	}

	return &input, nil
}

func (d *db) TodoGetByWhereInUserIDs(userIDs []int) ([]*models.Todo, error) {

	var todos []*models.Todo
	err := d.Tx.Where("user_id IN (?)", userIDs).Find(&todos).Error
	if err != nil {
		d.Tx.Rollback()
		panic(err)
	}

	return todos, nil
}

func (d *db) TodoGetByID(id int) (*models.Todo, error) {

	var todo models.Todo
	err := d.Tx.Model(&models.Todo{}).Where("id = ?", id).Take(&todo).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			d.Tx.Rollback()
			panic(err)
		} else {
			d.Tx.Rollback()
			return nil, fmt.Errorf("todo not available")
		}
	}

	return &todo, nil
}

func (d *db) TodoGetByUserID(userID int) ([]*models.Todo, error) {

	var todos []*models.Todo
	err := d.Tx.Model(&models.Todo{}).Where("user_id = ?", userID).Find(&todos).Error
	if err != nil {
		d.Tx.Rollback()
		panic(err)
	}

	return todos, nil
}

func (d *db) TodoUpdate(input models.Todo) (*models.Todo, error) {

	err := d.Tx.Model(&models.Todo{}).Where("id = ?", input.ID).Updates(&input).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			d.Tx.Rollback()
			panic(err)
		} else {
			d.Tx.Rollback()
			return nil, fmt.Errorf("todo not available")
		}
	}

	return &input, nil
}

func (d *db) TodoDelete(id int) (*int, error) {

	err := d.Tx.Model(&models.Todo{}).Where("id = ?", id).Delete(&models.Todo{}).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			d.Tx.Rollback()
			panic(err)
		} else {
			d.Tx.Rollback()
			return nil, fmt.Errorf("todo not available")
		}
	}

	return &id, nil
}

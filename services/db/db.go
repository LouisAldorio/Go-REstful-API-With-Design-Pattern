package services

import (
	"myapp/config"
	"myapp/interfaces"

	"gorm.io/gorm"
)

var (
	Database interfaces.DBInterface
)

type db struct {
	Tx *gorm.DB
}

func BeginTransaction() *gorm.DB {
	currentTransaction := config.GetDB().Begin()
	Database = &db{
		Tx: currentTransaction,
	}

	return currentTransaction
}




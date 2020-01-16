package util

import (
	"github.com/jinzhu/gorm"
)

func Transact(db *gorm.DB, txFunc func(*gorm.DB) (interface{}, error)) (data interface{}, err error) {
	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit().Error
		}
	}()
	data, err = txFunc(tx)
	return
}

func TransactTwo(db *gorm.DB, txFunc func(*gorm.DB) (interface{}, interface{}, error)) (data1 interface{}, data2 interface{}, err error) {
	tx := db.Begin()
	if tx.Error != nil {
		return nil, nil, tx.Error
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit().Error
		}
	}()
	data1, data2, err = txFunc(tx)
	return
}

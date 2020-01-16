package util_test

import (
	"testing"

	"github.com/jinzhu/gorm"
	util "github.com/maruware/gorm-util"
)

func TestTransaction(t *testing.T) {
	db, err := buildDatabase()
	if err != nil {
		t.Fatalf("failed to build database: %v", err)
	}

	user, err := util.Transact(db, func(tx *gorm.DB) (interface{}, error) {
		user := User{Name: "takashi"}
		if err := tx.Create(&user).Error; err != nil {
			return nil, err
		}

		post := Post{UserID: user.ID, Title: "My post"}
		if err := tx.Create(&post).Error; err != nil {
			return nil, err
		}

		return user, nil
	})

	if err != nil {
		t.Fatalf("failed transaction")
	}

	u, ok := user.(User)
	if !ok {
		t.Fatalf("failed type assertion")
	}
	if u.ID <= 0 {
		t.Errorf("expect to regist user")
	}
}

func TestTransactionTwo(t *testing.T) {
	db, err := buildDatabase()
	if err != nil {
		t.Fatalf("failed to build database: %v", err)
	}

	user, post, err := util.TransactTwo(db, func(tx *gorm.DB) (interface{}, interface{}, error) {
		user := User{Name: "takashi"}
		if err := tx.Create(&user).Error; err != nil {
			return nil, nil, err
		}

		post := Post{UserID: user.ID, Title: "My post"}
		if err := tx.Create(&post).Error; err != nil {
			return nil, nil, err
		}

		return user, post, nil
	})

	if err != nil {
		t.Fatalf("failed transaction")
	}

	u, ok := user.(User)
	if !ok {
		t.Fatalf("failed type assertion")
	}
	if u.ID <= 0 {
		t.Errorf("expect to regist user")
	}

	p, ok := post.(Post)
	if !ok {
		t.Fatalf("failed type assertion")
	}
	if p.ID <= 0 {
		t.Errorf("expect to regist post")
	}
}

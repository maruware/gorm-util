package util_test

import (
	util "github.com/maruware/gorm-util"
	"testing"
)

func TestPreloads(t *testing.T) {
	db, err := buildDatabase()
	if err != nil {
		t.Fatalf("failed to build database: %v", err)
	}

	user := User{Name: "takashi"}
	db.Create(&user)
	post1 := Post{UserID: user.ID, Title: "My post1"}
	db.Create(&post1)
	post2 := Post{UserID: user.ID, Title: "My post2"}
	db.Create(&post2)

	var u User
	q := util.Preloads(db, []string{"Posts"})
	if err := q.First(&u, user.ID).Error; err != nil {
		t.Fatalf("faild to find post: %v", err)
	}

	if len(u.Posts) != 2 {
		t.Errorf("expect user.Posts is preloaded")
	}
}

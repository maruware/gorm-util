package util

import "github.com/jinzhu/gorm"

func Preloads(db *gorm.DB, preloads []string) *gorm.DB {
	r := db
	for _, preload := range preloads {
		r = r.Preload(preload)
	}
	return r
}

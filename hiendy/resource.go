package hiendy

import (
	"github.com/jinzhu/gorm"
)

type Resource struct {
	gorm.Model

	Url  string
	From string
	Type string
}

// const (
// 	Image State = iota // value --> 0
// 	Video              // value --> 1
// 	Comic              // value --> 2
// 	Novel              // value --> 3
// )

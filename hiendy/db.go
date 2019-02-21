package hiendy

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	// "sync"
)

var connection *gorm.DB

// var mu sync.Mutex

func GetConnection() (db *gorm.DB, err error) {
	if connection != nil {
		return connection, nil
	}

	db, err = gorm.Open("mysql", "root@tcp(localhost:13307)/search?multiStatements=true&parseTime=True")
	if err != nil {
		log.Fatal("database connect failed", err)
	}

	db.SingularTable(true)
	db.LogMode(true)

	connection = db
	return db, err
}

// func GetSigle() *gorm.DB {
// 	if connection == nil {
// 		mu.Lock()
// 		defer mu.Unlock()
// 		if connection == nil {
// 			connection = &Program{Id: "update" + id, IsLws: false}
// 			cliProgram.Init()
// 			if err := connection.Start(); err != nil {
// 				log.Printf("client start failed")
// 			}
// 		}
// 	}
// 	return connection
// }

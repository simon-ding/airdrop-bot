package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	MetaIndex  int    `json:"meta_index"`
	Address    string `json:"address"`
	PrivateKey string `json:"private_key"`
	BindIp     string `json:"bind_ip"`
}

func Open() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

}

package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SetUpModels() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@(localhost)/tugaswebapigolang_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Error = Gagal koneksi ke dalam database!")
	}
	db.AutoMigrate(&Mahasiswa{}, &MataKuliah{})
	return db
}

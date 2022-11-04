package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
  
var db *gorm.DB
func init() {
    dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
      panic("failed to connect database")
    }
    db.AutoMigrate(&User{})
    db.AutoMigrate(&Todo{})
}
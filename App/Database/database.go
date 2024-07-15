package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"fmt"
	"errors"
	"App/Model"

)

var db *gorm.DB

func DatabaseConnection() (*gorm.DB, error) {
	//sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 3306, "root", "Joker", "Students")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", "root", "Joker", "localhost", 3306, "Students")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Errorf("failed to connect database")
		return nil, errors.New("Failed to Connect to DB:" + err.Error())
	}

	fmt.Println("Database connected")
	db.AutoMigrate(Model.Student{})
	
	return db, nil
}
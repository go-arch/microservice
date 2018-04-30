package services

import (
	_ "github.com/go-sql-driver/mysql"

	"students/util"
	"students/models"
	"github.com/jinzhu/gorm"
)


//var DB *sql.DB
var DB *gorm.DB

// db session
func GetDbSession() *gorm.DB{
	db, err := gorm.Open("mysql", "userms:randompass@tcp(userms.ctm6uenobr4l.us-west-2.rds.amazonaws.com:3306)/userms?charset=utf8&parseTime=true")
	util.Fatal(err)
	DB = db;
	return DB
}


func Migration(){
	DB.AutoMigrate(&models.User{})
}

func CloseDB() {
	DB.Close()
}

//func BootStrapDB(){
//	DB.Query("create table users(`id` INT(10) NOT NULL AUTO_INCREMENT,`name` VARCHAR(64) NULL DEFAULT NULL,`email` VARCHAR(64) NULL DEFAULT NULL,`password` VARCHAR(64) NULL DEFAULT NULL,`phoneNumber` VARCHAR(64) NULL DEFAULT NULL,PRIMARY KEY (`id`))")
//
//}
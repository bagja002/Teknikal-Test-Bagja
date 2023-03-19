package database

import (
	"fmt"
	"khub/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	dsn := "root:@tcp(127.0.0.1:3306)/kstyle_dev?charset=utf8mb4&parseTime=True&loc=Local"

	con, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Gagal Meng koneksikan database")
	}
	fmt.Println("database terkoneksi")
	DB = con
	err1:= DB.AutoMigrate(models.Member{}, models.Like_review{}, models.Review_product{}, models.Product{})
	if err1!= nil {
        fmt.Println("gagal Auto Migration")
    } else {
		fmt.Println("berhasil Auto Migration")
	}

}

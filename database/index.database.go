package database

import (
	"fmt"
	"gin-gorm/app/model"
	"gin-gorm/configs/db_config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var errConnection error
	fmt.Println(db_config.DB_DRIVER)
	if db_config.DB_DRIVER == "mysql" {
		dsnMysql := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db_config.DB_USER, db_config.DB_PASSWORD, db_config.DB_HOST, db_config.DB_PORT, db_config.DB_NAME)
		fmt.Println(dsnMysql)
		DB, errConnection = gorm.Open(mysql.Open(dsnMysql), &gorm.Config{})
	}

	if errConnection != nil {
		panic("Canot connect db mysql")
	}

	log.Println("Connected to DATABASE")

	data_karyawan := new(model.Karyawan)
	err := DB.Table("karyawan").
		Where("karyawan_email = ?", "tries1@mail.com").
		First(&data_karyawan).Error

	if err != nil {
		log.Println("ERR DB karyawan")
	} else {
		log.Println("Success DB karyawan" + *data_karyawan.KaryawanNama)
	}
}

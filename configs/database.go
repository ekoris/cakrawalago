package configs

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDB() {
	var err error
	dsn := "epiz_33738354:WhzHNhAitC@tcp(sql100.epizy.com:3306)/epiz_33738354_cakrawala?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}

	fmt.Println("ðŸš€ Connected Successfully to the Database")
}

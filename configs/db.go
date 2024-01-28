package configs

import (
	"fmt"
	"os"

	"example.com/gin-backend-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	dsn := os.Getenv("DATABASE_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("ไม่สามารถติดต่อกับ Database Server ได้")
		fmt.Println(err.Error())
		panic(err)
	}

	fmt.Println("ติดต่อฐานข้อมูลสำเร็จ")

	// Migration
	// db.Migrator().DropTable(&models.User{})

	db.AutoMigrate(&models.User{}, &models.Blog{})

	DB = db
}

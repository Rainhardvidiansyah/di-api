package main

import (
	"fmt"
	"g-mysql/entity"
	"g-mysql/handler"
	"g-mysql/repository"
	"g-mysql/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:maulida1010@tcp(127.0.0.1:3306)/web_apis?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("DB Connected")
	db.AutoMigrate(entity.Product{}, entity.User{})

	t := repository.Newrepository(db)
	y := service.NewService(t)
	z := handler.NewHandlerProduct(y)
	r := gin.Default()
	r.POST("/save/product", z.SaveData)
	r.GET("/api/find/all", z.FindAllData)

	fmt.Println("App is running on Port 9292")
	r.Run("localhost:9292")

}

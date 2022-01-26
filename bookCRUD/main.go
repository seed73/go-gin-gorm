package main

import (
	"bookCRUD/models"
	"bookCRUD/routes"
	"fmt"
)

/* 아래 항목이 swagger에 의해 문서화 된다. */
// @title Go Gin Swagger Example API
// @version 1.0
// @description GO Gin Gorm
// @host localhost:8080
// @BasePath /
func main() {
	//DB연결
	db := models.SetupDB()

	//DB자동생성
	db.AutoMigrate(&models.Task{})

	//Router설정
	r := routes.SetupRoutes(db)
	r.Run()

	fmt.Print()
}

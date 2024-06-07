package main

import (
	"latihan-restfull-api-go/app"
	"latihan-restfull-api-go/controller"
	"latihan-restfull-api-go/helper"
	"latihan-restfull-api-go/middleware"
	"latihan-restfull-api-go/repository"
	"latihan-restfull-api-go/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	// router := httprouter.New()

	// router.GET("/api/categories", categoryController.FindAll)
	// router.GET("/api/categories/:categoryId", categoryController.FindById)
	// router.POST("/api/categories", categoryController.Create)
	// router.PUT("/api/categories/:categoryId", categoryController.Update)
	// router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}
	err := server.ListenAndServe() //menjalankan server
	helper.PanicIfError(err)

}

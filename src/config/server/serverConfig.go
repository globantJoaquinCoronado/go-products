package server

import (
	"log"
	"net/http"
	"products/src/controller"
	"products/src/repository"
	"products/src/service"

	"github.com/gorilla/mux"
)

func StartServer() {

	productRepository := repository.NewProductRepositoryImpl()
	productService := service.NewProductServiceImpl(productRepository)
	productController := controller.NewProductControllerImpl(productService)

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/product/v1/product/{productId}", productController.GetById).Methods(http.MethodGet)
	r.HandleFunc("/product/v1/product", productController.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/product/v1/product", productController.Create).Methods(http.MethodPost)
	r.HandleFunc("/product/v1/product", productController.Update).Methods(http.MethodPut)
	r.HandleFunc("/product/v1/product/{productId}", productController.Delete).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8000", r))
}

package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"products/src/model"
	"products/src/service"
	"products/src/utils"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductControllerImpl struct {
	ProductService service.IProductService
}

func NewProductControllerImpl(productService service.IProductService) *ProductControllerImpl {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (p *ProductControllerImpl) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := p.ProductService.GetAll()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJson(w, http.StatusOK, products)
}

func (p *ProductControllerImpl) GetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productIdStr := params["productId"]
	productIdInt, err := strconv.Atoi(productIdStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	productId := uint(productIdInt)

	product, err := p.ProductService.GetById(productId)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if product == nil {
		params := map[string]string{
			"message": "Product not found",
		}
		utils.RespondWithJson(w, http.StatusNotFound, params)
		return
	}

	utils.RespondWithJson(w, http.StatusOK, product)
}

func (p *ProductControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	var product model.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Printf("body: %+v", product)

	err = p.ProductService.Create(&product)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJson(w, http.StatusCreated, product)
}

func (p *ProductControllerImpl) Update(w http.ResponseWriter, r *http.Request) {

	var product model.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Printf("body: %+v", product)

	err = p.ProductService.Update(product)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	updatedProduct, err := p.ProductService.GetById(product.Id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	utils.RespondWithJson(w, http.StatusOK, updatedProduct)
}

func (p *ProductControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productIdStr := params["productId"]
	productIdInt, err := strconv.Atoi(productIdStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	productId := uint(productIdInt)

	product, err := p.ProductService.GetById(productId)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if product == nil {
		params := map[string]string{
			"message": "Product not found",
		}
		utils.RespondWithJson(w, http.StatusNotFound, params)
		return
	}

	p.ProductService.Delete(productId)

	utils.RespondWithJson(w, http.StatusOK, product)
}

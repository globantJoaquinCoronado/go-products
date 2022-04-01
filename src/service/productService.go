package service

import (
	"products/src/model"
	"products/src/repository"
	"time"
)

type IProductService interface {
	GetAll() ([]model.Product, error)
	GetById(uint) (*model.Product, error)
	Create(*model.Product) error
	Update(model.Product) error
	Delete(uint) error
}

type ProductServiceImpl struct {
	ProductRepository repository.IProductRepository
}

func NewProductServiceImpl(productRepository repository.IProductRepository) *ProductServiceImpl {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
	}
}

func (p *ProductServiceImpl) GetAll() ([]model.Product, error) {
	return p.ProductRepository.GetAll()
}

func (p *ProductServiceImpl) GetById(id uint) (*model.Product, error) {
	return p.ProductRepository.GetById(id)
}

func (p *ProductServiceImpl) Create(product *model.Product) error {
	product.Id = uint(time.Now().UnixNano())
	return p.ProductRepository.Create(product)
}

func (p *ProductServiceImpl) Update(product model.Product) error {

	currentProduct, err := p.GetById(product.Id)
	if err != nil {
		return err
	}

	mergedProduct := p.merge(*currentProduct, product)
	return p.ProductRepository.Update(mergedProduct)
}

func (p *ProductServiceImpl) Delete(id uint) error {
	return p.ProductRepository.Delete(id)
}

func (p *ProductServiceImpl) merge(currentProduct, updateProduct model.Product) model.Product {

	if updateProduct.Name != "" {
		currentProduct.Name = updateProduct.Name
	}

	if updateProduct.SupplierId != 0 {
		currentProduct.SupplierId = updateProduct.SupplierId
	}

	if updateProduct.CategoryId != 0 {
		currentProduct.CategoryId = updateProduct.CategoryId
	}

	if updateProduct.UnitInStock != 0 {
		currentProduct.UnitInStock = updateProduct.UnitInStock
	}

	if updateProduct.UnitPrice != 0 {
		currentProduct.UnitPrice = updateProduct.UnitPrice
	}

	if updateProduct.Discontinued != currentProduct.Discontinued {
		currentProduct.Discontinued = updateProduct.Discontinued
	}

	return currentProduct
}

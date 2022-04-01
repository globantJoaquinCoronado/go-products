package repository_test

import (
	"products/src/model"
	"products/src/repository"
	"testing"
)

func TestCreate(t *testing.T) {
	repository := repository.NewProductRepositoryImpl()
	repository.Create(&model.Product{
		Id:           1,
		Name:         "Play Station 5",
		SupplierId:   1,
		CategoryId:   1,
		UnitInStock:  4,
		UnitPrice:    15000.00,
		Discontinued: false,
	})
}

package repository

import (
	"context"
	"fmt"
	"products/src/config/mongodb"
	"products/src/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IProductRepository interface {
	GetAll() ([]model.Product, error)
	GetById(uint) (*model.Product, error)
	Create(*model.Product) error
	Update(model.Product) error
	Delete(uint) error
}

var ctx = context.Background()

type ProductRepositoryImpl struct {
	collection *mongo.Collection
}

func NewProductRepositoryImpl() *ProductRepositoryImpl {
	collection := mongodb.GetCollection("product")
	return &ProductRepositoryImpl{
		collection: collection,
	}
}

func (p *ProductRepositoryImpl) GetAll() ([]model.Product, error) {
	products := make([]model.Product, 0)
	cursor, err := p.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var product model.Product
		err = cursor.Decode(&product)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (p *ProductRepositoryImpl) GetById(id uint) (*model.Product, error) {

	var product model.Product
	err := p.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return nil, nil
		}
		return nil, err
	}

	return &product, nil
}

func (p *ProductRepositoryImpl) Create(product *model.Product) error {
	_, err := p.collection.InsertOne(ctx, product)
	if err != nil {
		fmt.Println("Error saving product")
		return err
	}
	fmt.Println("Product was created")
	return nil
}

func (p *ProductRepositoryImpl) Update(product model.Product) error {

	filter := bson.M{"_id": product.Id}
	update := bson.D{{
		Key: "$set", Value: bson.D{
			{Key: "name", Value: product.Name},
			{Key: "supplierId", Value: product.SupplierId},
			{Key: "categoryId", Value: product.CategoryId},
			{Key: "unitInStock", Value: product.UnitInStock},
			{Key: "unitPrice", Value: product.UnitPrice},
			{Key: "discontinued", Value: product.Discontinued},
		},
	}}

	_, err := p.collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	return nil
}

func (p *ProductRepositoryImpl) Delete(id uint) error {
	_, err := p.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}

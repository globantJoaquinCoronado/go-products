package model

type Product struct {
	Id           uint    `json:"id,omitempty" bson:"_id,omitempty"`
	Name         string  `json:"name,omitempty" bson:"name,omitempty"`
	SupplierId   uint    `json:"supplierId,omitempty" bson:"supplierId,omitempty"`
	CategoryId   uint    `json:"categoryId,omitempty" bson:"categoryId,omitempty"`
	UnitInStock  uint    `json:"unitInStock,omitempty" bson:"unitInStock,omitempty"`
	UnitPrice    float64 `json:"unitPrice,omitempty" bson:"unitPrice,omitempty"`
	Discontinued bool    `json:"discontinued,omitempty" bson:"discontinued,omitempty"`
}

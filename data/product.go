package data

import "time"

// Product define product
type Product struct {
	ID          int
	Name        string
	Description string
	Price       float32
	SKU         string
	CreatedOn   string
	UpdatedOn   string
	DeletedOn   string
}

var products = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Froth  milky coffee",
		Price:       1.92,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       2.02,
		SKU:         "dfk123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}

// GetProducts returns products
func GetProducts() []*Product {
	return products
}

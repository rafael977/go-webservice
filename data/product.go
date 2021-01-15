package data

import (
	"regexp"
	"time"

	"github.com/go-playground/validator"
)

// Product define product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku"  validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
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

func AddProduct(p *Product) {
	p.ID = len(products) + 1
	products = append(products, p)
}

func UpdateProduct(id int, p *Product) {
	p.ID = id
	products[id-1] = p
}

func (p *Product) Validate() error {
	validator := validator.New()
	if err := validator.RegisterValidation("sku", validateSKU); err != nil {
		return err
	}

	return validator.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	re := regexp.MustCompile("[a-z]+-[a-z]+-[a-z]+")
	matches := re.FindAllString(fl.Field().String(), -1)

	return len(matches) > 0
}

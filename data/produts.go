package data

import (
	"errors"
	"regexp"
	"time"

	"github.com/go-playground/validator"
)

//Product defines the structure for an API product
// swagger:model
type Product struct {
	// the id for this product
	//
	// required: true
	//
	// example: 1
	ID int `json:"id" validate:"required"`

	// the name for this product
	//
	// required: true
	//
	// example: string
	Name string `json:"name" validate:"required"`

	// the Description for this product
	//
	// required: true
	// example: string
	Description string `json:"description" validate:"required"`

	// the price for this product
	//
	// required: true
	// min: 1
	//
	// example: 1.0
	Price float64 `json:"price" validate:"gt=0"`

	// the sku for this product
	//
	// required: true
	//
	// example: wwe-wea-aweasdas
	SKU      string `json:"sku" validate:"required,sku"`
	CreateOn string `json:"-"`
	UpdateOn string `json:"-"`
	DeleteOn string `json:"-"`
}

func validateSku(fl validator.FieldLevel) bool {
	rg := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := rg.FindAllString(fl.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}

	return true
}

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSku)
	err := validate.Struct(p)
	return err
}

func Delete(id int) error {
	pr, i := FindById(id)
	if pr == nil {
		return errors.New("Product not found")
	}
	produts = append(produts[:i], produts[i+1:]...)
	return nil
}

func Save(prod *Product) error {
	produts = append(produts, prod)
	return nil
}

func Update(prod *Product, id int) error {
	pr, i := FindById(id)
	if pr == nil {
		return errors.New("Product not found")
	}

	produts[i] = prod
	return nil
}

func FindById(id int) (*Product, int) {
	for i, p := range produts {
		if p.ID == id {
			return p, i
		}
	}
	return nil, -1
}

func FindAll() []*Product {
	return produts
}

var produts = []*Product{
	&Product{
		ID:          1,
		Name:        "Book- Go programming",
		Description: "Best pratices of write code Go",
		Price:       1.82,
		SKU:         "KFD123",
		CreateOn:    time.Now().String(),
		UpdateOn:    time.Now().String(),
	}, &Product{
		ID:          2,
		Name:        "Book- Documentation API",
		Description: "Best pratices to documented your microservice",
		Price:       1.82,
		SKU:         "KFD123",
		CreateOn:    time.Now().String(),
		UpdateOn:    time.Now().String(),
	},
}

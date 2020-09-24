package handlers

import (
	"encoding/json"
	"example01/data"
	"io"
	"net/http"
)

// A list of products
// swagger:response productsResponse
type productsResponse struct {
	//	All products in the system
	//	in: body
	Body []data.Product
}

// swagger:parameters deleteProduct updateProduct
type productIdParamater struct {
	//The id of the product to delete from the database
	// in: path
	// required:true
	ID int `json:"id"`
}

// swagger:parameters saveProducts updateProduct
type productSaveParam struct {
	//	Body to save product in database
	// in: body
	Product data.Product
}

// swagger:response noContent
type noContent struct {
}

//Produts
type Produts struct {
}

func NewProduts() Produts {
	return Produts{}
}

// 	swagger:route GET /products products listProducts
//  Return a list of products
//
//  Consumes:
//  - application/json
//
//  Produces:
//  - application/json
//
//  Schemes: http
//
//	Responses:
//	200: productsResponse
func (p Produts) All(wr http.ResponseWriter, r *http.Request) {
	wr.Header().Add("Content-Type", "application/json")
	lp := data.FindAll()
	err := ToJson(wr, lp)
	if err != nil {
		http.Error(wr, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func ToJson(w io.Writer, value interface{}) error {
	e := json.NewEncoder(w)
	return e.Encode(value)
}

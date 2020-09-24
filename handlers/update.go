package handlers

import (
	"encoding/json"
	"example01/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// 	swagger:route PUT /products/{id} products updateProduct
//  Save product in database
//
//  Consumes:
//  - application/json
//
//  Schemes: http
//
//	Responses:
//	200: noContent
func (p Produts) Update(wr http.ResponseWriter, r *http.Request) {

	path := mux.Vars(r)["id"]

	id, err := strconv.Atoi(path)
	if err != nil {
		http.Error(wr, "Invalid URI", http.StatusNotFound)
		return
	}

	var prod data.Product
	e := json.NewDecoder(r.Body)

	err = e.Decode(&prod)
	if err != nil {
		http.Error(wr, "Invalid Json", http.StatusInternalServerError)
		return
	}

	err = prod.Validate()
	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
		return
	}

	err = data.Update(&prod, id)
	if err != nil {
		http.Error(wr, "Error to save Produts", http.StatusInternalServerError)
	}

}

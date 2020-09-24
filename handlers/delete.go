package handlers

import (
	"example01/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// 	swagger:route DELETE /products/{id} products deleteProduct
//  Delete product from the database
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
//	201: noContent

//Delete delete product from database
func (p Produts) Delete(wr http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)["id"]

	id, err := strconv.Atoi(path)
	if err != nil {
		http.Error(wr, "Invalid URI", http.StatusNotFound)
		return
	}

	err = data.Delete(id)
	if err != nil {
		http.Error(wr, "Produto not found", http.StatusNotFound)
		return
	}
}

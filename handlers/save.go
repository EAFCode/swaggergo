package handlers

import (
	"encoding/json"
	"example01/data"
	"net/http"
)

// 	swagger:route POST /products products saveProducts
//  Save product in database
//
//  Consumes:
//  - application/json
//
//  Schemes: http
//
//	Responses:
//	200: noContent
func (p Produts) Save(wr http.ResponseWriter, r *http.Request) {
	var prod data.Product
	e := json.NewDecoder(r.Body)

	err := e.Decode(&prod)
	if err != nil {
		http.Error(wr, "Invalid Json", http.StatusInternalServerError)
		return
	}

	err = prod.Validate()
	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
		return
	}

	err = data.Save(&prod)
	if err != nil {
		http.Error(wr, "Error to save Produts", http.StatusInternalServerError)
	}

}

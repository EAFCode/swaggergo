package handlers

import (
	"example01/data"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (p Produts) GetID(wr http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)["id"]

	id, err := strconv.Atoi(path)
	if err != nil {
		http.Error(wr, "Invalid URI", http.StatusNotFound)
		return
	}

	pr, _ := data.FindById(id)
	if pr == nil {
		http.Error(wr, "Produto not found", http.StatusNotFound)
		return
	}

	fmt.Println(pr)
	ToJson(wr, pr)
}

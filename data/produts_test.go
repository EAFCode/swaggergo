package data

import (
	"testing"
)

func TestCheckValidate(t *testing.T) {
	pr := &Product{
		ID:          7,
		Name:        "Esaldino",
		Description: "Testando",
		Price:       15.0,
		SKU:         "asd-asd-asdasdas",
	}
	err := pr.Validate()

	if err != nil {
		t.Error(err)
	}
}

// func TestGet(t *testing.T) {
// 	prhandle := handlers.NewProduts()
// 	testflight.WithServer(http.HandlerFunc(func()
// 		prhandle.All(), func(r *testflight.Requester) {
// 		response := r.Get("/produts")

// 		assert.Equal(t, 200, response.StatusCode)
// 		assert.Equal(t, "hello, drew", response.Body)
// 	})
// }

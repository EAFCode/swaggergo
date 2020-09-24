//Package Classification API of products
//
//API for manager products
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	produces:
//	- application/json
//
// swagger:meta
package main

import (
	"context"
	"example01/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/nicholasjackson/env"
)

var bindAddress = env.String("BIND_ADDRESS", false, "127.0.0.1:8080", "Bind adress for server")

func main() {
	env.Parse()

	mux := mux.NewRouter()
	ph := handlers.NewProduts()

	groute := mux.Methods("GET").Subrouter()
	groute.HandleFunc("/products", ph.All)
	groute.HandleFunc("/products/{id:[0-9]+}", ph.GetID)

	proute := mux.Methods("POST").Subrouter()
	proute.HandleFunc("/products", ph.Save)

	droute := mux.Methods("DELETE").Subrouter()
	droute.HandleFunc("/products/{id:[0-9]+}", ph.Delete)

	uroute := mux.Methods("PUT").Subrouter()
	uroute.HandleFunc("/products/{id:[0-9]+}", ph.Update)

	opts := middleware.SwaggerUIOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.SwaggerUI(opts, nil)
	groute.Handle("/docs", sh)

	groute.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	server := &http.Server{
		Addr:         *bindAddress,
		Handler:      mux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	//No bloking
	go func() {
		err := server.ListenAndServe()

		if err != nil {
			panic(err.Error())
		}
	}()

	// server.Shutdown()
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan //blocking
	log.Println("Receive terminate, graceful shutdown", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(tc)
}

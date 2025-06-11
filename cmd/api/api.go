package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// define a new type called application that holds a field config of type config.
type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) mount() http.Handler {

	// mux := http.NewServeMux()
	// mux.HandleFunc("GET /v1/healthcheck", app.healthCheck)
	// return mux

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	// set a timeout value on the request that will signal through ctx.Done() that the request has timed out and further processing should be stopped.

	r.Use(middleware.Timeout(60 * time.Second))

	// the code below will group together all routes that start with /v1
	r.Route("/v1", func(r chi.Router) {
		r.Get("/healthcheck", app.healthCheck)
	})
	// r.Get("/", func(w http.ResponseWriter, r *http.Request){
	// 	w.Write([]byte("Welcome"))
	// })

	return r

}

// method for the application struct
// starts the HTTP server
// the app *application is a pointer receiver.
// without the * its a value receiver and that way, Go would nt be able to modify the application instance
// the error comes from an interface, built in
func (app *application) run(mux http.Handler) error {
	//mux is your router that handles incoming HTTP requests, and decides which handler function should process them.
	// creates a new, default router
	// mux := http.NewServeMux()
	// srv is  apointer to the server. We use a pointer so we can pass it around and modify the server later if needed.
	srv := &http.Server{
		Addr:    app.config.addr,
		Handler: mux,
		// if the server takes more than 30 seconds to write a response to the client, it will timeout
		WriteTimeout: time.Second * 30,
		// if the server takes more than 10 seconds to read a request from the client, it will timeout
		ReadTimeout: time.Second * 10,
		IdleTimeout: time.Minute,
	}

	// %s means insert a string here
	log.Printf("Server started on %s", app.config.addr)
	// starts the HTTP server and listens for incoming requests on the configured address.
	return srv.ListenAndServe()
}

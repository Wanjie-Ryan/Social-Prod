package main

import "net/http"

// define a new type called application that holds a field config of type config.
type application struct {
	config config
}

type config struct {
	addr string
}

// method for the application struct
// starts the HTTP server
// the app *application is a pointer receiver.
// without the * its a value receiver and that way, Go would nt be able to modify the application instance
// the rror comes from an interface, built in
func (app *application) run() error {
	//mux is your router that handles incoming HTTP requests, and decides which handler function should process them.
	// creates a new, default router
	mux := http.NewServeMux()
	// srv is  apointer to the server. We use a pointer so we can pass it around and modify the server later if needed.
	srv := &http.Server{
		Addr:    app.config.addr,
		Handler: mux,
	}

	// starts the HTTP server and listens for incoming requests on the configured address.
	return srv.ListenAndServe()
}

package main

import "net/http"

// the function below is a method on the application struct
func (app *application) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

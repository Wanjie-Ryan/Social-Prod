package main

import (
	"log"

	env "github.com/Wanjie-Ryan/Social-Prod/internal"
)

func main() {

	cfg := config{
		// addr: ":8080",
		// provide the name of the key, and also a fallback value
		addr: env.GetString("ADDR", ":8080"),
	}
	app := &application{
		config: cfg,
	}

	mux :=app.mount()
	log.Fatal(app.run(mux))

}

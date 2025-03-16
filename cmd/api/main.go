package main

import (
	"log"

	"github.com/arjkashyap/erlic.ai/internal/env"
)

func main() {

	conf := config{
		addr: env.GetString("ADDR", "localhost:8080"),
	}

	app := &application{
		config: conf,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}

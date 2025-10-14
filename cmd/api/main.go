package main

import "log"

func main() {
	cfg := Config{
		addr: ":8080",
	}
	app := &Application{
		config: cfg,
	}

	log.Printf("Server has started at %s", app.config.addr)
	mux := app.mount()
	log.Fatal(app.run(mux))
}

package main

import (
	"AskharovA/go-social-project/internal/env"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	cfg := Config{
		addr: env.GetString("ADDR", ":8080"),
	}

	app := &Application{
		config: cfg,
	}

	log.Printf("Server has started at %s", app.config.addr)
	mux := app.mount()
	log.Fatal(app.run(mux))
}

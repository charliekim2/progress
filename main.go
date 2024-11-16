package main

import (
	"log"
	"net/http"
	"os"

	"github.com/charliekim2/progress/handlers"
	"github.com/charliekim2/progress/middleware"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./css"))
	mux.Handle("/css/", http.StripPrefix("/css/", fs))

	mux.HandleFunc("/", handlers.Home)

	logMux := middleware.NewLogger(mux)

	addr := os.Getenv("ADDR") + ":" + os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(addr, logMux))
}

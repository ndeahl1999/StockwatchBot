package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	groupme "github.com/ndeahl1999/StockwatchBot/internal/messaging"
)

func pingRespond(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server is running!\n"))
}

func main() {
	//Load required environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Initialize server router
	router := mux.NewRouter()
	router.HandleFunc("/ping", pingRespond).Methods("GET", "OPTIONS")

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Access-Control-Allow-Origin", "Origin"})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	deployPort := os.Getenv("DEPLOY_PORT")

	if deployPort == "" {
		log.Fatal("ERROR: Deploy port not set in environment")
	}

	log.Println("Starting server on port " + deployPort)

	groupme.InitializeBot()

	log.Fatal(http.ListenAndServe(":"+deployPort, handlers.CORS(headers, origins, methods)(router)))

}

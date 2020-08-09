package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	mysqldb "github.com/ndeahl1999/StockwatchBot/internal/database"
	groupme "github.com/ndeahl1999/StockwatchBot/internal/messaging"
)

func pingRespond(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server is running!\n"))
}

//syntax is "SQL_DB_USERNAME:SQL_DB_PASSWORD#@tcp(SQL_DB_HOSTNAME:SQL_DB_PORT)/SQL_DB_NAME"
//TODO: Add checks for each parameter
func generateDBConnectionString() string {
	var dbConnectionString string
	dbUser := os.Getenv("SQL_DB_USERNAME")
	dbPassword := os.Getenv("SQL_DB_PASSWORD")
	dbHostname := os.Getenv("SQL_DB_HOSTNAME")
	dbPort := os.Getenv("SQL_DB_PORT")
	dbName := os.Getenv("SQL_DB_NAME")
	dbConnectionString = dbUser + ":" + dbPassword + "#@tcp(" + dbHostname + ":" + dbPort + ")/" + dbName
	return dbConnectionString
}

func main() {
	//Load required environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbConnectionString := generateDBConnectionString()
	//Initialize databse connection
	mysqldb.DBCon, err = sql.Open("mysql", dbConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer mysqldb.DBCon.Close()

	err = mysqldb.DBCon.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to remote mysql database")
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

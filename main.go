package main

import (
	"database/sql"
	"dcom_service/handler"
	"dcom_service/repository"
	"dcom_service/service"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	var port int

	flag.IntVar(&port, "port", 8080, "API server port")

	flag.Parse()
	logger := log.New(os.Stdout, "DCOM_SERVICE: ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	logger.Printf("Initialize config")
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server := os.Getenv("MSSQL_DB_SERVER")
	dbport := os.Getenv("MSSQL_DB_PORT")
	user := os.Getenv("MSSQL_DB_USER")
	password := os.Getenv("MSSQL_DB_PASSWORD")
	database := os.Getenv("MSSQL_DB_DATABASE")
	// Create connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;trustservercertificate=true;encrypt=DISABLE",
		server, user, password, dbport, database)

	// Create connection pool
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		logger.Fatal("Error creating connection pool: " + err.Error())
	}
	logger.Printf("Database Connected\n")

	// Close the database connection pool after program executes
	defer db.Close()

	rewardRepositoryDB := repository.NewRewardRepositoryDB(db)
	rewardService := service.NewRewardService(rewardRepositoryDB, logger)
	rewardHandler := handler.NewRewardHandler(rewardService)

	router := chi.NewRouter()

	router.Get("/api/rewards", rewardHandler.GetAllReward)
	logger.Printf("API running in port :%d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}

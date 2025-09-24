package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	httpdelivery "github.com/sterlingov/chadder-net-backend/internal/delivery/http"
	"github.com/sterlingov/chadder-net-backend/internal/repository/postgres"
	"github.com/sterlingov/chadder-net-backend/internal/service"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	dbCfg := postgres.Config{
		DSN:             os.Getenv("DB_DSN"),
		MaxOpenConns:    10,
		MaxIdleConns:    5,
		ConnMaxLifetime: 30 * time.Minute,
	}

	db, err := postgres.NewDB(dbCfg)
	if err != nil {
		log.Fatal("DB connection error: ", err)
	}
	defer db.Close()

	userRepo := postgres.NewUserRepo(db)
	userService := service.NewUserService(userRepo)

	routerServices := httpdelivery.Services{User: userService}
	router := httpdelivery.NewRouter(&routerServices)

	addr := os.Getenv("HOST") + ":" + os.Getenv("HTTP_PORT")
	log.Printf("Chadder-Net Backend %s\nServer started at %s", os.Getenv("VERSION"), addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal(err)
	}
}

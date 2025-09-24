package httpdelivery

import (
	"net/http"

	"github.com/sterlingov/chadder-net-backend/internal/service"

	"github.com/gorilla/mux"
)

type Services struct {
	User *service.UserService
}

func NewRouter(services *Services) http.Handler {
	r := mux.NewRouter()

	userHandler := NewUserHandler(services.User)

	r.HandleFunc("/users", userHandler.Register).Methods("POST")
	r.HandleFunc("/users/{id}", userHandler.GetByID).Methods("GET")
	r.HandleFunc("/users/by-username", userHandler.GetByUsername).Methods("GET")

	// позже можно добавлять:
	// r.HandleFunc("/login", userHandler.Login).Methods("POST")

	return r
}

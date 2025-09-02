package httpdelivery

import (
	"net/http"

	"github.com/sterlingov/chadder-net-backend/internal/service"

	"github.com/gorilla/mux"
)

func NewRouter(userService *service.UserService) http.Handler {
	r := mux.NewRouter()

	userHandler := NewUserHandler(userService)

	r.HandleFunc("/register", userHandler.Register).Methods("POST")

	// позже можно добавлять:
	// r.HandleFunc("/login", userHandler.Login).Methods("POST")

	return r
}

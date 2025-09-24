package httpdelivery

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	httpdto "github.com/sterlingov/chadder-net-backend/internal/delivery/http/dto"
	"github.com/sterlingov/chadder-net-backend/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req httpdto.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if _, err := h.service.Register(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetByID(id)
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
	}

	json.NewEncoder(w).Encode(user)
	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) GetByUsername(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	if username == "" {
		http.Error(w, "username parameter is required", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetByUsername(username)
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
	}

	json.NewEncoder(w).Encode(user)
	w.WriteHeader(http.StatusOK)
}

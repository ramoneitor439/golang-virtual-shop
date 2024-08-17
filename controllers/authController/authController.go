package authcontroller

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	authdtos "mystore.com/dtos/authDtos"
	"mystore.com/infrastructure/repositories/usersRepository"
	"mystore.com/services/auth"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	body, bodyErr := io.ReadAll(r.Body)
	if bodyErr != nil {
		http.Error(w, "Request body is in a bad format", http.StatusBadRequest)
		return
	}

	var request authdtos.SignInRequest
	if parseErr := json.Unmarshal(body, &request); parseErr != nil {
		http.Error(w, "Request body is in a bad format", http.StatusBadRequest)
		return
	}

	response, authErr := auth.SignIn(request.Email, request.Password)
	if authErr != nil {
		http.Error(w, authErr.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if jsonErr := json.NewEncoder(w).Encode(*response); jsonErr != nil {
		http.Error(w, jsonErr.Error(), http.StatusInternalServerError)
		return
	}
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	body, bodyErr := io.ReadAll(r.Body)
	if bodyErr != nil {
		http.Error(w, "Request body is in a bad format", http.StatusBadRequest)
		return
	}

	var request authdtos.SignUpRequest
	if parseErr := json.Unmarshal(body, &request); parseErr != nil {
		http.Error(w, "Request body is in a bad format", http.StatusBadRequest)
		return
	}

	err := auth.SignUp(request.Email, request.FirstName, request.LastName, request.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["userId"]
	if userId == "" {
		http.Error(w, "User Id is required", http.StatusBadRequest)
		return
	}

	id, parseErr := strconv.ParseUint(userId, 10, 64)
	if parseErr != nil {
		http.Error(w, "User Id is ina a bad format", http.StatusBadRequest)
		return
	}

	err := usersRepository.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func Update(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["userId"]
	if userId == "" {
		http.Error(w, "User Id is required", http.StatusBadRequest)
		return
	}

	id, parseErr := strconv.ParseUint(userId, 10, 64)
	if parseErr != nil {
		http.Error(w, "User Id is ina a bad format", http.StatusBadRequest)
		return
	}

	body, bodyErr := io.ReadAll(r.Body)
	if bodyErr != nil {
		http.Error(w, "Request body is in a bad format", http.StatusBadRequest)
		return
	}

	var request authdtos.UpdateRequest
	if parseErr := json.Unmarshal(body, &request); parseErr != nil {
		http.Error(w, "Request body is in a bad format", http.StatusBadRequest)
		return
	}

	updateErr := auth.Update(id, request)
	if updateErr != nil {
		http.Error(w, updateErr.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/valentinrb1/go-api-rest.git/middle"
	"github.com/valentinrb1/go-api-rest.git/models"
)

var users []models.User

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	tokenStr := r.Header.Get("Authorization")
	if tokenStr == "" {
		http.Error(w, "Token not found", http.StatusUnauthorized)
		return
	}

	tokenStr = strings.Replace(tokenStr, "Bearer ", "", 1)

	_, err := middle.VerifyToken(tokenStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	users, err := middle.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	users = append(users, user)

	createUserCmd := exec.Command("useradd", user.Username, "-m", "-s", "/bin/bash")
	err = createUserCmd.Run()
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	setPasswordCmd := exec.Command("bash", "-c", fmt.Sprintf("echo %s:%s | chpasswd", user.Username, user.Password))
	err = setPasswordCmd.Run()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	chownCmd := exec.Command("chown", fmt.Sprintf("%s:%s", user.Username, user.Username), filepath.Join("/home", user.Username))
	err = chownCmd.Run()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User created successfully: %s", user.Username)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	authenticated := middle.AuthenticateUser(user.Username, user.Password)
	if !authenticated {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := middle.CreateJWT(user.Username)
	if err != nil {
		http.Error(w, "Error creating JWT", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, token)
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var registeredUser = User{
	Username: "siam",
	Password: "1234",
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	fmt.Printf("Username: %s\n", user.Username)
	fmt.Printf("Password: %s\n", user.Password)

	if user.Username != registeredUser.Username || user.Password != registeredUser.Password {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Login successful")

}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/login", loginHandler)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", mux)
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	Email string `json:"email"`
}

var users = []User{
	{
		Id: 1,
		Name: "Mamun",
		Age: 26,
		Email: "mamun333089@gmail.com",
	},
	{
		Id: 2,
		Name: "Mamun2",
		Age: 25,
		Email: "testing2@gmail.com",
	},
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", rootHandler)
	mux.HandleFunc("GET /health", healthHandler)
	mux.HandleFunc("POST /createUser", createUserHandler)
	mux.HandleFunc("GET /users", getUserHandler)
	mux.HandleFunc("GET /users/{id}", getSingleUserHandler)
	mux.HandleFunc("PUT /users/{id}", updateSingleUserHandler)


	fmt.Println("Server is running at port 5000")
	err := http.ListenAndServe(":5000", mux)

	if err != nil {
		fmt.Println("Server error", err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to go server!")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server is up and healthy")
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid request")
		return
	}
	newUser.Id = len(users) + 1
	users = append(users, newUser)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// users, _ := json.Marshal(users)
	// w.Write(users)
	encoder := json.NewEncoder(w)
	encoder.Encode(users)
}

func getSingleUserHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	// fmt.Printf("the value of id is %v and the type of the id is %T", idParam, idParam)

	id, err := strconv.Atoi(idParam)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid user id")
		return
	}

	for _, user := range users{
		if user.Id == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "User not found")
}

func updateSingleUserHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid user id")
		return
	}

	var updatedUser User

	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid request body")
		return
	}

	for idx, user := range users {
		if user.Id == id {
			updatedUser.Id = id
			users[idx] = updatedUser

			w.Header().Set("Content-type", "application/json")
			json.NewEncoder(w).Encode(updatedUser)
			return
		}
	}

	
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "User not found")
}
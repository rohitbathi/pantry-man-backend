package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	// "github.com/go-sql-driver/mysql"
)

type User struct{
	Username string `json:"username"`
	Password string `json:"password"`
}

func userExists(username string, password string) bool{
	if username=="test1234" && password=="test1234"{
		return true
	}
	return false
}

func rootHandler(response http.ResponseWriter, request *http.Request){
	// fmt.Fprintf(response, "Welcome to pantryman-backend")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(
		map[string]string{"message":"Welcome to pantryman-backend"},
	)
}

func loginHandler(response http.ResponseWriter,request *http.Request){
	fmt.Println("login")
	var user User;
	// reading request body
	body, err := io.ReadAll(request.Body)
	if err!=nil{
		// fmt.Println("error in reading body")
		// panic(fmt.Sprintf("error in login: %v",err))
		http.Error(
			response, 
			"Error reading request body", 
			http.StatusInternalServerError,
		)
		return
	}
	defer request.Body.Close()

	// converting body to json object
	err = json.Unmarshal(body, &user)
	if err!=nil{
		// fmt.Println("error in unmarshalling")
		// panic(err)
		http.Error(
			response,
			"Invalid format, expected JSON",
			http.StatusInternalServerError,
		)
		return
	}

	if userExists(user.Username, user.Password){
		// response.WriteHeader(http.StatusOK)
		// fmt.Println("User found")
		// fmt.Fprintf(response, "User found")
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(
			map[string]string{"message":"User found"},
		)
	}else{
		// response.WriteHeader(http.StatusNotFound)
		// fmt.Println("User not found")
		// fmt.Fprintf(response, "User not found")
		response.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(response).Encode(
			map[string]string{"message":"Invalid username or password"},
		)
	}
}

func main(){

	// setting routes / Router
	http.HandleFunc("GET /", rootHandler)
	http.HandleFunc("POST /login",loginHandler)

	// creating and running Server
	server := http.Server{
		Addr: ":8080",
	}

	err := server.ListenAndServe()
	if err!=nil {
		fmt.Println("error in creating server")
		panic(err)
	}
}
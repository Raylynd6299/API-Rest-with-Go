package main

import (
	"encoding/json"
	"log"
	"net/http"

	"./connect"
	"./structures"
	"github.com/gorilla/mux"
)

//para server necesitamos http
//gorilla mux

func main() {

	connect.InitializaDatabase()
	defer connect.CloseConnectionBD()

	//varibale para manejar las fucniones
	r := mux.NewRouter()

	// primer argumente un URl a manejar, luego que funcion la manejara
	// luego Methods para especificar con que metodo se manejara
	// GET, POST,PUT ...
	r.HandleFunc("/user/{id}", GetUser).Methods("GET")
	r.HandleFunc("/user/new", NewUser).Methods("POST")
	r.HandleFunc("/user/update/{id}", UpdateUser).Methods("PATCH")
	r.HandleFunc("/user/delete/{id}", DeleteUser).Methods("DELETE")

	log.Println("El servidor se encuentra en el puerto 8001")
	log.Fatal(http.ListenAndServe(":8001", r))

}

//GetUser Esta funcion es para manejar la rute /user/
func GetUser(w http.ResponseWriter, r *http.Request) {
	//obtenemos las variables pasadas por url
	vars := mux.Vars(r)
	userID := vars["id"]

	status := "success"
	var message string
	user := connect.GetUser(userID)

	if user.ID <= 0 {
		status = "error"
		message = "User not found"
	}

	response := structures.Response{
		Status:  status,
		Data:    user,
		Message: message,
	}

	//para manejo de BD usaremos gorm

	//w.Write([]byte("Gorilla_REST!\n"))
	// user := structures.User{
	// 	Username:  "Raylynd6299",
	// 	FirstName: "Raymundo",
	// 	LastName:  "Pulido",
	// }

	//Version 1
	json.NewEncoder(w).Encode(response)

	//version 2
	jsso, _ := json.Marshal(user)
	w.Write(jsso)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	var user structures.User
	connect.DeleteUser(userID)

	response := structures.Response{Status: "Succes", Data: user, Message: ""}
	json.NewEncoder(w).Encode(response)
}
}
func NewUser(w http.ResponseWriter, r *http.Request) {
	user := GetUserRequest(r)
	response := structures.Response{Status: "success", Data: connect.CreateUser(user), Message: ""}

	json.NewEncoder(w).Encode(response)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	user := GetUserRequest(r)
	response := structures.Response{Status: "success", Data: connect.UpdateUser(userID, user), Message: ""}

	json.NewEncoder(w).Encode(response)
}

func GetUserRequest(r *http.Request) structures.User {
	var user structures.User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	return user
}

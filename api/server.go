package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ResponseError struct {
	Message string `json:"message"`
}

func (e ResponseError) Error() string {
	return e.Message
}

func NewResponseError(message string) *ResponseError {
	m := ResponseError{message}
	return &m
}

type Greet struct {
	Text string `json:"text"`
}

func NewGreet(text string) *Greet {
	greet := Greet{Text: text}
	return &greet
}

// Greets the user on root route
func Greeting(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("Root api got called %s\n", req.RequestURI)
	data := NewGreet("hello")
	json.NewEncoder(w).Encode(data)
}

type Person struct {
	Name string `json:"name"`
	Age  int16  `json:"age"`
}

func NewPerson(name string, age int16) *Person {
	person := Person{name, age}
	return &person
}

func PersonHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("Api got called %s with %s method\n", req.RequestURI, req.Method)
	w.Header().Set("Content-Type", "application/json")
	if req.Method == http.MethodPost {
		creatPerson(w, req)
		return
	} else if req.Method == http.MethodGet {
		fetchPeople(w, req)
		return
	}
	e := NewResponseError("You are calling with a wrong method")
	http.Error(w, e.Error(), http.StatusBadRequest)
}

func creatPerson(w http.ResponseWriter, req *http.Request) {
	var p Person
	err := json.NewDecoder(req.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//fmt.Fprintf(w, "Person: %+v", p)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

func fetchPeople(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("Api got called %s\n", req.RequestURI)
	people := [2]Person{{"amir", 32}, {"Alireza", 34}}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(people)
}

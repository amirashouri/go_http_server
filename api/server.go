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

func CreatPerson(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("Api got called %s\n", req.RequestURI)
	w.Header().Set("Content-Type", "application/json")

	if req.Method != http.MethodPost {
		e := NewResponseError("You are calling with a wrong method")
		http.Error(w, e.Error(), http.StatusBadRequest)
		return
	}
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

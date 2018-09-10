package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
)

//Book Strut (Module)
type Book struct {
	ID  string `json:"id"`
	Isbn  string `json:"isbn"`
	Title  string `json:"title"`
	Author  *Author `json:"author"`
}
//Auhtor struct
type Author struct {
	Firstname  string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//Init var Books as a slice Book Struct
var books []Book


//Get All Books
func getBooks(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r) //Get params
	//loop through books and find with id
	for _ , item := range books{
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}
func createBook(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	var book Book
	_ : json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000)) //Mock id - not safe
	books = append(books,book)
	fmt.Println(book)
	json.NewEncoder(w).Encode(book)
}
func updateBook(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index ,item :=range books{
		if item.ID ==params["id"] {
			books = append(books[:index],books[index+1:]...)
			var book Book
			_ : json.NewDecoder(r.Body).Decode(&book)
			book.ID = strconv.Itoa(rand.Intn(1000000)) //Mock id - not safe
			books = append(books,book)
			fmt.Println(book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}
func deleteBook(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index ,item :=range books{
		if item.ID ==params["id"] {
			books = append(books[:index],books[index+1:]...)
		}
	}
	json.NewEncoder(w).Encode(books)
}



func main() {
	//init Router
	r := mux.NewRouter()

	//Mock Data  --todo - impletent DB
	books = append(books,Book{ID:"1",Isbn:"34343",Title:"Books one",Author:&Author{Firstname:"liu",Lastname:"yunlong"}})

	books = append(books,Book{ID:"2",Isbn:"56565",Title:"Books two",Author:&Author{Firstname:"liu",Lastname:"tiantian"}})

	//Router Handlers  / Endpoints
	r.HandleFunc("/api/books",getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}",getBook).Methods("get")
	r.HandleFunc("/api/books",createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}",updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}",deleteBook).Methods("DELETE")
	http.ListenAndServe(":8000",r)
	//log.Fatal(http.ListenAndServe(":8000",r))
}

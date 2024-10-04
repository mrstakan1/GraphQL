package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	. "prac7/graphql"
)

func main() {
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		var requestBody struct {
			Query string `json:"query"`
		}

		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
			return
		}

		result := ExecuteQuery(requestBody.Query, Schema)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Сервер запущен на http://localhost:8080/graphql")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//Test queries for postman (body -> raw json)
//Add Book
//{
// "query": "{ books { id name genre author } }"
//}

//Show Books
//{
//"query": "mutation { addBook(name: \"War and Peace\", genre: \"Historical Fiction\", author: \"Lev Tolstoy\") { id name genre author } }"
//}

package main

import (
	"log"
	"net/http"

	"github.com/NaomiShinna/test-fe-be/tree/main/api-article/controllers/articlecontroller"
	"github.com/NaomiShinna/test-fe-be/tree/main/api-article/models"

	"github.com/gorilla/mux"
)

func main() {

	models.ConnectDatabase()
	router := mux.NewRouter()

	router.HandleFunc("/api/getArticles", articlecontroller.GetAllArticle).Methods("GET")

	router.HandleFunc("/api/article/{limit}/{offset}", articlecontroller.GetArticleByLimitOffset).Methods("GET")
	router.HandleFunc("/api/article/{id}", articlecontroller.GetArticleById).Methods("GET")
	router.HandleFunc("/api/article", articlecontroller.CreateArticle).Methods("POST")
	router.HandleFunc("/api/article/{id}", articlecontroller.UpdateArticleById).Methods("PUT")
	router.HandleFunc("/api/deleteArticle/{id}", articlecontroller.DeleteArticleById).Methods("PUT")

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		if r.Method == "OPTIONS" {
			w.Write([]byte("allowed"))
			return
		}

		w.Write([]byte("hello"))
	})

	log.Fatal(http.ListenAndServe(":8080", router))

}

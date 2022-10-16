package main

import (
	"log"
	"net/http"

	"github.com/NaomiShinna/test-fe-be/api-article/controllers/articlecontroller"
	"github.com/NaomiShinna/test-fe-be/api-article/models"

	"github.com/gorilla/mux"
)

func main() {
	models.ConnectDatabase()
	router := mux.NewRouter()

	router.HandleFunc("/api/getArticles", articlecontroller.GetAllArticle).Methods("GET")
	router.HandleFunc("/api/article/{id}", articlecontroller.GetArticleById).Methods("GET")
	router.HandleFunc("/api/article", articlecontroller.CreateArticle).Methods("POST")
	router.HandleFunc("/api/article/{id}", articlecontroller.UpdateArticleById).Methods("PUT")
	router.HandleFunc("/api/article/{id}", articlecontroller.DeleteArticleById).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}

// var db *gorm.DB
// var err error

// type PostArticle struct {
// 	Id           int       `form:"id" json:"id"`
// 	Title        string    `form:"title" json:"title"`
// 	Content      string    `form:"content" json:"content"`
// 	Category     string    `form:"category" json:"category"`
// 	Created_date time.Time `form:"created_date" json:"created_date"`
// 	Updated_date time.Time `form:"updated_date" json:"updated_date"`
// 	Status       string    `form:"status" json:"status"`
// }

// type Result struct {
// 	Code    int         `json:"code"`
// 	Data    interface{} `json:"data"`
// 	Message string      `json:"message"`
// }

// // import
// // "database/sql"
// // _ "github.com/go-sql-driver/mysql"

// // func GetConnection() *sql.DB {
// // 	db, err := sql.Open("mysql", "root:mysql@tcp(localhost:3306)/article")
// // 	if err != nil {
// // 		panic(err)
// // 	}

// // 	// test cek database
// // 	db.SetMaxIdleConns(10)
// // 	db.SetMaxOpenConns(100)
// // 	db.SetConnMaxIdleTime(5 * time.Minute)
// // 	db.SetConnMaxLifetime(60 & time.Minute)

// // 	return db
// // }

// func main() {
// 	db, err := gorm.Open("mysql", "root:@/article?charset=utf8mb4&parseTime=true")
// 	if err != nil {
// 		panic(err)
// 	} else {
// 		log.Println("Connection established")
// 	}

// 	db.AutoMigrate(&PostArticle{})

// 	handleRequests()
// }

// func handleRequests() {
// 	log.Println("Start the development server at http://127.0.0.1:9999")

// 	myRouter := mux.NewRouter().StrictSlash(true)

// 	myRouter.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusNotFound)

// 		res := Result{Code: 404, Message: "Method not found"}
// 		response, _ := json.Marshal(res)
// 		w.Write(response)
// 	})

// 	myRouter.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusMethodNotAllowed)

// 		res := Result{Code: 403, Message: "Method not allowed"}
// 		response, _ := json.Marshal(res)
// 		w.Write(response)
// 	})

// 	myRouter.HandleFunc("/", homePage)
// 	myRouter.HandleFunc("/api/postArticle", createArticle).Methods("POST")
// 	myRouter.HandleFunc("/api/getArticles", getArticles).Methods("GET")
// 	myRouter.HandleFunc("/api/getArticleById/{id}", getArticle).Methods("GET")
// 	myRouter.HandleFunc("/api/updateArticle/{id}", updateArticle).Methods("PUT")
// 	myRouter.HandleFunc("/api/deleteArticle/{id}", deleteArticle).Methods("DELETE")

// 	log.Fatal(http.ListenAndServe(":9999", myRouter))
// }

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Welcome to api")
// }

// func createArticle(w http.ResponseWriter, r *http.Request) {
// 	payloads, _ := ioutil.ReadAll(r.Body)

// 	var article PostArticle
// 	json.Unmarshal(payloads, &article)

// 	// now := time.Now()

// 	// db.Created_date = now
// 	// db.Updated_date = now
// 	db.Create(&article)

// 	res := Result{Code: 200, Data: article, Message: "Success create article"}
// 	result, err := json.Marshal(res)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(result)
// }

// func getArticles(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint hit: get products")

// 	article := []PostArticle{}
// 	db.Find(&article)

// 	res := Result{Code: 200, Data: article, Message: "Success get Articles"}
// 	results, err := json.Marshal(res)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(results)
// }

// func getArticle(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	articleId := vars["id"]

// 	var article PostArticle

// 	db.First(&article, articleId)

// 	res := Result{Code: 200, Data: article, Message: "Success get Article"}
// 	result, err := json.Marshal(res)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(result)
// }

// func updateArticle(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	articleId := vars["id"]

// 	payloads, _ := ioutil.ReadAll(r.Body)

// 	var articleUpdates PostArticle
// 	json.Unmarshal(payloads, &articleUpdates)

// 	var article PostArticle
// 	db.First(&article, articleId)
// 	db.Model(&article).Updates(articleUpdates)

// 	res := Result{Code: 200, Data: article, Message: "Success update article"}
// 	result, err := json.Marshal(res)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(result)
// }

// func deleteArticle(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	articleId := vars["id"]

// 	var article PostArticle

// 	db.First(&article, articleId)
// 	db.Delete(&article)

// 	res := Result{Code: 200, Message: "Success delete article"}
// 	result, err := json.Marshal(res)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(result)
// }

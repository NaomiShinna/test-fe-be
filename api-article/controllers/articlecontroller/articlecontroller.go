package articlecontroller

import (
	"encoding/json"
	"net/http"

	"github.com/NaomiShinna/test-fe-be/tree/main/api-article/helper"

	"gorm.io/gorm"

	"github.com/NaomiShinna/test-fe-be/tree/main/api-article/models"

	"fmt"

	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
)

var ResponseJson = helper.ResponseJson
var ResponseError = helper.ResponseError

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to API!")
}

func GetAllArticle(w http.ResponseWriter, r *http.Request) {
	var articles []models.Article

	if err := models.DB.Find(&articles).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseJson(w, http.StatusOK, articles)
}

func GetArticleByLimitOffset(w http.ResponseWriter, r *http.Request) {
	var articles []models.Article
	vars := mux.Vars(r)

	batas, _ := strconv.Atoi(vars["limit"])
	halaman, _ := strconv.Atoi(vars["offset"])

	if err := models.DB.Limit(batas).Offset(halaman * batas).Find(&articles).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseJson(w, http.StatusOK, articles)
}

func GetArticleById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var article models.Article
	if err := models.DB.First(&article, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseError(w, http.StatusNotFound, "Article tidak ditemukan")
			return
		default:
			ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	ResponseJson(w, http.StatusOK, article)
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {

	var article models.Article

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&article); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	if err := models.DB.Create(&article).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseJson(w, http.StatusCreated, article)

}

func UpdateArticleById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var article models.Article

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&article); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	if models.DB.Where("id = ?", id).Updates(&article).RowsAffected == 0 {
		ResponseError(w, http.StatusBadRequest, "Tidak dapat mengupdate article")
		return
	}

	article.Id = id

	ResponseJson(w, http.StatusOK, article)

}

func DeleteArticleById(w http.ResponseWriter, r *http.Request) {

	// vars := mux.Vars(r)
	// id, err := strconv.ParseInt(vars["id"], 10, 64)
	// if err != nil {
	// 	ResponseError(w, http.StatusBadRequest, err.Error())
	// 	return
	// }

	var article models.Article
	// if err := models.DB.Where("id = ?", id).Updates(&article).Error; err != nil {
	// 	switch err {
	// 	case gorm.ErrRecordNotFound:
	// 		ResponseError(w, http.StatusBadRequest, "Tidak dapat menghapus article")
	// 		return
	// 	default:
	// 		ResponseError(w, http.StatusInternalServerError, err.Error())
	// 		return
	// 	}
	// }

	// models.DB.Update("status", "Thrash").Save(&article)

	// response := map[string]string{"message": "Article berhasil dihapus"}
	// ResponseJson(w, http.StatusOK, response)

	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/article?parseTime=true"))
	if err != nil {
		panic(err)
	}
	// db.AutoMigrate(&article{})

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	db.Model(&article).Where("id = ?", id).Update("status", "Thrash")
	response := map[string]string{"message": "Article berhasil dihapus"}
	ResponseJson(w, http.StatusOK, response)
}

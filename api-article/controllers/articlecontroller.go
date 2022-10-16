package articlecontroller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/NaomiShinna/api-article/helper"

	"gorm.io/gorm"

	"api-article/models"

	"github.com/gorilla/mux"
)

var ResponseJson = helper.ResponseJson
var ResponseError = helper.ResponseError

func GetAllArticlendex(w http.ResponseWriter, r *http.Request) {
	var articles []models.Article

	if err := models.DB.Find(&articles).Error; err != nil {
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

	input := map[string]string{"id": ""}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	var article models.Article
	if models.DB.Delete(&article, input["id"]).RowsAffected == 0 {
		ResponseError(w, http.StatusBadRequest, "Tidak dapat menghapus product")
		return
	}

	response := map[string]string{"message": "Product berhasil dihapus"}
	ResponseJson(w, http.StatusOK, response)
}

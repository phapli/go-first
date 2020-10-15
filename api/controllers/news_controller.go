package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/phapli/go-kit/api/models"
	"github.com/phapli/go-kit/api/responses"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (server *Server) CreateNews(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	news := models.News{}
	err = json.Unmarshal(body, &news)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	userCreated, err := news.SaveNews(server.DB)

	if err != nil {

		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))
	responses.JSON(w, http.StatusCreated, userCreated)
}

func (server *Server) GetAllNews(w http.ResponseWriter, _ *http.Request) {
	newsModel := models.News{}

	news, err := newsModel.FindAllNews(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, news)
}

func (server *Server) GetANews(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	newsModel := models.News{}

	news, err := newsModel.FindByID(server.DB, uint(id))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, news)
}

func (server *Server) UpdateNews(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Check if the news id is valid
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Check if the news exist
	newsModel := models.News{}
	err = server.DB.Debug().Model(models.News{}).Where("id = ?", pid).Take(&newsModel).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("News not found"))
		return
	}

	// Read the data posted
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Start processing the request data
	newsUpdate := models.News{}
	err = json.Unmarshal(body, &newsUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	newsUpdate.ID = newsModel.ID //this is important to tell the model the newsModel id to update, the other update field are set above

	postUpdated, err := newsUpdate.UpdateANews(server.DB)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, postUpdated)
}

func (server *Server) DeleteNews(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Is a valid newsModel id given to us?
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Check if the newsModel exist
	newsModel := models.News{}
	err = server.DB.Debug().Model(models.News{}).Where("id = ?", pid).Take(&newsModel).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	_, err = newsModel.DeleteANews(server.DB, uint(pid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	responses.JSON(w, http.StatusNoContent, "")
}

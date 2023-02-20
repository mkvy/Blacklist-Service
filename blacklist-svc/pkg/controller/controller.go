package controller

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/dto"
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/service"
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/utils"
	"log"
	"net/http"
)

type Controller struct {
	service service.UserBlackListSvc
}

func NewController(service service.UserBlackListSvc) *Controller {
	return &Controller{service}
}

func (c *Controller) AddHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("AddHandler: handling method POST api/v1/test/")
	var data dto.BlacklistRequestDto
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	validate := validator.New()
	err = validate.Struct(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("AddHandler send to service")
	id, err := c.service.Add(data)
	if err != nil {
		log.Println("AddHandler err in service", err)
		if err == utils.ErrAlreadyExists {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(map[string]string{"id": id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *Controller) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	log.Println("DeleteHandler: delete request income with id: ", id)
	err := c.service.Delete(id)
	if err != nil {
		log.Println("DeleteHandler Error while deleting: ", err)
		if err == utils.ErrNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("success delete")
	w.WriteHeader(http.StatusNoContent)
}

func (c *Controller) GetByUsernameHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("user_name")
	log.Println("GetByUsernameHandler call to service with username: ", username)
	reqdata, err := c.service.GetByUsername(username)
	if err != nil {
		log.Println("error retrieving data")
		if err == utils.ErrNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(reqdata); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *Controller) GetByPhoneHandler(w http.ResponseWriter, r *http.Request) {
	phone_number := r.FormValue("phone_number")
	log.Println("GetByPhoneHandler call to service with phone_number: ", phone_number)
	reqdata, err := c.service.GetByPhoneNumber(phone_number)
	if err != nil {
		log.Println("GetByPhoneHandler error retrieving data")
		if err == utils.ErrNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(reqdata); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

package controller

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/dto"
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/service"
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
	log.Println("handling method api/v1/test/")
	log.Println("handling method post")
	var data dto.BlacklistRequestDto
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println(err.Error())
		log.Println(data)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		//todo
		return
	}
	validate := validator.New()
	err = validate.Struct(data)
	log.Println("after validate")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("send to service")
	id, err := c.service.Add(data)
	if err != nil {
		log.Println("err  in service", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, id)
}

func (c *Controller) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	err := c.service.Delete(id)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("success delete")
}

func (c *Controller) GetByUsernameHandler(w http.ResponseWriter, r *http.Request) {
	var data dto.GetByUsernameDto
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	validate := validator.New()
	err = validate.Struct(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("GetByUsernameHandler call to service")
	reqdata, err := c.service.GetByUsername(data)
	if err != nil {
		log.Println("error retrieving data")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for _, v := range reqdata {
		fmt.Println(v)
	}
	log.Println("end")
}

func (c *Controller) GetByPhoneHandler(w http.ResponseWriter, r *http.Request) {
	var data dto.GetByPhoneDto
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	validate := validator.New()
	err = validate.Struct(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("GetByPhoneHandler call to service")
	reqdata, err := c.service.GetByPhoneNumber(data)
	if err != nil {
		log.Println("error retrieving data")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for _, v := range reqdata {
		fmt.Println(v)
	}
	log.Println("end")
}

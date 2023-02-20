package controller

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/dto"
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/service"
	"net/http"
)

type Controller struct {
	service service.UserBlackListSvc
}

func NewController(service service.UserBlackListSvc) *Controller {
	return &Controller{service}
}

func (c *Controller) AddHandler(w http.ResponseWriter, r *http.Request) {
	var data dto.BlacklistRequestDto
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		//todo
		return
	}
	validate := validator.New()
	err = validate.Struct(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, err := c.service.Add(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(id))
}

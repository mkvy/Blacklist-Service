package controller

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/mkvy/BlacklistTestTask/internal/dto"
	"github.com/mkvy/BlacklistTestTask/internal/service"
	"github.com/mkvy/BlacklistTestTask/pkg/utils"
	"log"
	"net/http"
)

type Controller struct {
	service service.UserBlackListSvc
}

func NewController(service service.UserBlackListSvc) *Controller {
	return &Controller{service}
}

// Create godoc
// @Summary Create user in blacklist
// @Description Добавление пользователя в черный список
// @Tags blacklist
// @Success 201 {object} dto.ResponseId
// @Router /blacklist/ [post]
// @Param reqeust body dto.BlacklistRequestDto true "Create body"
// @Failure      400
// @Failure      401
// @Failure      404
// @Failure      500
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
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
	resp := &dto.ResponseId{Id: id}
	err = json.NewEncoder(w).Encode(&resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Delete godoc
// @Summary Get black list by user's phone
// @Description Удаление пользователя из черного списка
// @Tags blacklist
// @Success 204
// @Router /blacklist/{id} [delete]
// @Param id path string true "Delete by id"
// @Failure      401
// @Failure      404
// @Failure      500
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
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

// Get godoc
// @Summary Get black list by username or phone_number
// @Description Поиск пользователя в черном списке по user_name или phone_number (требуется только один параметр)
// @Tags blacklist
// @Produce json
// @Success 200 {array} models.BlacklistData
// @Router /blacklist [get]
// @Param user_name query string false "Get by username only"
// @Param phone_number query string false "Get by phone number only"
// @Failure      401
// @Failure      404
// @Failure      500
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
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

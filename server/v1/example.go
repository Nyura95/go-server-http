package v1

import (
	"encoding/json"
	"go-server-http/models"
	"go-server-http/server/response"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type payloadUser struct {
	User models.User
}

type payloadUsers struct {
	Users []*models.User
}

// FindUserByID [GET] get a user
func FindUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		json.NewEncoder(w).Encode(response.BasicResponse(new(payloadUser), "Error parsing id", -1))
	}
	user := models.User{IDUser: id}
	if err = user.FindOne(); err != nil {
		json.NewEncoder(w).Encode(response.BasicResponse(new(payloadUser), "user not found", -1))
	}
	json.NewEncoder(w).Encode(response.BasicResponse(payloadUser{User: user}, "ok", 1))
}

// FindAllActif user
func FindAllActif(w http.ResponseWriter, r *http.Request) {
	users, err := models.FindAllActif()
	if err != nil {
		json.NewEncoder(w).Encode(response.BasicResponse(new(payloadUser), "Error server", -4))
	}
	json.NewEncoder(w).Encode(response.BasicResponse(payloadUsers{Users: users}, "ok", 1))
}

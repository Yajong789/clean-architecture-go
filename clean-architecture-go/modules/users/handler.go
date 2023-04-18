package users

import (
	"encoding/json"
	"net/http"
	"time"
)

type Handler struct {
	Usecase Usecase
}

var APPLICATION_NAME = "yaza barudak"
var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNATURE_KEY = []byte("opewfjdi3f84f339fu3")

func (handler Handler) Login(w http.ResponseWriter, r *http.Request) {
	var userInput User
	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "failed to decode json"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}
	defer r.Body.Close()

	signedToken, err := handler.Usecase.Login(userInput.Username, userInput.Password)
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	response := map[string]string{"token": signedToken}
	json.NewEncoder(w).Encode(response)

}

func (handler Handler) Register(w http.ResponseWriter, r *http.Request) {

	// mengambil inputan dari json
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "failed to decode json"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}
	defer r.Body.Close()

	// hass pass menggunakan bcrypt
	err := handler.Usecase.Register(&user)
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	response := map[string]string{"message": "user data added successfully"}
	json.NewEncoder(w).Encode(response)
}

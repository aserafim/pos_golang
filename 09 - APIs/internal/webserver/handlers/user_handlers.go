package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/aserafim/pos_golang/09_APIs/internal/database"
	"github.com/aserafim/pos_golang/09_APIs/internal/dto"
	"github.com/aserafim/pos_golang/09_APIs/internal/entity"
	"github.com/go-chi/jwtauth"
)

type UserHandler struct {
	UserDB       database.UserInterface
	Jwt          *jwtauth.JWTAuth
	JwtExpiresIn int
}

func NewUserHandler(db database.UserInterface, jwt *jwtauth.JWTAuth, expiresIn int) *UserHandler {
	return &UserHandler{
		UserDB:       db,
		Jwt:          jwt,
		JwtExpiresIn: expiresIn,
	}
}

func (h *UserHandler) GetJwt(w http.ResponseWriter, r *http.Request) {
	var getJwt dto.GetJwt
	err := json.NewDecoder(r.Body).Decode(&getJwt)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userDB, err := h.UserDB.FindByEmail(getJwt.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if !userDB.ValidatePassword(getJwt.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// Cria um map de string para interface
	// para nao se preocupar com os tipos
	_, tokenString, err := h.Jwt.Encode(map[string]interface{}{
		"sub": userDB.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JwtExpiresIn)).Unix(),
	})
	if err != nil {
		panic(err)
	}

	// Cria uma struct que sera retornada
	accessToken := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: tokenString,
	}
	print(accessToken.AccessToken)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)

}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var userInput dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := entity.NewUser(userInput.Name, userInput.Email, userInput.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.UserDB.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)

}

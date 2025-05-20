package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/ShubhamkumarAnand/melkey-go/mel_project/internal/store"
	"github.com/ShubhamkumarAnand/melkey-go/mel_project/internal/tokens"
	"github.com/ShubhamkumarAnand/melkey-go/mel_project/internal/utils"
)

type TokenHandler struct {
	tokenStore store.TokenStore
	userStore  store.UserStore
	logger     *log.Logger
}

type createTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewTokenHandler(tokenStore store.TokenStore, userStore store.UserStore, logger *log.Logger) *TokenHandler {
	return &TokenHandler{
		tokenStore: tokenStore,
		userStore:  userStore,
		logger:     logger,
	}
}

func (h *TokenHandler) HandleCreateToken(w http.ResponseWriter, r *http.Request) {
	var req createTokenRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		h.logger.Printf("Error: createTokenRequest - %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"Error": "Invalid Request Payload"})
		return
	}

	user, err := h.userStore.GetUserByUsername(req.Username)

	if err != nil || user == nil {
		h.logger.Printf("ERROR: GetUserByUsername - %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"Error": "internal server error"})
		return
	}

	passwordDoMatch, err := user.PasswordHash.Matches(req.Password)

	if err != nil {
		h.logger.Printf("Error: PasswordHash.Matches - %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"Error": " internal server error"})
		return
	}

	if !passwordDoMatch {
		utils.WriteJSON(w, http.StatusUnauthorized, utils.Envelope{"Error": "Invalid Credentials"})
		return
	}

	token, err := h.tokenStore.CreateNewToken(user.ID, 24*time.Hour, tokens.ScopeAuth)

	if err != nil {
		h.logger.Printf("Error: Creating Token %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"Error": " internal server error"})
		return
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"auth_token": token})
}

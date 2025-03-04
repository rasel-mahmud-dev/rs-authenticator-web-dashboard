package mfaSecurity

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"rs/auth/app/dto"
	"rs/auth/app/models"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/repositories"
	"rs/auth/app/response"
	"rs/auth/app/utils"
	"time"
)

func Finalize2FASecret(w http.ResponseWriter, r *http.Request) {

	authSession := (*r).Context().Value("authSession").(*models.AuthSession)

	if authSession == nil {
		response.Respond(w, statusCode.UNAUTHORIZED, "UNAUTHORIZED", nil)
		return
	}

	var body dto.Completed2FASecretBody
	err := json.NewDecoder((*r).Body).Decode(&body)
	if err != nil {
		response.Respond(w, statusCode.INVALID_JSON_FORMAT, "Invalid JSON format", nil)
		return
	}

	el := models.MfaSecurityToken{
		UserID:    authSession.UserId,
		CodeName:  &body.CodeName,
		Secret:    body.Secret,
		QrCodeURL: &body.QrCodeURL,
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		AppName:   "RsAuth",
	}

	token, err := repositories.MfaSecurityTokenRepo.InsertMfaSecurityToken(el)
	if err != nil {
		utils.LoggerInstance.Error(err.Error())
		response.Respond(w, statusCode.INTERNAL_SERVER_ERROR, "Unable to completed authenticator app setup", nil)
		return
	}

	utils.LoggerInstance.Info("Successfully completed authenticator app setup", token)
	response.Respond(w, statusCode.OK, "OK", token)
}

func GetAllConnectedAuthenticatorApps(w http.ResponseWriter, r *http.Request) {

	authSession := (*r).Context().Value("authSession").(*models.AuthSession)

	if authSession == nil {
		response.Respond(w, statusCode.UNAUTHORIZED, "UNAUTHORIZED", nil)
		return
	}

	items, err := repositories.MfaSecurityTokenRepo.GetAllItems(authSession.UserId)
	if err != nil {
		utils.LoggerInstance.Error(err.Error())
		response.Respond(w, statusCode.INVALID_JSON_FORMAT, "Fetch all token failed", nil)
		return
	}

	response.Respond(w, statusCode.OK, "OK", items)
}

func RemoveAuthenticator(w http.ResponseWriter, r *http.Request) {
	authSession := (*r).Context().Value("authSession").(*models.AuthSession)
	if authSession == nil {
		response.Respond(w, statusCode.UNAUTHORIZED, "UNAUTHORIZED", nil)
		return
	}
	id := mux.Vars(r)["id"]
	repositories.MfaSecurityTokenRepo.RemoveAuthenticator(authSession.UserId, id)
	response.Respond(w, statusCode.OK, "OK", nil)
}

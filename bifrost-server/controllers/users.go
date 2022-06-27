package controllers

import (
	"context"
	conn "github.com/aloks98/bifrost/bifrost-server/db"
	"github.com/aloks98/bifrost/bifrost-server/models"
	"github.com/aloks98/bifrost/bifrost-server/utils"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

func CreateNewUser(Username, Password, Email string) models.Response {
	_, emailInvalidErr := utils.ValidateEmail(Email)
	if emailInvalidErr != nil {
		return utils.CreateResponse(http.StatusBadRequest, emailInvalidErr, nil)
	}

	isUsernameValid, usernameInvalidErr := utils.ValidateUsername(strings.ToLower(Username))
	if usernameInvalidErr != nil {
		return utils.CreateInternalServerErrorResponse("")
	}
	if !isUsernameValid {
		log.Error("Username doesn't match required criteria.")
		return utils.CreateResponse(http.StatusBadRequest, "Username doesn't match required criteria.", nil)
	}

	isPasswordValid, passwordInvalidErr := utils.ValidatePassword(Password)
	if passwordInvalidErr != nil {
		return utils.CreateInternalServerErrorResponse("")
	}
	if !isPasswordValid {
		log.Error("Password doesn't match required criteria.")
		return utils.CreateResponse(http.StatusBadRequest, "Password doesn't match required criteria.", nil)
	}

	db := conn.CreateConn()
	defer db.Close()

	ExistingUser := ""
	ExistingEmail := ""
	existingUserErr := db.QueryRow(context.Background(), utils.SelectUsernameSQL, strings.ToLower(Username)).Scan(&ExistingUser)
	if existingUserErr != nil && existingUserErr != pgx.ErrNoRows {
		log.Error("Fetching Username failed with error: ", existingUserErr)
		return utils.CreateInternalServerErrorResponse("")
	}
	existingEmailErr := db.QueryRow(context.Background(), utils.SelectEmailSQL, Email).Scan(&ExistingEmail)
	if existingEmailErr != nil && existingEmailErr != pgx.ErrNoRows {
		log.Error("Fetching Email failed with error: ", existingEmailErr)
		return utils.CreateInternalServerErrorResponse("")
	}
	if ExistingUser != "" {
		log.Error("Username ", Username, " already exists.")
		return utils.CreateResponse(http.StatusConflict, "Username already exists.", nil)
	}
	if ExistingEmail != "" {
		log.Error("Email ", Email, " already exists.")
		return utils.CreateResponse(http.StatusConflict, "Email already exists.", nil)

	}

	hashedPassword, hashingErr := utils.HashPassword(Password)
	if hashingErr != nil {
		log.Error("Cannot hash password: ", hashingErr)
		return utils.CreateInternalServerErrorResponse("")
	}
	NewUser := models.User{
		Id:             uuid.NewString(),
		Username:       strings.ToLower(Username),
		Password:       hashedPassword,
		Email:          Email,
		AccountEnabled: true,
		EmailVerified:  false,
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
	}
	tag, err := db.Exec(
		context.Background(),
		utils.InsertNewUserSQL,
		NewUser.Id,
		NewUser.Username,
		NewUser.Password,
		NewUser.Email,
		NewUser.AccountEnabled,
		NewUser.EmailVerified,
		NewUser.CreatedAt,
		NewUser.UpdatedAt)
	if err != nil {
		log.Error("Error while creating new user: ", err)
		return utils.CreateInternalServerErrorResponse("")
	}
	log.Info("CREATE: User successful: ", tag)
	return utils.CreateResponse(http.StatusCreated, nil, nil)
}

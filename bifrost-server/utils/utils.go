package utils

import (
	"errors"
	emailverifier "github.com/AfterShip/email-verifier"
	"github.com/aloks98/bifrost/bifrost-server/models"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"regexp"
)

func CreateResponse(status int, err interface{}, data interface{}) models.Response {
	return models.Response{
		Status: status,
		Error:  err,
		Data:   data,
	}
}

func CreateInternalServerErrorResponse(errorString string) models.Response {
	if errorString == "" {
		errorString = "Something bad happened."
	}
	return models.Response{
		Status: http.StatusInternalServerError,
		Error:  errorString,
		Data:   nil,
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func ComparePassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func ValidateEmail(email string) (bool, error) {
	verifier := emailverifier.NewVerifier()
	emailData, err := verifier.Verify(email)
	if err != nil {
		log.Error("Email verification failed with error: ", err)
		return false, err
	}
	if emailData.Disposable {
		freeErr := errors.New("cannot use disposable email")
		log.Error("Email verification failed with error: ", freeErr)
		return false, freeErr
	}
	return true, nil
}

func ValidateUsername(username string) (bool, error) {
	re, err := regexp.Compile(`^(?!.*\.\.)(?!.*\.$)\w[\w.]{0,29}$`)
	if err != nil {
		log.Error("Error while compiling regular expression for username validation: ", err)
		return false, err
	}
	return re.MatchString(username), nil
}

func ValidatePassword(password string) (bool, error) {
	re, err := regexp.Compile(`^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[#?!@$%^&*-]).{8,}$`)
	if err != nil {
		log.Error("Error while compiling regular expression for password validation: ", err)
		return false, err
	}
	return re.MatchString(password), nil
}

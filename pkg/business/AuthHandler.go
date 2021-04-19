package business

import (
	"fmt"
	"net/http"

	"rastebazaar/pkg/domain"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
)

// AuthHandler will hold everything that controller needs
type AuthHandler struct {
	authRepo domain.AuthRepository
}

// NewAuthHandler returns a new BaseHandler
func NewAuthHandler(AuthRepo domain.AuthRepository) *AuthHandler {
	return &AuthHandler{
		authRepo: AuthRepo,
	}
}

var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
}

// Login , in this function we hanlde send sms process
func (authHandler *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {

	phone := r.FormValue("phone")

	if phone == "" {
		w.Write([]byte(fmt.Sprintf("enter phone")))
		return
	}

	if user, err := authHandler.authRepo.Login(phone); err != nil {
		fmt.Println("Error", user)
	}

	w.Write([]byte(fmt.Sprintf("phone %v \n", phone)))
}

// Verify , in this function we hanlde send sms process
func (authHandler *AuthHandler) Verify(w http.ResponseWriter, r *http.Request) {

	phone := r.FormValue("phone")
	code := r.FormValue("code")

	if phone == "" || code == "" {
		w.Write([]byte(fmt.Sprintf("enter phone and code")))
		return

	}

	if user, err := authHandler.authRepo.Login(phone); err != nil {
		fmt.Println("Error", user)
	}

	_, tokenString, _ := tokenAuth.Encode(jwt.MapClaims{"user_id": 123})

	w.Write([]byte(fmt.Sprintf("token %v \n", tokenString)))
}

package business

import (
	"fmt"
	"net/http"

	"rastebazaar/pkg/domain"

	"github.com/dgrijalva/jwt-go"
)

// MarketHandler will hold everything that controller needs
type MarketHandler struct {
	marketRepo domain.MarketRepository
}

// NewMarketHandler returns a new BaseHandler
func NewMarketHandler(marketRepo domain.MarketRepository) *MarketHandler {
	return &MarketHandler{
		marketRepo: marketRepo,
	}
}

// MarketerLogin , in this function we hanlde send sms process for marketer
func (marketHandler *MarketHandler) MarketerLogin(w http.ResponseWriter, r *http.Request) {

	phone := r.FormValue("phone")

	if phone == "" {
		w.Write([]byte(fmt.Sprintf("enter phone")))
		return
	}

	if user, err := marketHandler.marketRepo.MarketerLogin(phone); err != nil {
		fmt.Println("Error", user)
	}

	w.Write([]byte(fmt.Sprintf("phone %v \n", phone)))
}

// MarketerVerify , in this function we hanlde check OTP process for marketer
func (marketHandler *MarketHandler) MarketerVerify(w http.ResponseWriter, r *http.Request) {

	phone := r.FormValue("phone")
	code := r.FormValue("code")

	if phone == "" || code == "" {
		w.Write([]byte(fmt.Sprintf("enter phone and code")))
		return

	}

	if user, err := marketHandler.marketRepo.MarketerVerify(phone, code); err != nil {
		fmt.Println("Error", user)
	}

	_, tokenString, _ := tokenAuth.Encode(jwt.MapClaims{"user_id": 123})

	w.Write([]byte(fmt.Sprintf("token %v \n", tokenString)))
}

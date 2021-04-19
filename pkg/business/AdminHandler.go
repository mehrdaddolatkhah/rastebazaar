package business

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"rastebazaar/pkg/business/utils"

	"rastebazaar/pkg/domain"

	"rastebazaar/pkg/domain/database"

	"github.com/dgrijalva/jwt-go"
)

// AdminHandler will hold everything that controller needs
type AdminHandler struct {
	adminRepo domain.AdminRepository
	tmpl      *template.Template
}

// NewAdminHandler returns a new BaseHandler
func NewAdminHandler(adminRepo domain.AdminRepository, tmpl *template.Template) *AdminHandler {
	return &AdminHandler{
		adminRepo: adminRepo,
		tmpl:      tmpl,
	}
}

// Register , in this fucntion we will login as admin
func (handler *AdminHandler) Register(w http.ResponseWriter, r *http.Request) {

	admin := database.Admin{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&admin); err != nil {
		w.Write([]byte(fmt.Sprintf("error: %s", err)))
		return
	}

	defer r.Body.Close()

	data, err := handler.adminRepo.AdminRegister(&admin)

	if err != nil {
		panic(err)
	}

	utils.RespondJSON(w, 201, true, "", data)
}

// Login , in this fucntion we will login as admin
func (handler *AdminHandler) Login(w http.ResponseWriter, r *http.Request) {

	admin := database.Admin{}

	phone := r.FormValue("phone")
	password := r.FormValue("password")

	if phone == "" || password == "" {
		utils.RespondJSON(w, 400, false, "phone or password are empty", nil)
		return
	}

	admin, _ = handler.adminRepo.AdminLogin(phone, password)

	if admin.Mobile == "" || admin.Password == "" {
		utils.RespondJSON(w, 403, false, "not found", nil)
		return
	}

	//encodedToken := utils.TokenExtractor(jwtauth.TokenFromHeader(r))
	// todo : add Roles into some Constants
	_, tokenString, _ := tokenAuth.Encode(jwt.MapClaims{"id": admin.ID, "role": "Admin"})

	utils.RespondJSON(w, 200, true, "success", tokenString)

}

// GetAdminRegisterView , in this fucntion we will login as admin
func (handler *AdminHandler) GetAdminRegisterView(w http.ResponseWriter, r *http.Request) {

	err := handler.tmpl.ExecuteTemplate(w, "admin-register.html", nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// GetAdminLoginView , in this fucntion we will login as admin
func (handler *AdminHandler) GetAdminLoginView(w http.ResponseWriter, r *http.Request) {

	err := handler.tmpl.ExecuteTemplate(w, "admin-login.html", nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// GetAdminPanelView , in this fucntion we will login as admin
func (handler *AdminHandler) GetAdminPanelView(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value("userId")
	id := userID.(string)

	admin, err := handler.adminRepo.GetAdminByID(id)

	if err != nil {
		panic(err)
	}

	if admin.Mobile == "" {
		utils.RespondJSON(w, 404, false, "not found", nil)
		return
	}

	err = handler.tmpl.ExecuteTemplate(w, "admin-panel.html", nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// PostAdminRegisterView , in this fucntion we will login as admin
func (handler *AdminHandler) PostAdminRegisterView(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	//admin := database.Admin{}
	firstName := r.FormValue("et_first_name")
	lastName := r.FormValue("et_last_name")
	mobile := r.FormValue("et_mobile")
	password := r.FormValue("et_password")

	if firstName == "" || lastName == "" || mobile == "" || password == "" {
		utils.RespondJSON(w, 400, false, "all fields need filed", nil)
		return
	}

	admin := database.Admin{
		Firstname: firstName,
		Lastname:  lastName,
		Mobile:    mobile,
		Password:  password,
	}

	fmt.Printf("****** %s", admin)

	_, err := handler.adminRepo.AdminRegister(&admin)

	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/admin/login", http.StatusSeeOther)

	//encodedToken := utils.TokenExtractor(jwtauth.TokenFromHeader(r))
	// todo : add Roles into some Constants
	_, tokenString, _ := tokenAuth.Encode(jwt.MapClaims{"id": admin.ID, "role": "Admin"})

	utils.RespondJSON(w, 200, true, "success", tokenString)

}

// PostAdminLoginView , in this fucntion we will login as admin
func (handler *AdminHandler) PostAdminLoginView(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("test test test")

	err := r.ParseForm()
	if err != nil {
		return
	}

	//admin := database.Admin{}
	mobile := ""
	password := ""

	//mobile := r.FormValue("phone")
	//password := r.FormValue("password")

	for key, value := range r.Form {
		if key == "et_mobile" {
			mobile = value[0]
		}

		if key == "et_password" {
			password = value[0]
		}
	}

	if mobile == "" || password == "" {
		utils.RespondJSON(w, 400, false, "phone or password is empty", nil)
		return
	}

	admin, err := handler.adminRepo.AdminLogin(mobile, password)

	if err != nil {
		panic(err)
	}

	if admin.Mobile == "" || admin.Password == "" {
		utils.RespondJSON(w, 403, false, "not found", nil)
		return
	}

	_, tokenString, _ := tokenAuth.Encode(jwt.MapClaims{"userId": admin.ID, "role": "Admin"})
	expiration := time.Now().Add(time.Hour * 48)

	accessTokenCookie := http.Cookie{
		Name:    "token",
		Path:    "/secure",
		Value:   tokenString,
		Expires: expiration,
	}

	http.SetCookie(w, &accessTokenCookie)
	http.Redirect(w, r, "/secure/admin/panel", http.StatusSeeOther)
}

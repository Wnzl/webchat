package controllers

import (
	"errors"
	"github.com/Wnzl/webchat/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm"
	"net/http"
)

type AuthController struct {
	DB      *gorm.DB
	jwtAuth *jwtauth.JWTAuth
}

type loginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type registerRequest struct {
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8,eqfield=PasswordConfirm"`
	PasswordConfirm string `json:"password_confirmation" validate:"required,min=8"`
}

var validate *validator.Validate

func NewAuthController(db *gorm.DB, jwtAuth *jwtauth.JWTAuth) *AuthController {
	validate = validator.New()

	return &AuthController{
		DB:      db,
		jwtAuth: jwtAuth,
	}
}

func (auth *AuthController) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/login", auth.login)
	r.Post("/signup", auth.signup)

	return r
}

func (auth *AuthController) login(w http.ResponseWriter, r *http.Request) {
	var request loginRequest
	if err := render.Bind(r, &request); err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	var user models.User
	auth.DB.First(&user, "email = ?", request.Email)
	if user.ID == 0 {
		render.Status(r, http.StatusUnauthorized)
		_ = render.Render(w, r, ErrInvalidRequest(errors.New("wrong credentials")))
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		render.Status(r, http.StatusUnauthorized)
		_ = render.Render(w, r, ErrInvalidRequest(errors.New("wrong credentials")))
		return
	}

	_, tokenString, err := auth.jwtAuth.Encode(map[string]interface{}{"user_id": user.ID})
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]string{
		"access token": tokenString,
	})
}

func (auth *AuthController) signup(w http.ResponseWriter, r *http.Request) {
	var request registerRequest
	if err := render.Bind(r, &request); err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	user := models.User{
		Email:    request.Email,
		Password: request.Password,
	}

	result := auth.DB.Create(&user)
	if result.Error != nil {
		_ = render.Render(w, r, ErrInvalidRequest(result.Error))
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]interface{}{
		"id":    user.ID,
		"email": user.Email,
	})
}

func (u *loginRequest) Bind(r *http.Request) error {
	if err := validate.Struct(u); err != nil {
		errs := err.(validator.ValidationErrors)
		return errs
	}

	return nil
}

func (u *registerRequest) Bind(r *http.Request) error {
	if err := validate.Struct(u); err != nil {
		errs := err.(validator.ValidationErrors)
		return errs
	}

	return nil
}
package controller

import (
	"Musicstreamuser/service"
	"net/http"

	"github.com/gorilla/mux"
)

type UserController struct {
	userService *service.UserService
}

func NewController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (user *UserController) RegisterRoutes(router *mux.Router) {
	subRouter := router.PathPrefix("/api/user").Subrouter()
	subRouter.HandleFunc("/register", user.registrationHandler).Methods(http.MethodPost)
	subRouter.HandleFunc("/login", user.loginHandler).Methods(http.MethodGet)
}

//	func (user *UserController) registrationHandler(rw http.ResponseWriter, r *http.Request) {
//		user.userService.createUser(r.Body)
//	}
func (user *UserController) registrationHandler(rw http.ResponseWriter, r *http.Request) {
	success := user.userService.CreateUser(r.Body)
	if success {
		rw.WriteHeader(http.StatusCreated)
	} else {
		http.Error(rw, "Cannot create new user", http.StatusInternalServerError)
	}

}

func (user *UserController) loginHandler(rw http.ResponseWriter, r *http.Request) {
	user.userService.LoginUser(r, rw)
}

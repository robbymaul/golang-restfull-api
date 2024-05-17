package service

import web "github.com/robbymaul/golang-swagger.git/web"

type UserService interface {
	Register(request *web.UserRegisterRequest) (web.UserRegisterResponse, error)
	Login(request web.UserLoginRequest) (web.UserLoginResponse, error)
}

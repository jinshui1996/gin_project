package controller

import (
	"gin_project/utils/jwt"
	"fmt"
)

var name_ = "admin"
var password_ = "123456"

func Login(name, password string) (string, error) {
	if name_ == name && password_ == password {
	    claims := &jwt.Claims{
			Uid: 5666,
			Username: "admin",}
		return jwt.GenerateToken(claims)
	}
	return "", fmt.Errorf("用户名或密码错误")
}
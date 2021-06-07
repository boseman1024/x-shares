package serializer

import "shares/model"

type Response struct {
	Code int         `json:"code`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type LoginResponse struct {
	Token string
	User  model.User
}

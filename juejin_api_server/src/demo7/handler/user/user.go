package user

import "demo7/model"

type CreateRequest struct {
	Username string `json:username`
	Password string `json:password`
}

type CreateResponse struct {
	Username string `json:username`
}

type ListRequest struct {
	Username string `json:"username"`
	Offset int `json:"offset"`
	Limit int `json:"limit"`
}

type ListResponse struct {
	TotalCount int `json:"totalCount"`
	UserList []*model.UserInfo `json:"userList"`
}

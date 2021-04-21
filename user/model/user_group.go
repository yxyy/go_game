package model


type UserGroup struct {
	Id int `json:"id"`
	Name string `json:"name" form:"name"`
	Mark string `json:"mark" form:"mark"`
	Description string `json:"description" form:"description"`
	CreateTime int64 `json:"create_time"`
	UpdateTime int64 `json:"update_time"`
}

package model

type Page struct {
	Page int `form:"page"`
	Length int `form:"length"`
}

func NewPage() *Page  {
	return &Page{}
}

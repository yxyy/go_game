package model

type Params struct {
	Satrt int `form:"start"`
	Length int `form:"length"`
}

type Page struct {
	Page int `form:"page"`
	Length int `form:"length"`
}

func NewPage() *Page  {
	return &Page{}
}

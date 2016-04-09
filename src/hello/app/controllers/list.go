package controllers

import "github.com/revel/revel"

type List struct {
	*revel.Controller
}

func (c List) Index() revel.Result {
	list := "List search"
	return c.Render(list)
}


func (c List) Hello() revel.Result {
	return c.Render()
}
package controllers

import "github.com/revel/revel"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	greeting := "test Hello apps"
	return c.Render(greeting)
}


func (c App) Hello() revel.Result {
	return c.Render()
}
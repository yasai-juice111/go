package controllers

import (
	"github.com/revel/revel"
	"hello/app/controllers"
	"hello/app/models"
)

type List struct {
	*revel.Controller
}

func (c List) Index() revel.Result {

 	lists := []models.List{}
 	if err := controllers.DB.Order("id desc").Find(&comments).Error; err != nil {
 		return c.HandleInternalServerError("Record Find Failure")
 	}
	// list := "List search",
	return c.Render(lists)
}
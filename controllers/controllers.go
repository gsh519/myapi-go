package controllers

import (
	"github.com/yourname/reponame/controllers/services"
)

type MyAppController struct {
	service services.MyAppServicer
}

func NewMyAppController(s services.MyAppServicer) *MyAppController {
	return &MyAppController{service: s}
}

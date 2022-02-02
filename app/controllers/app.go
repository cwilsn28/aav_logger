package controllers

import (
	"aav_logger/app/utils"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

var DBCONN = utils.OpenDB()

func (c App) Index() revel.Result {
	return c.Render()
}

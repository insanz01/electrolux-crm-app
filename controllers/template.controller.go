package controllers

import (
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/clients/coster"
	"github.com/labstack/echo/v4"
)

type TemplateController interface {
	GetMessageByTemplate(ctx echo.Context) error
}

type templateController struct {
	costerClient coster.CosterClient
}

func NewTemplateController(costerClient coster.CosterClient) TemplateController {
	return &templateController{
		costerClient: costerClient,
	}
}

func (tc *templateController) GetMessageByTemplate(ctx echo.Context) error {
	return nil
}

package controllers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/repository"

	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/helpers"
	"github.com/labstack/echo/v4"
)

type LoginController interface {
	CheckLogin(c echo.Context) error
	GenerateHashPassword(c echo.Context) error
}

type (
	loginController struct {
		repository *repository.Repository
	}
)

func NewLoginController(repository *repository.Repository) LoginController {
	return &loginController{
		repository: repository,
	}
}

func (lc *loginController) CheckLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	res, err := lc.repository.CheckLogin(username, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"messages": err.Error(),
		})
	}

	if !res {
		return echo.ErrUnauthorized
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["level"] = "application"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func (lc *loginController) GenerateHashPassword(c echo.Context) error {
	password := c.Param("password")

	hash, _ := helpers.HashPassword(password)

	return c.JSON(http.StatusOK, hash)
}

package routes

import (
	"net/http"
	// "git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/middleware"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/controllers"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/db"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/repository"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/services"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	dbConnection := db.CreateCon()

	repo := repository.New(dbConnection)

	customerService := services.NewCustomerService(repo)
	fileService := services.NewFileService(repo)
	giftService := services.NewGiftService(repo)

	loginController := controllers.NewLoginController(repo)
	customerController := controllers.NewCustomerController(customerService)
	fileController := controllers.NewFileController(fileService)
	giftController := controllers.NewGiftController(giftService)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, this is echo!")
	})

	e.GET("/generate-hash/:password", loginController.GenerateHashPassword)
	e.POST("/login", loginController.CheckLogin)

	// e.POST("/customer", customerController.)
	e.GET("/customers", customerController.FindAll)
	e.GET("/customers/:id", customerController.FindById)
	e.PUT("/customers/:id", customerController.Update)
	e.DELETE("/customers/:id", customerController.Delete)

	e.GET("/gift_claims", giftController.FindAll)
	e.GET("/gift_claims/:id", giftController.FindById)
	e.PUT("/gift_claims/:id", giftController.Update)
	e.DELETE("/gift_claims/:id", giftController.Delete)

	e.POST("/files", fileController.Upload)
	e.GET("/files/:uuid", fileController.GetFile)

	e.GET("/test-struct-validation", controllers.TestStructValidation)
	e.GET("/test-variable-validation", controllers.TestVariableValidation)

	return e
}

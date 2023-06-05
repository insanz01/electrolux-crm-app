package routes

import (
	"net/http"
	// "git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/middleware"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/controllers"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/db"
	_ "git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/docs"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/repository"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/services"

	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
)

// @title Swagger API Electrolux
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:1234
// @BasePath /api/v1
func Init() *echo.Echo {
	e := echo.New()

	dbConnection := db.CreateCon()

	repo := repository.New(dbConnection)

	customerService := services.NewCustomerService(repo)
	fileService := services.NewFileService(repo)
	giftService := services.NewGiftService(repo)
	campaignService := services.NewCampaignService(repo)

	loginController := controllers.NewLoginController(repo)
	customerController := controllers.NewCustomerController(customerService)
	fileController := controllers.NewFileController(fileService)
	giftController := controllers.NewGiftController(giftService)
	campaignController := controllers.NewCampaignController(campaignService)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	api := e.Group("api/v1")

	api.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, this is echo!")
	})

	api.GET("/generate-hash/:password", loginController.GenerateHashPassword)
	api.POST("/login", loginController.CheckLogin)

	// api.POST("/customer", customerController.)
	api.GET("/customers", customerController.FindAll)
	api.GET("/customers/:id", customerController.FindById)
	api.PUT("/customers/:id", customerController.Update)
	api.DELETE("/customers/:id", customerController.Delete)

	api.GET("/gift_claims", giftController.FindAll)
	api.GET("/gift_claims/:id", giftController.FindById)
	api.PUT("/gift_claims/:id", giftController.Update)
	api.DELETE("/gift_claims/:id", giftController.Delete)

	api.GET("/campaigns", campaignController.FindAll)
	api.GET("/campaigns/:id", campaignController.FindById)
	api.POST("/campaigns", campaignController.Insert)

	api.POST("/files", fileController.Upload)
	api.GET("/files/:uuid", fileController.GetFile)

	api.GET("/test-struct-validation", controllers.TestStructValidation)
	api.GET("/test-variable-validation", controllers.TestVariableValidation)

	return e
}

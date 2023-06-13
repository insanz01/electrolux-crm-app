package routes

import (
	"fmt"
	"net/http"

	// "git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/middleware"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/controllers"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/db"
	_ "git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/docs"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/repository"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/services"

	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	productLineService := services.NewProductLineService(repo)

	loginController := controllers.NewLoginController(repo)
	customerController := controllers.NewCustomerController(customerService)
	fileController := controllers.NewFileController(fileService)
	giftController := controllers.NewGiftController(giftService)
	campaignController := controllers.NewCampaignController(campaignService)
	productLineController := controllers.NewProductLineController(productLineService)

	// Middleware CORS
	e.Use(middleware.CORS())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/hostname", func(c echo.Context) error {
		req := c.Request()
		urlSchema := req.URL.Scheme
		if urlSchema == "" {
			urlSchema = "http"
		}

		url := fmt.Sprintf("%s://%s%s", urlSchema, req.Host, req.URL.Path)
		if req.URL.RawQuery != "" {
			url += "?" + req.URL.RawQuery
		}

		return c.String(http.StatusOK, "URL: "+url)
	})

	e.GET("/assets/:filename", fileController.Download)

	api := e.Group("api/v1")

	api.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, this is echo!")
	})

	api.GET("/generate-hash/:password", loginController.GenerateHashPassword)
	api.POST("/login", loginController.CheckLogin)

	// api.POST("/customer", customerController.)
	api.GET("/customers", customerController.FindAll)
	api.POST("/customers", customerController.FindAll)
	api.POST("/customers/:id", customerController.FindById)
	api.PUT("/customers/:id", customerController.Update)
	api.DELETE("/customers/:id", customerController.Delete)

	api.GET("/gift_claims", giftController.FindAll)
	api.POST("/gift_claims", giftController.FindAll)
	api.POST("/gift_claims/search", giftController.Search)

	api.POST("/gift_claims/:id", giftController.FindById)
	api.PUT("/gift_claims/:id", giftController.Update)
	api.DELETE("/gift_claims/:id", giftController.Delete)

	api.GET("/campaigns", campaignController.FindAll)
	api.GET("/campaigns/:id", campaignController.FindById)
	api.POST("/campaigns", campaignController.Insert)
	api.GET("/campaigns/:campaign_id/summary", campaignController.Summary)

	api.GET("/files", fileController.GetAllFile)
	api.POST("/files/filter", fileController.GetAllFileFilter)
	api.POST("/files", fileController.Upload)
	api.GET("/files/:uuid", fileController.GetFile)
	api.GET("/files/invalid", fileController.GetAllInvalidFile)

	api.GET("/product-lines", productLineController.FindAll)
	api.POST("/product-lines", productLineController.Insert)

	api.GET("/lists", customerController.List)

	api.GET("/test-struct-validation", controllers.TestStructValidation)
	api.GET("/test-variable-validation", controllers.TestVariableValidation)

	return e
}

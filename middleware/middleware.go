package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/helpers"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func AuthSSO() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Logika middleware kustom di sini
			token := helpers.GetTokenFromHeader(c.Request())
			if token == "" {
				return c.JSON(http.StatusUnauthorized, models.Response{
					Status:  0,
					Message: "invalid sso token",
					Data:    nil,
				})
			}

			tokenDetail, err := CheckToken(c.Request().Context(), token)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, helpers.GenerateOpenAPIErrorResponse(http.StatusUnauthorized))
			}

			ctx := c.Request().Context()
			newCtx := setTokenToCtx(ctx, *tokenDetail)

			c.Set("auth_token", tokenDetail)

			c.SetRequest(c.Request().WithContext(newCtx))

			return next(c)
		}
	}
}

func EchoMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := helpers.GetTokenFromHeader(c.Request())
			if token == "" {
				return c.JSON(http.StatusUnauthorized, helpers.GenerateOpenAPIErrorResponse(http.StatusUnauthorized))
			}

			tokenDetail, err := CheckToken(c.Request().Context(), token)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, helpers.GenerateOpenAPIErrorResponse(http.StatusUnauthorized))
			}

			ctx := c.Request().Context()
			newCtx := setTokenToCtx(ctx, *tokenDetail)

			c.SetRequest(c.Request().WithContext(newCtx))

			return next(c)
		}
	}
}

func CheckToken(ctx context.Context, token string) (*models.AuthSSO, error) {
	logger := logrus.WithContext(ctx).WithFields(logrus.Fields{
		"method": "CheckToken",
		"token":  token,
	})

	baseURL := "https://login.coster.id/introspect"
	queryURL := "?code="
	queryURL = fmt.Sprintf("%s%s", queryURL, token)

	url := fmt.Sprintf("%s%s", baseURL, queryURL)

	// url, err := url.JoinPath(baseURL, queryURL)
	// if err != nil {
	// 	logger.WithError(err).Error("failed to join path")
	// 	return nil, err
	// }

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		logger.WithError(err).Error("failed to create request")
		return nil, err
	}

	// req.Header.Add("Authorization", "Bearer "+token)

	httpReq := http.Client{}

	resp, err := httpReq.Do(req)
	if err != nil {
		logger.WithError(err).Error("failed to do request")
		return nil, err
	}

	defer helpers.WrapCloser(resp.Body.Close)

	fmt.Println(resp)

	switch resp.StatusCode {
	default:
		fmt.Println("gagal")
		return nil, handleCheckTokenErrorResponse(resp.Body)
	case http.StatusOK:
		fmt.Println("berhasil")
		return handleCheckTokenResponse(resp.Body)
	}
}

func setTokenToCtx(ctx context.Context, token models.AuthSSO) context.Context {
	return context.WithValue(ctx, AuthTokenKey, token)
}

func handleCheckTokenErrorResponse(body io.ReadCloser) error {
	var resp *CheckTokenErrorResponse
	if err := json.NewDecoder(body).Decode(&resp); err != nil {
		return err
	}

	return fmt.Errorf("failed to check token: message: %s status: %s code: %s ", resp.Message, resp.Status, resp.Code)
}

func handleCheckTokenResponse(body io.ReadCloser) (*models.AuthSSO, error) {
	var resp *models.AuthSSO
	if err := json.NewDecoder(body).Decode(&resp); err != nil {
		return nil, err
	}

	return resp, nil
}

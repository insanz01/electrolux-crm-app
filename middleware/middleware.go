package middleware

import (
	"github.com/labstack/echo/v4"
)

func AuthSSO(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Logika middleware kustom di sini

		// cek token di db

		// cek token di coster

		// implementasi token di db jika ada perubahan

		// tolak token jika gagal

		// Panggil fungsi berikutnya di rantai middleware
		return next(c)
	}
}

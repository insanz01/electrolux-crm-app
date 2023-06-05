package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models/dto"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/services"
	"github.com/labstack/echo/v4"
)

type FileController interface {
	Upload(c echo.Context) error
	GetFile(c echo.Context) error
}

type (
	fileController struct {
		fileService services.FileService
	}
)

func NewFileController(service services.FileService) FileController {
	return &fileController{
		fileService: service,
	}
}

func (fc *fileController) Upload(c echo.Context) error {
	// Menerima file dari permintaan HTTP
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error,
			"data":    nil,
		})
	}

	var fileUpload dto.FileRequest
	err = c.Bind(&fileUpload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error,
			"data":    nil,
		})
	}

	fileUpload.File = file

	// Membuka file yang diterima
	src, err := fileUpload.File.Open()
	if err != nil {
		fmt.Println("gagal 1")
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error,
			"data":    nil,
		})
	}
	defer src.Close()

	// Menentukan direktori tujuan penyimpanan file
	dstDir := "uploads"
	if _, err := os.Stat(dstDir); os.IsNotExist(err) {
		os.Mkdir(dstDir, 0755)
	}

	// Membuat file tujuan
	dstPath := filepath.Join(dstDir, fileUpload.File.Filename)
	dst, err := os.Create(dstPath)
	if err != nil {
		fmt.Println("gagal 2")
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error,
			"data":    nil,
		})
	}
	defer dst.Close()

	// Menyalin isi file ke file tujuan
	if _, err = io.Copy(dst, src); err != nil {
		fmt.Println("gagal 3")
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error,
			"data":    nil,
		})
	}

	fileResponse, err := fc.fileService.Insert(c, fileUpload)
	if err != nil {
		fmt.Println("gagal 4")
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error,
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "File uploaded successfully",
		"data":    fileResponse,
	})
}

func (fc *fileController) GetFile(c echo.Context) error {
	// Mendapatkan nama file dari URL parameter
	uuid := c.Param("uuid")

	if uuid == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "invalid parameter request",
			"data":    nil,
		})
	}

	// Menentukan direktori aset file
	assetDir := "uploads"

	uploadedFile, err := fc.fileService.GetDocument(c, uuid)
	if err != nil {
		return c.Stream(http.StatusInternalServerError, "application/octet-stream", nil)
	}

	// Menggabungkan direktori aset dengan nama file
	filepath := filepath.Join(assetDir, uploadedFile.Filename)

	// Membuka file
	file, err := os.Open(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return echo.NewHTTPError(http.StatusNotFound, "File not found")
		}
		return err
	}
	defer file.Close()

	// Mengirimkan file sebagai respons HTTP
	return c.Stream(http.StatusOK, "application/octet-stream", file)
}

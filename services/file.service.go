package services

import (
	"fmt"

	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models/dto"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/repository"
	"github.com/labstack/echo/v4"
)

type FileService interface {
	GetAllDocument(c echo.Context) ([]*dto.FileResponse, error)
	GetDocument(c echo.Context, uuid string) (*dto.FileResponse, error)
	Insert(c echo.Context, fileUpload dto.FileRequest) (*dto.FileResponse, error)
	GetAllInvalidDocument(c echo.Context) ([]*dto.InvalidFileResponse, error)
	GetInvalidDocument(c echo.Context, uuid string) (*dto.InvalidFileResponse, error)
}

type fileService struct {
	repository *repository.Repository
}

func NewFileService(repository *repository.Repository) FileService {
	return &fileService{
		repository: repository,
	}
}

func (fs *fileService) GetAllDocument(c echo.Context) ([]*dto.FileResponse, error) {
	files, err := fs.repository.GetAllFile()
	if err != nil {
		return nil, err
	}

	req := c.Request()
	urlSchema := req.URL.Scheme
	if urlSchema == "" {
		urlSchema = "http"
	}

	url := fmt.Sprintf("%s://%s/assets/", urlSchema, req.Host)

	fileResponse := []*dto.FileResponse{}
	for _, file := range files {
		fileResponse = append(fileResponse, &dto.FileResponse{
			UUID:      file.Id,
			Filename:  file.Filename,
			Status:    file.Status,
			Category:  file.Category,
			FilePath:  url + file.Filename,
			UpdatedAt: file.UpdatedAt,
		})
	}

	return fileResponse, nil
}

func (fs *fileService) GetDocument(c echo.Context, uuid string) (*dto.FileResponse, error) {
	file, err := fs.repository.GetFile(uuid)
	if err != nil {
		return nil, err
	}

	req := c.Request()
	urlSchema := req.URL.Scheme
	if urlSchema == "" {
		urlSchema = "http"
	}

	url := fmt.Sprintf("%s://%s/assets/", urlSchema, req.Host)

	fileResponse := dto.FileResponse{
		UUID:      file.Id,
		Filename:  file.Filename,
		Status:    file.Status,
		Category:  file.Category,
		FilePath:  url + file.Filename,
		UpdatedAt: file.UpdatedAt,
	}

	fmt.Println(fileResponse)

	return &fileResponse, nil
}

func (fs *fileService) Insert(c echo.Context, fileUpload dto.FileRequest) (*dto.FileResponse, error) {
	inputFile := models.FileExcelDocument{
		Filename:     fileUpload.File.Filename,
		Category:     fileUpload.Category,
		NumOfFailed:  0,
		NumOfSuccess: 0,
		Status:       "process",
	}

	fmt.Println(inputFile)

	uuid, err := fs.repository.UploadFile(inputFile)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	response := dto.FileResponse{
		UUID:     uuid,
		Filename: fileUpload.File.Filename,
		Status:   "process",
	}

	return &response, nil
}

func (fs *fileService) GetAllInvalidDocument(c echo.Context) ([]*dto.InvalidFileResponse, error) {
	files, err := fs.repository.GetAllInvalidFile()
	if err != nil {
		return nil, err
	}

	req := c.Request()
	urlSchema := req.URL.Scheme
	if urlSchema == "" {
		urlSchema = "http"
	}

	url := fmt.Sprintf("%s://%s/assets/", urlSchema, req.Host)

	fileResponse := []*dto.InvalidFileResponse{}
	for _, file := range files {
		fileResponse = append(fileResponse, &dto.InvalidFileResponse{
			Id:              file.Id,
			ExcelDocumentId: file.ExcelDocumentId,
			Filename:        file.Filename,
			IsValid:         file.IsValid,
			FilePath:        url + file.Filename,
			UpdatedAt:       file.UpdatedAt,
		})
	}

	return fileResponse, nil
}

func (fs *fileService) GetInvalidDocument(c echo.Context, uuid string) (*dto.InvalidFileResponse, error) {
	file, err := fs.repository.GetInvalidFile(uuid)
	if err != nil {
		return nil, err
	}

	req := c.Request()
	urlSchema := req.URL.Scheme
	if urlSchema == "" {
		urlSchema = "http"
	}

	url := fmt.Sprintf("%s://%s/invalid_uploads/", urlSchema, "localhost:4321")

	fileResponse := dto.InvalidFileResponse{
		Id:              file.Id,
		ExcelDocumentId: file.ExcelDocumentId,
		Filename:        file.Filename,
		IsValid:         file.IsValid,
		FilePath:        url + file.Filename,
		UpdatedAt:       file.UpdatedAt,
	}

	return &fileResponse, nil
}

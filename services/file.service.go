package services

import (
	"fmt"

	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models/dto"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/repository"
	"github.com/labstack/echo"
)

type FileService interface {
	GetDocument(c echo.Context, uuid string) (*dto.FileResponse, error)
	Insert(c echo.Context, fileUpload dto.FileRequest) (*dto.FileResponse, error)
}

type fileService struct {
	repository *repository.Repository
}

func NewFileService(repository *repository.Repository) FileService {
	return &fileService{
		repository: repository,
	}
}

func (fs *fileService) GetDocument(c echo.Context, uuid string) (*dto.FileResponse, error) {
	fmt.Println(uuid)

	file, err := fs.repository.GetFile(uuid)
	if err != nil {
		return nil, err
	}

	fileResponse := dto.FileResponse{
		UUID:     file.Id,
		Filename: file.Filename,
		Status:   file.Status,
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

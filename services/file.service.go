package services

import (
	"fmt"

	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models/dto"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/repository"
	"github.com/labstack/echo"
)

type FileService interface {
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

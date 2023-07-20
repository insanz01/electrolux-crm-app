package services

import (
	"fmt"

	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/clients/coster"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models/dto"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/repository"
	"github.com/labstack/echo/v4"
)

type FileService interface {
	GetAllDocument(c echo.Context) ([]*dto.FileResponse, error)
	GetAllDocumentWithFilter(c echo.Context, filters dto.FileFilterRequest) ([]*dto.FileResponse, error)
	GetDocument(c echo.Context, uuid string) (*dto.FileResponse, error)
	Insert(c echo.Context, fileUpload dto.FileRequest, userInfo *models.AuthSSO, division *coster.Division) (*dto.FileResponse, error)
	GetAllInvalidDocument(c echo.Context) ([]*dto.InvalidFileResponse, error)
	GetInvalidDocument(c echo.Context, uuid string) (*dto.InvalidFileResponse, error)
	List(c echo.Context, property string) (*dto.FileListResponse, error)
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
		invalidFile := dto.InvalidFile{}

		if file.InvalidFilename != nil {
			invalidFile.Filename = file.InvalidFilename
			invalidFile.IsValid = file.IsValid
			invalidFile.FilePath = url + *file.InvalidFilename
		}

		divisionId := ""
		divisionName := ""
		if file.UploadByDivisionId != nil {
			divisionId = *file.UploadByDivisionId
		}
		if file.UploadByDivisionName != nil {
			divisionName = *file.UploadByDivisionName
		}

		fileResponse = append(fileResponse, &dto.FileResponse{
			UUID:                 file.Id,
			Filename:             file.Filename,
			NumOfFailed:          file.NumOfFailed,
			NumOfSuccess:         file.NumOfSuccess,
			Status:               file.Status,
			Category:             file.Category,
			FilePath:             url + file.Filename,
			InvalidFile:          &invalidFile,
			UpdatedAt:            file.UpdatedAt,
			UploadByUserId:       *file.UploadByUserId,
			UploadByUserName:     *file.UploadByUserName,
			UploadByDivisionId:   divisionId,
			UploadByDivisionName: divisionName,
		})
	}

	return fileResponse, nil
}

func (fs *fileService) GetAllDocumentWithFilter(c echo.Context, filter dto.FileFilterRequest) ([]*dto.FileResponse, error) {
	files, err := fs.repository.GetAllFileWithFilter(filter.Filters)
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
		invalidFile := dto.InvalidFile{}

		if file.InvalidFilename != nil {
			invalidFile.Filename = file.InvalidFilename
			invalidFile.IsValid = file.IsValid
			invalidFile.FilePath = url + *file.InvalidFilename
		}

		fileResponse = append(fileResponse, &dto.FileResponse{
			UUID:                 file.Id,
			Filename:             file.Filename,
			NumOfFailed:          file.NumOfFailed,
			NumOfSuccess:         file.NumOfSuccess,
			Status:               file.Status,
			Category:             file.Category,
			FilePath:             url + file.Filename,
			InvalidFile:          &invalidFile,
			UpdatedAt:            file.UpdatedAt,
			UploadByUserId:       *file.UploadByUserId,
			UploadByUserName:     *file.UploadByUserName,
			UploadByDivisionId:   *file.UploadByDivisionId,
			UploadByDivisionName: *file.UploadByDivisionName,
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
		UUID:         file.Id,
		Filename:     file.Filename,
		NumOfFailed:  file.NumOfFailed,
		NumOfSuccess: file.NumOfSuccess,
		Status:       file.Status,
		Category:     file.Category,
		FilePath:     url + file.Filename,
		UpdatedAt:    file.UpdatedAt,
	}

	fmt.Println(fileResponse)

	return &fileResponse, nil
}

func (fs *fileService) Insert(c echo.Context, fileUpload dto.FileRequest, userInfo *models.AuthSSO, division *coster.Division) (*dto.FileResponse, error) {
	inputFile := models.FileExcelDocument{
		Filename:             fileUpload.File.Filename,
		Category:             fileUpload.Category,
		NumOfFailed:          0,
		NumOfSuccess:         0,
		Status:               "process",
		UploadByUserId:       userInfo.User.ID,
		UploadByUserName:     userInfo.User.Name,
		UploadByDivisionId:   &division.ID,
		UploadByDivisionName: &division.Name,
	}

	fmt.Println(inputFile)

	uuid, err := fs.repository.UploadFile(inputFile)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	response := dto.FileResponse{
		UUID:             uuid,
		Filename:         fileUpload.File.Filename,
		Status:           "process",
		UploadByUserId:   *inputFile.UploadByUserId,
		UploadByUserName: *inputFile.UploadByUserName,
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

	url := fmt.Sprintf("%s://%s/uploads/", urlSchema, "localhost:4321")

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

func (fs *fileService) List(c echo.Context, property string) (*dto.FileListResponse, error) {
	lists, err := fs.repository.GetAllFile()
	if err != nil {
		return nil, err
	}

	listResponse := []string{}
	unique := make(map[string]bool)
	for _, list := range lists {
		divisionName := ""
		if list.UploadByDivisionName != nil {
			divisionName = *list.UploadByDivisionName
		}
		if !unique[divisionName] && divisionName != "" {
			unique[divisionName] = true
			listResponse = append(listResponse, divisionName)
		}
	}

	return &dto.FileListResponse{
		ListData: listResponse,
	}, nil
}

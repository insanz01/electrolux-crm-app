package services

import (
	"fmt"

	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models/dto"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/repository"
	"github.com/labstack/echo/v4"
)

type ProductLineService interface {
	FindAll(c echo.Context) (*dto.ProductLineResponse, error)
	FindById(c echo.Context, id string) (*dto.ProductLineResponse, error)
	Insert(c echo.Context, productLine dto.ProductLineInsertRequest) (*dto.ProductLineResponse, error)
	Update(c echo.Context, productLine dto.ProductLineUpdateRequest, id string) (*dto.ProductLineResponse, error)
	Delete(c echo.Context, id string) error
}

type productLineService struct {
	repository *repository.Repository
}

func NewProductLineService(repository *repository.Repository) ProductLineService {
	return &productLineService{
		repository: repository,
	}
}

func (pls *productLineService) FindAll(c echo.Context) (*dto.ProductLineResponse, error) {
	productLines, err := pls.repository.GetAllProductLine()
	if err != nil {
		return nil, err
	}

	groupedProductLines := make(map[string][]*models.ProductLineProperties)
	for _, productLine := range productLines {
		groupId := productLine.TableDataID
		groupedProductLines[groupId] = append(groupedProductLines[groupId], productLine)
	}

	return &dto.ProductLineResponse{
		ProductLine: groupedProductLines,
	}, nil
}

func (pls *productLineService) FindById(c echo.Context, id string) (*dto.ProductLineResponse, error) {
	productLines, err := pls.repository.GetSingleProductLine(id)
	if err != nil {
		return nil, err
	}

	groupedProductLines := make(map[string][]*models.ProductLineProperties)
	for _, productLine := range productLines {
		groupId := productLine.TableDataID
		groupedProductLines[groupId] = append(groupedProductLines[groupId], productLine)
	}

	return &dto.ProductLineResponse{
		ProductLine: groupedProductLines,
	}, nil
}

func (pls *productLineService) Insert(c echo.Context, productLine dto.ProductLineInsertRequest) (*dto.ProductLineResponse, error) {
	tableList, err := pls.repository.FindIdTableCategoryByName("product_line")
	if err != nil {
		return nil, err
	}

	insertTableData := models.TableData{
		TableId: tableList.Id,
	}

	tableDataId, err := pls.repository.InsertTableData(insertTableData)
	if err != nil {
		return nil, err
	}

	productLineCodeInsert := models.TableProperty{
		TableDataID: tableDataId,
		OrderNumber: 1,
		Name:        "Code",
		Key:         "code",
		Value:       productLine.Code,
		Datatype:    "string",
		IsMandatory: true,
		InputType:   "text",
	}

	codeId, err := pls.repository.InsertTableProperty(productLineCodeInsert)
	if err != nil {
		return nil, err
	}

	productLineValueInsert := models.TableProperty{
		TableDataID: tableDataId,
		OrderNumber: 2,
		Name:        "Value",
		Key:         "value",
		Value:       productLine.Value,
		Datatype:    "string",
		IsMandatory: true,
		InputType:   "text",
	}

	valueId, err := pls.repository.InsertTableProperty(productLineValueInsert)
	if err != nil {
		return nil, err
	}

	fmt.Println(codeId, valueId)

	productLines, err := pls.repository.GetSingleProductLine(tableDataId)
	if err != nil {
		return nil, err
	}

	groupedProductLines := make(map[string][]*models.ProductLineProperties)
	for _, productLine := range productLines {
		groupId := productLine.TableDataID
		groupedProductLines[groupId] = append(groupedProductLines[groupId], productLine)
	}

	return &dto.ProductLineResponse{
		ProductLine: groupedProductLines,
	}, nil
}

func (pls *productLineService) Update(c echo.Context, productLine dto.ProductLineUpdateRequest, id string) (*dto.ProductLineResponse, error) {
	return nil, nil
}

func (pls *productLineService) Delete(c echo.Context, id string) error {
	return nil
}

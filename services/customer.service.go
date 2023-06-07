package services

import (
	"fmt"

	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models/dto"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/repository"
	"github.com/labstack/echo/v4"
)

type CustomerService interface {
	FindAll(c echo.Context) (*dto.CustomerResponse, error)
	FindByPropsOrFilter(c echo.Context, customerProperties dto.CustomerProperties) (*dto.CustomerResponse, error)
	FindById(c echo.Context, uuid string) (*dto.CustomerResponse, error)
	FindByIdWithPropsOrFilter(c echo.Context, customerProperties dto.CustomerProperties, id string) (*dto.CustomerResponse, error)
	// Insert(c echo.Context, customer models.CustomerInsert) (dto.CustomerResponse, error)
	Update(c echo.Context, customer dto.CustomerUpdateRequest, uuid string) (*dto.CustomerResponse, error)
	Delete(c echo.Context, uuid string) error
	List(c echo.Context, property string) (*dto.ListResponse, error)
}

type customerService struct {
	repository *repository.Repository
}

func NewCustomerService(repository *repository.Repository) CustomerService {
	return &customerService{
		repository: repository,
	}
}

func (cs *customerService) FindAll(c echo.Context) (*dto.CustomerResponse, error) {

	pagination := c.Get("pagination").(models.Pagination)

	customers, err := cs.repository.GetAll(pagination)
	if err != nil {
		return nil, err
	}

	groupedCustomers := make(map[string][]*models.CustomerProperties)
	for _, customer := range customers {
		groupId := customer.TableDataID
		groupedCustomers[groupId] = append(groupedCustomers[groupId], customer)
	}

	// allCustomers := dto.GroupCustomer{}

	// for groupId, gCustomer := range groupedCustomers {
	// 	singleCustomer := []dto.Customer{}
	// 	for _, sCustomer := range gCustomer {
	// 		singleCustomer = append(singleCustomer, dto.Customer{
	// 			Id:          sCustomer.ID,
	// 			TableDataID: sCustomer.TableDataID,
	// 			Key:         sCustomer.Key,
	// 			Value:       sCustomer.Value,
	// 			Datatype:    sCustomer.Datatype,
	// 			IsMandatory: sCustomer.IsMandatory,
	// 			InputType:   sCustomer.InputType,
	// 		})
	// 	}

	// 	allCustomers.CustomerData = append(allCustomers.CustomerData, singleCustomer...)
	// 	allCustomers.GroupId = groupId
	// }

	return &dto.CustomerResponse{
		Customer: groupedCustomers,
	}, nil
}

func (cs *customerService) FindByPropsOrFilter(c echo.Context, customerProperties dto.CustomerProperties) (*dto.CustomerResponse, error) {
	customers, err := cs.repository.GetAllWithFilter(customerProperties)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	groupedCustomers := make(map[string][]*models.CustomerProperties)
	for _, customer := range customers {
		groupId := customer.TableDataID
		groupedCustomers[groupId] = append(groupedCustomers[groupId], customer)
	}

	return &dto.CustomerResponse{
		Customer: groupedCustomers,
	}, nil
}

func (cs *customerService) FindById(c echo.Context, uuid string) (*dto.CustomerResponse, error) {
	customer, err := cs.repository.GetSingle(uuid)
	if err != nil {
		return nil, err
	}

	groupedCustomers := make(map[string][]*models.CustomerProperties)
	for _, customer := range customer {
		groupId := customer.TableDataID
		groupedCustomers[groupId] = append(groupedCustomers[groupId], customer)
	}

	return &dto.CustomerResponse{
		Customer: groupedCustomers,
	}, nil
}

func (cs *customerService) FindByIdWithPropsOrFilter(c echo.Context, customerProperties dto.CustomerProperties, id string) (*dto.CustomerResponse, error) {
	customer, err := cs.repository.GetSingleCustomerWithFilter(customerProperties, id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	groupedCustomers := make(map[string][]*models.CustomerProperties)
	for _, customer := range customer {
		groupId := customer.TableDataID
		groupedCustomers[groupId] = append(groupedCustomers[groupId], customer)
	}

	return &dto.CustomerResponse{
		Customer: groupedCustomers,
	}, nil
}

func (cs *customerService) Insert(c echo.Context, customer models.CustomerInsert) (dto.CustomerResponse, error) {
	return dto.CustomerResponse{}, nil
}

func (cs *customerService) Update(c echo.Context, customer dto.CustomerUpdateRequest, uuid string) (*dto.CustomerResponse, error) {

	updateData := []models.CustomerProperties{}
	for _, data := range customer.Customers {
		updateData = append(updateData, models.CustomerProperties{
			ID:          data.ID,
			TableDataID: uuid,
			Name:        data.Name,
			Key:         data.Key,
			Value:       data.Value,
			Datatype:    data.Datatype,
			IsMandatory: data.IsMandatory,
			InputType:   data.InputType,
		})
	}

	err := cs.repository.UpdateCustomer(&updateData)
	if err != nil {
		return nil, err
	}

	err = cs.repository.UpdateDate(uuid)
	if err != nil {
		return nil, err
	}

	singleCustomer, err := cs.repository.GetSingle(uuid)
	if err != nil {
		return nil, err
	}

	groupedCustomers := make(map[string][]*models.CustomerProperties)
	for _, customer := range singleCustomer {
		groupId := customer.TableDataID
		groupedCustomers[groupId] = append(groupedCustomers[groupId], customer)
	}

	return &dto.CustomerResponse{
		Customer: groupedCustomers,
	}, nil
}

func (cs *customerService) Delete(c echo.Context, uuid string) error {
	err := cs.repository.DeleteCustomer(uuid)
	if err != nil {
		return err
	}

	return nil
}

func (cs *customerService) List(c echo.Context, property string) (*dto.ListResponse, error) {
	lists, err := cs.repository.GetList(property)
	if err != nil {
		return nil, err
	}

	listResponse := []dto.ListData{}

	unique := make(map[dto.ListData]bool)
	for _, list := range lists {
		tempList := dto.ListData{
			Key:   list.Key,
			Value: list.Value,
		}

		if !unique[tempList] {
			unique[tempList] = true
			listResponse = append(listResponse, tempList)
		}
	}

	return &dto.ListResponse{
		ListData: listResponse,
	}, nil
}

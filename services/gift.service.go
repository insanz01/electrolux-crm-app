package services

import (
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models/dto"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/repository"
	"github.com/labstack/echo/v4"
)

type GiftService interface {
	FindAll(c echo.Context) (*dto.GiftClaimResponse, error)
	FindByPropsOrFilter(c echo.Context, giftClaimResponse dto.GiftClaimProperties) (*dto.GiftClaimResponse, error)
	FindById(c echo.Context, uuid string) (*dto.GiftClaimResponse, error)
	FindByIdWithPropsOrFilter(c echo.Context, giftProperties dto.GiftClaimProperties, uuid string) (*dto.GiftClaimResponse, error)
	Update(c echo.Context, giftClaim dto.GiftClaimUpdateRequest, uuid string) (*dto.GiftClaimResponse, error)
	Delete(c echo.Context, uuid string) error
}

type giftService struct {
	repository *repository.Repository
}

func NewGiftService(repository *repository.Repository) GiftService {
	return &giftService{
		repository: repository,
	}
}

func (gs *giftService) FindAll(c echo.Context) (*dto.GiftClaimResponse, error) {

	giftClaims, err := gs.repository.GetAllGiftClaim()
	if err != nil {
		return nil, err
	}

	groupedGiftClaims := make(map[string][]*models.GiftProperties)
	for _, gift := range giftClaims {
		groupId := gift.TableDataID
		groupedGiftClaims[groupId] = append(groupedGiftClaims[groupId], gift)
	}

	return &dto.GiftClaimResponse{
		GiftClaim: groupedGiftClaims,
	}, nil
}

func (gs *giftService) FindByPropsOrFilter(c echo.Context, giftProperties dto.GiftClaimProperties) (*dto.GiftClaimResponse, error) {
	giftClaims, err := gs.repository.GetAllGiftClaimWithFilter(giftProperties)
	if err != nil {
		return nil, err
	}

	groupedGiftClaims := make(map[string][]*models.CustomerProperties)
	for _, giftClaim := range giftClaims {
		groupId := giftClaim.TableDataID
		groupedGiftClaims[groupId] = append(groupedGiftClaims[groupId], giftClaim)
	}

	return &dto.GiftClaimResponse{
		GiftClaim: groupedGiftClaims,
	}, nil
}

func (gs *giftService) FindById(c echo.Context, uuid string) (*dto.GiftClaimResponse, error) {
	giftClaim, err := gs.repository.GetSingleGiftClaim(uuid)
	if err != nil {
		return nil, err
	}

	groupedGiftClaim := make(map[string][]*models.GiftProperties)
	for _, gift := range giftClaim {
		groupId := gift.TableDataID
		groupedGiftClaim[groupId] = append(groupedGiftClaim[groupId], gift)
	}

	return &dto.GiftClaimResponse{
		GiftClaim: groupedGiftClaim,
	}, nil
}

func (gs *giftService) FindByIdWithPropsOrFilter(c echo.Context, giftProperties dto.GiftClaimProperties, id string) (*dto.GiftClaimResponse, error) {
	giftClaims, err := gs.repository.GetSingleGiftClaimWithFilter(giftProperties, id)
	if err != nil {
		return nil, err
	}

	groupedGiftClaims := make(map[string][]*models.CustomerProperties)
	for _, giftClaim := range giftClaims {
		groupId := giftClaim.TableDataID
		groupedGiftClaims[groupId] = append(groupedGiftClaims[groupId], giftClaim)
	}

	return &dto.GiftClaimResponse{
		GiftClaim: groupedGiftClaims,
	}, nil
}

func (gs *giftService) Update(c echo.Context, giftClaim dto.GiftClaimUpdateRequest, uuid string) (*dto.GiftClaimResponse, error) {
	updateData := []models.GiftProperties{}
	for _, data := range giftClaim.GiftClaims {
		updateData = append(updateData, models.GiftProperties{
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

	err := gs.repository.UpdateGiftClaim(&updateData)
	if err != nil {
		return nil, err
	}

	err = gs.repository.UpdateDate(uuid)
	if err != nil {
		return nil, err
	}

	singleGiftClaim, err := gs.repository.GetSingleGiftClaim(uuid)
	if err != nil {
		return nil, err
	}

	groupedGiftClaim := make(map[string][]*models.GiftProperties)
	for _, gift := range singleGiftClaim {
		groupId := gift.TableDataID
		groupedGiftClaim[groupId] = append(groupedGiftClaim[groupId], gift)
	}

	return &dto.GiftClaimResponse{
		GiftClaim: groupedGiftClaim,
	}, nil
}

func (gs *giftService) Delete(c echo.Context, uuid string) error {
	err := gs.repository.DeleteGiftClaim(uuid)
	if err != nil {
		return err
	}

	return nil
}

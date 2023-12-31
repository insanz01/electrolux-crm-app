package repository

import (
	"fmt"

	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models/dto"
	"github.com/jmoiron/sqlx"
)

type GiftRepository interface {
	GetAllGiftClaim() ([]*models.GiftProperties, error)
	GetSingleGiftClaim(id string) ([]*models.GiftProperties, error)
	UpdateGiftClaim(gifts *[]models.GiftProperties) error
	DeleteGiftClaim(tableId string) error
	// refactor nanti
	SearchGiftClaim(keyValue models.KeyValue) ([]*models.GiftProperties, error)
}

const (
	getAllTableIdGiftQuery       = "SELECT public.table_data.id, public.table_data.table_id FROM public.table_data JOIN public.table_list ON public.table_data.table_id = public.table_list.id WHERE public.table_list.name = 'gift_claim' AND public.table_data.deleted_at is null"
	getAllGiftQuery              = "SELECT public.properties.id, public.properties.table_data_id, public.properties.order_number, public.properties.name, public.properties.key, public.properties.value, public.properties.datatype, public.properties.is_mandatory, public.properties.input_type, public.table_data.updated_at FROM public.properties JOIN public.table_data ON public.properties.table_data_id = public.table_data.id JOIN public.table_list ON public.table_data.table_id = public.table_list.id WHERE public.table_list.name = 'gift_claim' AND public.table_data.deleted_at is null"
	getAllGiftWithFilterQuery    = "SELECT public.properties.id, public.properties.table_data_id, public.properties.order_number, public.properties.name, public.properties.key, public.properties.value, public.properties.datatype, public.properties.is_mandatory, public.properties.input_type, public.table_data.updated_at FROM public.properties JOIN public.table_data ON public.properties.table_data_id = public.table_data.id JOIN public.table_list ON public.table_data.table_id = public.table_list.id WHERE public.table_list.name = 'gift_claim' AND public.table_data.deleted_at is null"
	getSingleGiftQuery           = "SELECT public.properties.id, public.properties.table_data_id, public.properties.order_number, public.properties.name, public.properties.key, public.properties.value, public.properties.datatype, public.properties.is_mandatory, public.properties.input_type, public.table_data.updated_at FROM public.properties JOIN public.table_data ON public.properties.table_data_id = public.table_data.id JOIN public.table_list ON public.table_data.table_id = public.table_list.id WHERE public.table_list.name = 'gift_claim' AND public.properties.table_data_id = $1 AND public.table_data.deleted_at is null"
	getSingleGiftWithFilterQuery = "SELECT public.properties.id, public.properties.table_data_id, public.properties.order_number, public.properties.name, public.properties.key, public.properties.value, public.properties.datatype, public.properties.is_mandatory, public.properties.input_type, public.table_data.updated_at FROM public.properties JOIN public.table_data ON public.properties.table_data_id = public.table_data.id JOIN public.table_list ON public.table_data.table_id = public.table_list.id WHERE public.table_list.name = 'gift_claim' AND public.properties.table_data_id = ? AND public.table_data.deleted_at is null"
	updateGiftQuery              = "UPDATE public.properties SET properties.value = :value, properties.updated_at = now() WHERE properties.key = :key AND properties.table_data_id = :table_data_id"
	deleteGiftQuery              = "UPDATE public.table_data SET deleted_at = NOW() WHERE properties.id = :id"
	searchByKeyValueQuery        = "SELECT public.properties.id, public.properties.table_data_id, public.properties.order_number, public.properties.name, public.properties.key, public.properties.value, public.properties.datatype, public.properties.is_mandatory, public.properties.input_type, public.table_data.updated_at FROM public.properties JOIN public.table_data ON public.properties.table_data_id = public.table_data.id JOIN public.table_list ON public.table_data.table_id = public.table_list.id WHERE public.table_list.name = 'gift_claim' AND public.table_data.deleted_at is null AND public.properties.key = $1 AND public.properties.value = $2"

// insertQuery = "INSERT INTO customer () VALUES ()"
)

func (r *Repository) GetAllGiftClaim() ([]*models.GiftProperties, error) {
	var giftClaims []*models.GiftProperties

	err := r.db.Select(&giftClaims, getAllGiftQuery)
	if err != nil {
		return nil, err
	}

	return giftClaims, nil
}

func (r *Repository) GetAllGiftClaimWithFilter(properties dto.GiftClaimProperties) ([]*models.GiftProperties, error) {
	var giftClaims []*models.GiftProperties

	finalQuery := getAllGiftWithFilterQuery
	var tableIds []string

	useFilter := false
	useProperties := false

	if properties.Properties != nil {
		useProperties = true

		finalQuery = fmt.Sprintf("%s AND public.properties.key in (?)", finalQuery)
	}

	if properties.Filters != nil {
		tempTableIds, err := r.GetTableIdByValue(properties.Filters)
		if err != nil {
			if err.Error() == "many_rows_data" {
				return nil, nil
			}

			return nil, err
		}

		tableIds = tempTableIds

		finalQuery = fmt.Sprintf("%s AND public.properties.table_data_id IN (?)", finalQuery)

		useFilter = true
	}

	if useFilter && useProperties {
		// Persiapan query
		query, args, err := sqlx.In(finalQuery, properties.Properties, tableIds)
		if err != nil {
			return nil, err
		}

		query = sqlx.Rebind(sqlx.DOLLAR, query)

		// Eksekusi query
		err = r.db.Select(&giftClaims, query, args...)
		if err != nil {
			return nil, err
		}

		return giftClaims, err
	} else if useFilter {
		// Persiapan query
		query, args, err := sqlx.In(finalQuery, tableIds)
		if err != nil {
			return nil, err
		}

		query = sqlx.Rebind(sqlx.DOLLAR, query)

		// Eksekusi query
		err = r.db.Select(&giftClaims, query, args...)
		if err != nil {
			return nil, err
		}

		return giftClaims, err
	}

	// Persiapan query
	query, args, err := sqlx.In(finalQuery, properties.Properties)
	if err != nil {
		return nil, err
	}

	query = sqlx.Rebind(sqlx.DOLLAR, query)

	// Eksekusi query
	err = r.db.Select(&giftClaims, query, args...)
	if err != nil {
		return nil, err
	}

	return giftClaims, err
}

func (r *Repository) GetSingleGiftClaim(id string) ([]*models.GiftProperties, error) {
	var giftClaims []*models.GiftProperties

	err := r.db.Select(&giftClaims, getSingleGiftQuery, id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if len(giftClaims) == 0 {
		return nil, nil
	}

	return giftClaims, nil
}

func (r *Repository) GetSingleGiftClaimWithFilter(props dto.GiftClaimProperties, id string) ([]*models.CustomerProperties, error) {
	var giftClaim []*models.CustomerProperties

	// Persiapan query
	query, args, err := sqlx.In(getSingleGiftWithFilterQuery, id, props.Properties)
	if err != nil {
		return nil, err
	}

	query = sqlx.Rebind(sqlx.DOLLAR, query)

	// Eksekusi query
	err = r.db.Select(&giftClaim, query, args...)
	if err != nil {
		return nil, err
	}

	return giftClaim, err
}

func (r *Repository) UpdateGiftClaim(gifts *[]models.GiftProperties) error {
	for _, gift := range *gifts {
		_, err := r.db.NamedExec(updateGiftQuery, gift)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Repository) DeleteGiftClaim(tableId string) error {
	_, err := r.db.NamedExec(deleteGiftQuery, models.DeleteTableData{Id: tableId})
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) SearchGiftClaim(keyValue models.KeyValue) ([]*models.GiftProperties, error) {
	var giftClaims []*models.GiftProperties

	err := r.db.Select(&giftClaims, searchByKeyValueQuery, keyValue.Key, keyValue.Value)
	if err != nil {
		return nil, err
	}

	return giftClaims, nil
}

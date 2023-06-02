package repository

import (
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models/dto"
)

type ProductLineRepository interface {
	GetAllProductLine() ([]*models.ProductLineProperties, error)
	GetSingleProductLine(id string) ([]*models.ProductLineProperties, error)
	// insert using table repository
	// InsertProductLine() (string, error)
	UpdateProductLine(updateProductLine dto.ProductLineUpdateRequest) error
}

const (
	getAllProductLineQuery    = "SELECT public.properties.id, public.properties.table_data_id, public.properties.order_number, public.properties.name, public.properties.key, public.properties.value, public.properties.datatype, public.properties.is_mandatory, public.properties.input_type FROM public.properties JOIN public.table_data ON public.properties.table_data_id = public.table_data.id JOIN public.table_list ON public.table_data.table_id = public.table_list.id WHERE public.table_list.name = 'product_line' AND public.table_data.deleted_at is null"
	getSingleProductLineQuery = "SELECT public.properties.id, public.properties.table_data_id, public.properties.order_number, public.properties.name, public.properties.key, public.properties.value, public.properties.datatype, public.properties.is_mandatory, public.properties.input_type FROM public.properties JOIN public.table_data ON public.properties.table_data_id = public.table_data.id JOIN public.table_list ON public.table_data.table_id = public.table_list.id WHERE public.table_list.name = 'product_line' AND public.properties.table_data_id = $1 AND public.table_data.deleted_at is null"
)

func (r *Repository) GetAllProductLine() ([]*models.ProductLineProperties, error) {
	var productLines []*models.ProductLineProperties

	err := r.db.Select(&productLines, getAllProductLineQuery)
	if err != nil {
		return nil, err
	}

	return productLines, nil
}

func (r *Repository) GetSingleProductLine(id string) ([]*models.ProductLineProperties, error) {
	var productLines []*models.ProductLineProperties

	err := r.db.Select(&productLines, getSingleProductLineQuery, id)
	if err != nil {
		return nil, err
	}

	if len(productLines) == 0 {
		return nil, nil
	}

	return productLines, nil
}

package repository

import (
	"fmt"

	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models"
)

type CustomerRepository interface {
	GetAll() ([]*models.CustomerProperties, error)
	GetSingle(id string) ([]*models.CustomerProperties, error)
	UpdateCustomer(customers *[]models.CustomerProperties) error
	DeleteCustomer(tableId string) error
}

const (
	getAllTableId       = "SELECT public.table_data.id, public.table_data.table_id FROM public.table_data JOIN public.table_list ON public.table_data.table_id = public.table_list.id WHERE public.table_list.name = 'customer' AND public.table_data.deleted_at is null"
	getAllQuery         = "SELECT public.properties.id, public.properties.table_data_id, public.properties.order_number, public.properties.name, public.properties.key, public.properties.value, public.properties.datatype, public.properties.is_mandatory, public.properties.input_type FROM public.properties JOIN public.table_data ON public.properties.table_data_id = public.table_data.id JOIN public.table_list ON public.table_data.table_id = public.table_list.id WHERE public.table_list.name = 'customer' AND public.table_data.deleted_at is null"
	getSingleQuery      = "SELECT public.properties.id, public.properties.table_data_id, public.properties.order_number, public.properties.name, public.properties.key, public.properties.value, public.properties.datatype, public.properties.is_mandatory, public.properties.input_type FROM public.properties JOIN public.table_data ON public.properties.table_data_id = public.table_data.id JOIN public.table_list ON public.table_data.table_id = public.table_list.id WHERE public.table_list.name = 'customer' AND public.properties.table_data_id = $1 AND public.table_data.deleted_at is null"
	updateCustomerQuery = "UPDATE public.properties SET value = :value, updated_at = NOW() WHERE key = :key AND table_data_id = :table_data_id"
	deleteCustomerQuery = "UPDATE public.properties SET deleted_at = NOW() WHERE id = :id"

// insertQuery = "INSERT INTO customer () VALUES ()"
)

func (r *Repository) GetAll() ([]*models.CustomerProperties, error) {
	var customers []*models.CustomerProperties

	err := r.db.Select(&customers, getAllQuery)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (r *Repository) GetSingle(id string) ([]*models.CustomerProperties, error) {
	var customers []*models.CustomerProperties

	err := r.db.Select(&customers, getSingleQuery, id)
	if err != nil {
		return nil, err
	}

	if len(customers) < 1 {
		return nil, nil
	}

	return customers, nil
}

func (r *Repository) UpdateCustomer(customers *[]models.CustomerProperties) error {
	for _, customer := range *customers {
		_, err := r.db.NamedExec(updateCustomerQuery, customer)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}

func (r *Repository) DeleteCustomer(tableId string) error {
	_, err := r.db.NamedExec(deleteCustomerQuery, models.DeleteTableData{Id: tableId})
	if err != nil {
		return err
	}
	return nil
}

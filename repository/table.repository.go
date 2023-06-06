package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models"
)

type TableRepository interface {
	FindIdTableCategoryByName(name string) (*models.TableCategory, error)
	InsertTableData(tableData models.TableData) (string, error)
	InsertTableProperty(property models.TableProperty) (string, error)
	UpdatedDate(tableId string) (string, error)
}

const (
	getTableListCategoryQuery = "SELECT id, name FROM public.table_list WHERE name = $1"
	getUpdatedDateQuery       = "SELECT public.table_data.updated_at FROM public.table_data WHERE public.table_data.id = $1"
	insertTableDataQuery      = "INSERT INTO public.table_data (table_id) VALUES (:table_id) returning id"
	insertTableProperty       = "INSERT INTO public.properties (table_data_id, order_number, name, key, value, datatype, is_mandatory, input_type) VALUES (:table_data_id, :order_number, :name, :key, :value, :datatype, :is_mandatory, :input_type) returning id"
	updateDateQuery           = "UPDATE public.table_data SET updated_at = NOW() WHERE id = :table_data_id"
)

func (r *Repository) FindIdTableCategoryByName(name string) (*models.TableCategory, error) {
	tableList := []models.TableCategory{}

	err := r.db.Select(&tableList, getTableListCategoryQuery, name)
	if err == sql.ErrNoRows {
		fmt.Println("No data found")
		return nil, err
	}

	if err != nil {
		fmt.Println("Query error")
		return nil, err
	}

	if len(tableList) < 1 {
		return nil, errors.New("tidak ada data")
	}

	return &tableList[0], nil
}

func (r *Repository) InsertTableData(tableData models.TableData) (string, error) {
	stmt, err := r.db.PrepareNamed(insertTableDataQuery)
	if err != nil {
		return "", errors.New("insert_table_data" + err.Error())
	}
	defer stmt.Close()

	var uuid string
	err = stmt.QueryRow(&tableData).Scan(&uuid)
	if err != nil {
		return "", errors.New("insert_table_data" + err.Error())
	}
	return uuid, nil
}

func (r *Repository) InsertTableProperty(property models.TableProperty) (string, error) {
	stmt, err := r.db.PrepareNamed(insertTableProperty)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("insert_table_property" + err.Error())
	}
	defer stmt.Close()

	var uuid string
	err = stmt.QueryRow(&property).Scan(&uuid)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("insert_table_property" + err.Error())
	}
	return uuid, nil
}

func (r *Repository) UpdateDate(tableId string) error {
	var updatedData struct {
		TableDataId string `db:"table_data_id"`
	}

	_, err := r.db.NamedExec(updateDateQuery, updatedData)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetUpdatedDate(tableId string) (string, error) {
	var theDated struct {
		UpdatedAt string `db:"updated_at"`
	}

	err := r.db.Select(&theDated, getUpdatedDateQuery, tableId)
	if err != nil {
		return "", err
	}

	return theDated.UpdatedAt, nil
}

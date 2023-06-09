package repository

import (
	"errors"
	"fmt"

	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models/dto"
)

type FileRepository interface {
	UploadFile(filename string) (string, error)
	GetAllFile() ([]*models.FileExcelDocument, error)
	GetFile(id string) (*models.FileExcelDocument, error)
	GetAllFileWithFilter(filters []dto.FileFilter) ([]*models.FileExcelDocument, error)
	GetAllInvalidFile() ([]*models.InvalidFileExcelDocument, error)
	GetInvalidFile(id string) (*models.InvalidFileExcelDocument, error)
}

const (
	insertFileQuery        = "INSERT INTO excel_document (filename, category, num_of_failed, num_of_success, status) VALUES (:filename, :category, :num_of_failed, :num_of_success, :status) returning id"
	getFileQuery           = "SELECT id, filename, category, num_of_failed, num_of_success, status, created_at, updated_at, deleted_at FROM public.excel_document WHERE id = $1"
	getAllFileQuery        = "SELECT id, filename, category, num_of_failed, num_of_success, status, created_at, updated_at, deleted_at FROM public.excel_document WHERE deleted_at is NULL"
	getAllFileV2Query      = "SELECT ed.id, ed.filename, ed.category, ed.num_of_failed, ed.num_of_success, ed.status, edi.filename as invalid_filename, edi.is_valid, ed.created_at, ed.updated_at, ed.deleted_at FROM public.excel_document ed LEFT JOIN public.excel_document_invalid edi ON ed.id = edi.id WHERE ed.deleted_at is NULL"
	getInvalidFileQuery    = "SELECT id, excel_document_id, filename, is_valid, created_at, updated_at, deleted_at FROM public.excel_document_invalid WHERE id = $1"
	getAllInvalidFileQuery = "SELECT id, excel_document_id, filename, is_valid, created_at, updated_at, deleted_at FROM public.excel_document_invalid"
)

func (r *Repository) UploadFile(insertFile models.FileExcelDocument) (string, error) {
	stmt, err := r.db.PrepareNamed(insertFileQuery)
	if err != nil {
		return "", errors.New("insert_excel_document" + err.Error())
	}
	defer stmt.Close()

	var uuid string
	err = stmt.QueryRow(&insertFile).Scan(&uuid)
	if err != nil {
		return "", errors.New("insert_excel_document" + err.Error())
	}
	return uuid, nil
}

func (r *Repository) GetAllFile() ([]*models.FileExcelDocument, error) {
	var files []*models.FileExcelDocument

	// err := r.db.Select(&files, getAllFileQuery)
	err := r.db.Select(&files, getAllFileV2Query)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func (r *Repository) GetFile(id string) (*models.FileExcelDocument, error) {
	var files []*models.FileExcelDocument

	err := r.db.Select(&files, getFileQuery, id)
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, nil
	}

	return files[0], nil
}

// on testing
func (r *Repository) GetAllFileWithFilter(filters []*dto.FileFilter) ([]*models.FileExcelDocument, error) {
	var files []*models.FileExcelDocument

	filterQuery := ""

	for _, filter := range filters {
		switch filter.Key {
		case "id":
			filterQuery = fmt.Sprintf("%s AND public.excel_document.id = '%s'", filterQuery, filter.Value)
		case "status":
			filterQuery = fmt.Sprintf("%s AND public.excel_document.status = '%s'", filterQuery, filter.Value)
		case "upload_at":
			filterQuery = fmt.Sprintf("%s AND DATE(public.excel_document.created_at) = DATE('%s')", filterQuery, filter.Value)
		}
	}

	finalQuery := fmt.Sprintf("%s%s", getAllFileQuery, filterQuery)

	err := r.db.Select(&files, finalQuery)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func (r *Repository) GetAllInvalidFile() ([]*models.InvalidFileExcelDocument, error) {
	var files []*models.InvalidFileExcelDocument

	err := r.db.Select(&files, getAllInvalidFileQuery)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func (r *Repository) GetInvalidFile(id string) (*models.InvalidFileExcelDocument, error) {
	var files []*models.InvalidFileExcelDocument

	err := r.db.Select(&files, getInvalidFileQuery, id)
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, nil
	}

	return files[0], nil
}

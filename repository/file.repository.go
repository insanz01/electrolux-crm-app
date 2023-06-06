package repository

import (
	"errors"

	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models"
)

type FileRepository interface {
	UploadFile(filename string) (string, error)
}

const (
	insertFileQuery = "INSERT INTO excel_document (filename, category, num_of_failed, num_of_success, status) VALUES (:filename, :category, :num_of_failed, :num_of_success, :status) returning id"
	getFileQuery    = "SELECT id, filename, category, num_of_failed, num_of_success, status, created_at, updated_at, deleted_at FROM public.excel_document WHERE id = $1"
	getAllFileQuery = "SELECT id, filename, category, num_of_failed, num_of_success, status, created_at, updated_at, deleted_at FROM public.excel_document"
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

	err := r.db.Select(&files, getAllFileQuery)
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

	if len(files) < 1 {
		return nil, nil
	}

	return files[0], nil
}

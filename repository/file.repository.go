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

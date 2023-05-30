package models

type FileExcelDocument struct {
	Filename     string `db:"filename"`
	Category     string `db:"category"`
	NumOfFailed  int    `db:"num_of_failed"`
	NumOfSuccess int    `db:"num_of_success"`
	Status       string `db:"status"`
}

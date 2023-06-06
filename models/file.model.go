package models

type FileExcelDocument struct {
	Id           string  `db:"id"`
	Filename     string  `db:"filename"`
	Category     string  `db:"category"`
	NumOfFailed  int     `db:"num_of_failed"`
	NumOfSuccess int     `db:"num_of_success"`
	Status       string  `db:"status"`
	CreatedAt    string  `db:"created_at"`
	UpdatedAt    string  `db:"updated_at"`
	DeletedAt    *string `db:"deleted_at"`
}

type InvalidFileExcelDocument struct {
	Id              string  `db:"id"`
	ExcelDocumentId string  `db:"excel_document_id"`
	Filename        string  `db:"filename"`
	IsValid         bool    `db:"is_valid"`
	CreatedAt       string  `db:"created_at"`
	UpdatedAt       string  `db:"updated_at"`
	DeletedAt       *string `db:"deleted_at"`
}

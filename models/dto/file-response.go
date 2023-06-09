package dto

type InvalidFile struct {
	Filename *string `json:"filename"`
	IsValid  *bool   `json:"is_valid"`
	FilePath string  `json:"file_path"`
}

type FileResponse struct {
	UUID         string       `json:"uuid"`
	Filename     string       `json:"filename"`
	NumOfFailed  int          `json:"num_of_failed"`
	NumOfSuccess int          `json:"num_of_success"`
	Status       string       `json:"status"`
	Category     string       `json:"category"`
	InvalidFile  *InvalidFile `json:"invalid_file"`
	FilePath     string       `json:"file_path"`
	UpdatedAt    string       `json:"updated_at"`
}

type InvalidFileResponse struct {
	Id              string `json:"id"`
	ExcelDocumentId string `json:"excel_document_id"`
	Filename        string `json:"filename"`
	IsValid         bool   `json:"is_valid"`
	FilePath        string `json:"file_path"`
	UpdatedAt       string `json:"updated_at"`
}

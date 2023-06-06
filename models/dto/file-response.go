package dto

type FileResponse struct {
	UUID      string `json:"uuid"`
	Filename  string `json:"filename"`
	Status    string `json:"status"`
	Category  string `json:"category"`
	FilePath  string `json:"file_path"`
	UpdatedAt string `json:"updated_at"`
}

type InvalidFileResponse struct {
	Id              string `json:"id"`
	ExcelDocumentId string `json:"excel_document_id"`
	Filename        string `json:"filename"`
	IsValid         bool   `json:"is_valid"`
	FilePath        string `json:"file_path"`
	UpdatedAt       string `json:"updated_at"`
}

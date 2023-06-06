package dto

type FileResponse struct {
	UUID      string `json:"uuid"`
	Filename  string `json:"filename"`
	Status    string `json:"status"`
	Category  string `json:"category"`
	FilePath  string `json:"file_path"`
	UpdatedAt string `json:"updated_at"`
}

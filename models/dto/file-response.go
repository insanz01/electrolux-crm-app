package dto

type FileResponse struct {
	UUID     string `json:"uuid"`
	Filename string `json:"filename"`
	Status   string `json:"status"`
}

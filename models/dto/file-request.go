package dto

import "mime/multipart"

type FileRequest struct {
	Category   string                `form:"category"`
	File       *multipart.FileHeader `form:"file"`
	DivisionId string                `form:"division_id"`
}

type FileFilter struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type FileFilterRequest struct {
	Filters []*FileFilter `json:"filters"`
}

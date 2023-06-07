package dto

import "mime/multipart"

type FileRequest struct {
	Category string                `form:"category"`
	File     *multipart.FileHeader `form:"file"`
}

type FileFilter struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type FileFilterRequest struct {
	Filters []*FileFilter `json:"filters"`
}

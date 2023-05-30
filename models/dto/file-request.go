package dto

import "mime/multipart"

type FileRequest struct {
	Category string                `form:"category"`
	File     *multipart.FileHeader `form:"file"`
}

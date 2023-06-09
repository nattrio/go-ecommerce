package files

import "mime/multipart"

type FileReq struct {
	File        *multipart.FileHeader `form:"file"`
	Destination string                `form:"destination"`
	Extension   string
	FileName    string
}

type FileRes struct {
	Filename string `json:"filename"`
	Url      string `json:"url"`
}

type DeleteFileReq struct {
	Destination string `json:"destination"`
}

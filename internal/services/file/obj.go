package file_service

type fileObject struct {
	Path string `json:"path"`
	Name string `json:"name"`
}

type dirObject struct {
	Path string `json:"path"`
}

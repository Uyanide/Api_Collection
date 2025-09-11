package file_service

import "github.com/gin-gonic/gin"

func (s *FileService) setupRoutes(r *gin.Engine) {
	for urlPath := range s.fileMap {
		path := urlPath
		r.GET(path, func(c *gin.Context) { s.serveFile(c, path) })
	}
}

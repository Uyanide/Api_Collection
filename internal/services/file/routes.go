package file_service

import "github.com/gin-gonic/gin"

func (s *FileService) setupRoutes(r *gin.Engine) {
	// Single files
	for urlPath := range s.fileMap {
		path := urlPath
		r.GET(path, func(c *gin.Context) { s.serveFile(c, path) })
	}

	// Directories
	for urlPath := range s.dirMap {
		path := urlPath
		r.GET(path, func(c *gin.Context) { s.serveDirFile(c, path) })
		g := r.Group(path)
		g.GET("/*filepath", func(c *gin.Context) { s.serveDirFile(c, path) })
	}
}

package services

import (
	"github.com/gin-gonic/gin"
)

type GeneralService interface {
	Init(*gin.Engine)
}

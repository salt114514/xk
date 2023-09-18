package midwares

import (
	"STU/app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleNotFound(c *gin.Context) {
	utils.Response(c, 404, 200404, http.StatusText(http.StatusNotFound), nil)
}

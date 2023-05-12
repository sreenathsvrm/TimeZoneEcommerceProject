package utilhandler

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)



func GetAdminIdFromContext(c *gin.Context) (int, error) {
	id := c.Value("adminId")
	adminId, err := strconv.Atoi(fmt.Sprintf("%v", id))
	return adminId, err
}


func GetUserIdFromContext(c *gin.Context) (int, error) {
	id := c.Value("userId")
	userId, err := strconv.Atoi(fmt.Sprintf("%v", id))
	return userId, err
}
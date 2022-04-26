package middleWare

import (
	"github.com/gin-gonic/gin"
	"github.com/jtao539/autocode/model/response"
	"strconv"
)

func IdFilter(c *gin.Context) {
	ids := c.Param("id")
	_, err := strconv.Atoi(ids)
	if err == nil {
		c.Next()
		return
	}
	response.FailByIdValid(c)
	c.Abort()
}

func containArray(a int, args []int) bool {
	for i := 0; i < len(args); i++ {
		if a == args[i] {
			return true
		}
	}
	return false
}

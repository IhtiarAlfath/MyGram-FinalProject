package helpers

import (
	"kominfo/h8-myGram-finalproject/pkg/errs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetParamID(c *gin.Context, key string) (int, errs.MessageErr) {
	value := c.Param(key)

	id, err := strconv.Atoi(value)

	if err != nil {
		return 0, errs.NewBadRequest("invalid parameter id")
	}

	return id, nil
}

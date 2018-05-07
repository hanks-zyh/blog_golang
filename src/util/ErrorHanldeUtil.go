package util

import (
	"../model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func CheckNil(c *gin.Context, field interface{}) {
	if field == nil || field == "" {
		res := model.Result{
			StatusCode: http.StatusBadRequest,
			Code:       http.StatusBadRequest,
			Msg:        "field is nil",
			Data:       nil,
		}
		bytes, _ := json.Marshal(res)
		c.JSON(res.StatusCode, string(bytes))
		panic("field is nil")
	}
}

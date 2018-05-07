package controller

import (
	"../model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendResult(c *gin.Context, err error, data interface{}) {
	var res model.Result
	if err == nil {
		res = model.Result{
			StatusCode: http.StatusOK,
			Code:       http.StatusOK,
			Msg:        "ok",
			Data:       data,
		}
	} else {
		res = model.Result{
			StatusCode: http.StatusBadRequest,
			Code:       http.StatusBadRequest,
			Msg:        err.Error(),
			Data:       nil,
		}
	}
	bytes, _ := json.Marshal(res)
	c.JSON(res.StatusCode, string(bytes))
}

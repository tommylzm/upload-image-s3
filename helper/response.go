package helper

import (
	"encoding/json"
	"net/http"
	"reflect"
	"upload-image-s3/errorcode"

	"github.com/gin-gonic/gin"
)

// Response : api 回應格式
type Response struct {
	Code    string      `json:"Code"`
	Message string      `json:"Message,omitempty"`
	Result  interface{} `json:"Result"`
}

func init() {
	ResJSON = resJSON
}

// ResJSON : RESTful API統一回應錯誤處理
var ResJSON func(c *gin.Context, res interface{}, err error)

func resJSON(c *gin.Context, res interface{}, err error) {
	output := &Response{Code: "1", Result: res}
	if err == nil {
		if output.Result == nil {
			output.Result = make([]string, 0)
		} else if v := reflect.ValueOf(output.Result); v.Kind() == reflect.Ptr && v.IsNil() {
			output.Result = []string{}
		} else if reflect.ValueOf(output.Result).Kind() == reflect.Slice {
			if reflect.ValueOf(output.Result).IsNil() {
				output.Result = []string{}
			}
		}

		c.JSON(http.StatusOK, output)
		return
	}

	output.Result = []string{}
	c.Error(err)

	var httpcode int
	if ce, ok := err.(*errorcode.CustomError); ok {
		httpcode, output.Code, output.Message = ce.Decode()
	} else {
		httpcode = http.StatusInternalServerError
		output.Code = errorcode.ExceptionError
		output.Message = err.Error()
	}

	c.AbortWithStatusJSON(httpcode, output)
	return
}

func resJSONWithDebug(c *gin.Context, res interface{}, err error) {
	output := &Response{Code: "1", Result: res}
	if err == nil {
		raw, _ := json.Marshal(res)
		println(string(raw))
		c.JSON(http.StatusOK, output)
		return
	}

	c.Error(err)

	var httpcode int
	if ce, ok := err.(*errorcode.CustomError); ok {
		httpcode, output.Code, output.Message = ce.Decode()
	} else {
		httpcode = http.StatusInternalServerError
		output.Code = errorcode.ExceptionError
		output.Message = err.Error()
	}

	c.AbortWithStatusJSON(httpcode, output)
	return
}

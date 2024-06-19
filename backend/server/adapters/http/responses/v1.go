package responses

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Total   *int        `json:"total,omitempty"`
	HasNext *bool       `json:"has_next,omitempty"`
	Data    interface{} `json:"data"`
}

func HandleSuccess(ctx *gin.Context, data interface{}, code int) {
	offsetParams := ctx.Request.Form.Get("offset")
	if offsetParams == "" {
		offsetParams = "0"
	}
	offset, _ := strconv.Atoi(offsetParams)

	if data == nil {
		data = map[string]interface{}{}
	}
	hasNext := (offset != 0)

	length, err := LengthOfInterfaceData(data)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Length of data stored in interface: %d\n", length)
	}

	resp := Response{Code: errorCodeMap[ErrSuccess], Message: ErrSuccess.Error(), Data: data, Total: &length, HasNext: &hasNext}
	if _, ok := errorCodeMap[ErrSuccess]; !ok {
		resp = Response{Code: 0, Message: "", Data: data, Total: &length, HasNext: &hasNext}
	}
	ctx.JSON(code, resp)
}

func HandleError(ctx *gin.Context, httpCode int, err error, data interface{}) {
	if data == nil {
		data = map[string]string{}
	}

	resp := Response{Code: errorCodeMap[err], Message: err.Error(), Data: data}
	if _, ok := errorCodeMap[ErrSuccess]; !ok {
		resp = Response{Code: 500, Message: "unknown error", Data: data}
	}
	ctx.JSON(httpCode, resp)
}

type Error struct {
	Code    int
	Message string
}

var errorCodeMap = map[error]int{}

func newError(code int, msg string) error {
	err := errors.New(msg)
	errorCodeMap[err] = code
	return err
}
func (e Error) Error() string {
	return e.Message
}
func LengthOfInterfaceData(i interface{}) (int, error) {
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Slice, reflect.Array, reflect.String:
		return v.Len(), nil
	case reflect.Map:
		return len(v.MapKeys()), nil
	case reflect.Chan:
		return v.Len(), nil
	default:
		return 0, fmt.Errorf("unsupported type: %s", v.Kind())
	}
}

// author: jiazujiang
// date: 2023/6/14
package render

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ————————————————————————————岁月史书——————————————————————————————
var (
	statusSuccess             int = http.StatusOK                  //HTTP.StatusOK = 200
	statusOK                  int = http.StatusOK                  //正确
	statusBadRequest          int = http.StatusBadRequest          //错误请求
	statusUnauthorized        int = http.StatusUnauthorized        //鉴权失败
	statusForbidden           int = http.StatusForbidden           //禁止访问
	statusNotFound            int = http.StatusNotFound            //资源不存在
	statusMethodNotAllowed    int = http.StatusMethodNotAllowed    //方法不允许
	statusInternalServerError int = http.StatusInternalServerError //服务器内部错误
)

func init() {
	statusOK = 1
	statusBadRequest = 0
	statusUnauthorized = -1

}

type renderDataResp struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func ErrBadRequest(ctx *gin.Context, msg string) {
	ctx.JSON(statusSuccess, &renderDataResp{
		Code: statusBadRequest,
		Msg:  msg,
		Data: nil})
}

func ErrCustom(ctx *gin.Context, code int, msg string) {
	ctx.JSON(statusSuccess, &renderDataResp{
		Code: code,
		Msg:  msg,
		Data: nil})
}

func ErrUnauthorized(ctx *gin.Context, msg string) {
	ctx.JSON(statusSuccess, &renderDataResp{
		Code: statusUnauthorized,
		Msg:  msg,
		Data: nil})
}

func ErrForbidden(ctx *gin.Context, msg string) {
	ctx.JSON(statusSuccess, &renderDataResp{
		Code: statusForbidden,
		Msg:  msg,
		Data: nil})
}

func ErrInternalServerError(ctx *gin.Context, msg string) {
	ctx.JSON(statusSuccess, &renderDataResp{
		Code: statusInternalServerError,
		Msg:  msg,
		Data: nil})
}

func ErrNotFound(ctx *gin.Context, msg string) {
	ctx.JSON(statusSuccess, &renderDataResp{
		Code: statusNotFound,
		Msg:  msg,
		Data: nil})
}

func ErrMethodNotAllowed(ctx *gin.Context, msg string) {
	ctx.JSON(statusSuccess, &renderDataResp{
		Code: statusMethodNotAllowed,
		Msg:  msg,
		Data: nil})
}

func Success(ctx *gin.Context, msg string) {
	ctx.JSON(statusSuccess, &renderDataResp{
		Code: statusOK,
		Msg:  msg,
		Data: nil})
}

func Data(ctx *gin.Context, data interface{}, msg string) {
	ctx.JSON(statusSuccess, &renderDataResp{
		Code: statusOK,
		Data: data,
		Msg:  msg})
}

//————————————————————————————统一错误处理——————————————————————————————

type ErrType int

// 公用码
const (
	httpStatusOk = http.StatusOK

	UnKnowErrorCode     ErrType = -2
	Unauthorized        ErrType = -1
	BadRequest          ErrType = 0
	StatusOk            ErrType = 1
	BindError           ErrType = 2
	ValidateError       ErrType = 3
	Forbidden           ErrType = http.StatusForbidden
	NotFound            ErrType = http.StatusNotFound
	MethodNotAllowed    ErrType = http.StatusMethodNotAllowed
	InternalServerError ErrType = http.StatusInternalServerError
)

var message = map[ErrType]string{
	UnKnowErrorCode:     "未知错误",
	Unauthorized:        "鉴权失败,无权访问!",
	BadRequest:          "错误请求!",
	StatusOk:            "成功",
	BindError:           "参数绑定错误",
	ValidateError:       "参数验证错误",
	Forbidden:           "禁止访问",
	NotFound:            "资源不存在",
	MethodNotAllowed:    "方法不允许",
	InternalServerError: "服务器内部错误",
}

var rpcMessage = map[string]ErrType{}

func (s ErrType) Error() string {
	return message[s]
}

func (s ErrType) Code() int {
	return int(s)
}

func Err(ctx *gin.Context, err error) {
	v, ok := err.(ErrType)
	var errStr string
	if !ok {
		v = UnKnowErrorCode
		errStr = err.Error()
	} else {
		errStr = v.Error()
	}

	ctx.JSON(httpStatusOk, &renderDataResp{
		Code: v.Code(),
		Data: nil,
		Msg:  errStr})
}

func ErrRpc(ctx *gin.Context, errMsg string) {
	var errStr string
	v, ok := rpcMessage[errMsg]
	if !ok {
		v = UnKnowErrorCode
		errStr = errMsg
	} else {
		errStr = v.Error()
	}
	ctx.JSON(httpStatusOk, &renderDataResp{
		Code: v.Code(),
		Data: nil,
		Msg:  errStr})
}

func ErrWithData(ctx *gin.Context, data interface{}, err error) {
	v, ok := err.(ErrType)
	var errStr string
	if !ok {
		v = UnKnowErrorCode
		errStr = err.Error()
	} else {
		errStr = v.Error()
	}

	ctx.JSON(httpStatusOk, &renderDataResp{
		Code: v.Code(),
		Data: data,
		Msg:  errStr})
}

func Ok(ctx *gin.Context, data interface{}) {
	ctx.JSON(httpStatusOk, &renderDataResp{
		Code: StatusOk.Code(),
		Data: data,
		Msg:  StatusOk.Error()})
}

func InjErr(m map[ErrType]string) {
	for k, v := range m {
		message[k] = v
		rpcMessage[v] = k
	}
}

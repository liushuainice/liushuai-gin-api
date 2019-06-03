package util

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CodeOK 成功
const CodeOK int = 0

// BadRequest 错误请求
func BadRequest(c *gin.Context) {
	c.AbortWithStatus(400)
}

// Failure 请求失败
func Failure(c *gin.Context, status int) {
	// 自定义错误码
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": status,
	})
}

// FailureWithData 请求失败
func FailureWithData(c *gin.Context, status int, h gin.H) {
	_, ok := h["code"]
	if ok {
		panic(errors.New("code is keyword"))
	}
	h["code"] = status
	c.AbortWithStatusJSON(http.StatusOK, h)
}

// Success 请求成功
func Success(c *gin.Context, v interface{}) {
	h := gin.H{}
	if v != nil {
		// 都解析成MAP结构
		data, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}

		if err := json.Unmarshal(data, &h); err != nil {
			panic(err)
		}
		_, ok := h["code"]
		if ok {
			panic(errors.New("code is keyword"))
		}
	}
	// 注入状态码code
	h["code"] = CodeOK
	c.JSON(http.StatusOK, h)
}

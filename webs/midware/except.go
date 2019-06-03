package midware

import (
	"encoding/json"
	"errors"
	"net/http"
	"runtime"
	"strings"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
	"liushuai-gin-api/util/golog"
)

// RecoverPanic panic catch and process method for martini
func RecoverPanic() gin.HandlerFunc {
	// var err error
	return func(c *gin.Context) {
		defer func() {
			r := recover()
			if r != nil {
				var err error
				switch t := r.(type) {
				case int:
					// 通用错误，如:401 Unauthorized
					c.AbortWithStatus(t)
					return
				case string:
					err = errors.New(t)
				case error:
					err = t

				default:
					err = errors.New("Unknown error")
				}

				if err == gorm.ErrRecordNotFound || err == redis.Nil {
					// 未找到资源错误，返回400
					c.Status(http.StatusNotFound)
				} else {
					// 服务器异常，返回500错误
					buf := make([]byte, 1<<16)
					size := runtime.Stack(buf, false)
					text := string(buf[:size])
					slice := strings.Split(text, "\n")

					shortTrace := ""

					for i := 5; i < len(slice); i++ {
						line := slice[i]
						if strings.Contains(line, "/golog") {
							continue
						}
						if strings.HasPrefix(line, "reflect.") {
							break
						}
						shortTrace += (strings.Repeat(" ", 28) + line + "\n")
					}

					log.Errorf("PANIC: %s\n%s", err, shortTrace)
					str, _ := json.Marshal(map[string]string{"success": "-1", "error": "InternalServerError"})
					c.JSON(http.StatusInternalServerError, string(str))
				}
			}
		}()
		c.Next()
	}
}

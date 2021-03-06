/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/1/27 10:08
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Authorize(ctx *gin.Context) {
	auth := ctx.GetHeader("Authorize")
	// įŽåé´æ
	if auth != viper.GetString("authorize") {
		ctx.Abort()
		return
	}
	ctx.Next()
}

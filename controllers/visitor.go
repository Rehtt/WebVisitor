/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/1/26 10:59
 */

package controllers

import (
	"WebVisitor/extends/mysql"
	"WebVisitor/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetVisitorInfo(ctx *gin.Context) {
	info := &mysql.Visitor{}
	ip, _ := ctx.RemoteIP()
	info.IP = ip.String()
	if ctx.GetHeader("X-Forwarded-For") != "" {
		info.IP = ctx.GetHeader("X-Forwarded-For")
	}
	info.UserAgent = ctx.GetHeader("User-Agent")
	frequency := models.VisitorInfo(info)
	ctx.Writer.WriteString(fmt.Sprintf("%d", frequency))
}

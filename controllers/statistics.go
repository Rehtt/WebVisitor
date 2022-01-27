/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/1/27 10:07
 */

package controllers

import (
	"WebVisitor/models"
	"github.com/gin-gonic/gin"
	"time"
)

func GetStatistics(ctx *gin.Context) {
	startTime, _ := time.Parse("2006-01-02T15:04:05Z", ctx.Query("sTime"))
	endTime, _ := time.Parse("2006-01-02T15:04:05Z", ctx.Query("eTime"))
	out := models.GetStatistics(&startTime, &endTime)
	ctx.Writer.Write(out)
}

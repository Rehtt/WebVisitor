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
	var startTime, endTime *time.Time
	start, err := time.Parse("2006-01-02T15:04:05Z", ctx.Query("sTime"))
	if err == nil {
		startTime = &start
	}
	end, err := time.Parse("2006-01-02T15:04:05Z", ctx.Query("eTime"))
	if err == nil {
		endTime = &end
	}
	out := models.GetStatistics(startTime, endTime)
	ctx.Writer.Write(out)
}

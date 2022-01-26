/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/1/26 11:22
 */

package models

import (
	"WebVisitor/extends/ip"
	"WebVisitor/extends/mysql"
	"fmt"
	"log"
	"time"
)

var visitorMap = map[string]time.Time{}
var frequency int64
var timeout = 5 * time.Minute
var cleanTime = 10 * time.Minute
var vInit = false

func VisitorInfo(info *mysql.Visitor) int64 {
	if !vInit {
		visitorInit()
	}
	now := time.Now()
	if v, ok := visitorMap[info.IP]; ok && now.Sub(v) < timeout {
		return frequency
	}

	var visitor mysql.Visitor
	mysql.DBQuery(&visitor, map[string]interface{}{"ip": info.IP}).GetContent()
	visitor.Frequency++
	if visitor.Model == nil {
		ipinfo, err := ip.IP.GetIpInfo(info.IP)
		if err != nil {
			log.Println(err)
			return frequency
		}
		visitor.IP = info.IP
		visitor.UserAgent = info.UserAgent
		visitor.Country = ipinfo.Country
		visitor.CountryId = ipinfo.CountryId
		visitor.Province = ipinfo.Region
		visitor.City = ipinfo.City
		visitor.ISP = ipinfo.ISP
	}
	mysql.DBUpdate(&visitor).Save()
	frequency++

	return frequency
}

func visitorInit() {
	// 获取总访问量
	mysql.DB.Self.Model(&mysql.Visitor{}).Select("coalesce(sum(frequency),0) as frequency").Find(&frequency)

	// 获取最近访问者
	var visitors []mysql.Visitor
	mysql.DBQuery(&visitors, fmt.Sprintf("updated_at > '%s'", time.Now().Add(-timeout).Format("2006-01-02 15:04:05")))
	for _, v := range visitors {
		visitorMap[v.IP] = v.UpdatedAt
	}

	// 定时清除
	go func() {
		t := time.NewTicker(cleanTime)
		for {
			<-t.C
			cleanVisitorMap()
		}
	}()
	vInit = true
}
func cleanVisitorMap() {
	for k, v := range visitorMap {
		if time.Now().Sub(v) > timeout {
			delete(visitorMap, k)
		}
	}
}

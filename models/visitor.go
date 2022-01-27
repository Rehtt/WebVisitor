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
	"sync"
	"time"
)

var visitorMap = sync.Map{}
var timeout = 5 * time.Minute
var cleanTime = 10 * time.Minute

type visitorStruct struct {
	visitor   *sync.Map
	frequency int64
}

func VisitorInfo(info *mysql.Visitor) int64 {
	now := time.Now()
	visitorMapPtr, _ := visitorMap.LoadOrStore(info.Host, visitorInit(info.Host))

	frequency := visitorMapPtr.(*visitorStruct).frequency
	if mm, ook := visitorMapPtr.(*visitorStruct).visitor.Load(info.IP); ook {
		if now.Sub(mm.(time.Time)) < timeout {
			return frequency
		}
	}

	var visitor mysql.Visitor
	mysql.DBQuery(&visitor, map[string]interface{}{"ip": info.IP, "host": info.Host}).GetContent()
	visitor.Frequency++
	if visitor.Model == nil {
		ipinfo, err := ip.IP.GetIpInfo(info.IP)
		if err != nil {
			log.Println(err)
			//return frequency
		}
		visitor.IP = info.IP
		visitor.UserAgent = info.UserAgent
		visitor.Country = ipinfo.Country
		visitor.CountryId = ipinfo.CountryId
		visitor.Province = ipinfo.Region
		visitor.City = ipinfo.City
		visitor.ISP = ipinfo.ISP
		visitor.Host = info.Host
	}
	mysql.DBUpdate(&visitor).Save()
	frequency++
	visitorMapPtr.(*visitorStruct).visitor.Store(info.IP, now)
	visitorMapPtr.(*visitorStruct).frequency = frequency
	return frequency
}

func visitorInit(host string) *visitorStruct {
	visitor := &visitorStruct{
		visitor:   &sync.Map{},
		frequency: 0,
	}
	// 获取总访问量
	mysql.DB.Self.Model(&mysql.Visitor{}).Where(map[string]interface{}{"host": host}).Select("coalesce(sum(frequency),0) as frequency").Find(&visitor.frequency)

	// 获取最近访问者
	var visitors []mysql.Visitor
	mysql.DBQuery(&visitors, fmt.Sprintf("updated_at > '%s'", time.Now().Add(-timeout).Format("2006-01-02 15:04:05")))
	for _, v := range visitors {
		visitor.visitor.Store(v.IP, v.UpdatedAt)
	}
	return visitor
}
func init() {
	// 定时清除
	go func() {
		t := time.NewTicker(cleanTime)
		for {
			<-t.C
			cleanVisitorMap()
		}
	}()
}
func cleanVisitorMap() {
	now := time.Now()
	visitorMap.Range(func(key, value interface{}) bool {
		value.(*visitorStruct).visitor.Range(func(k, v interface{}) bool {
			if now.Sub(v.(time.Time)) > timeout {
				value.(*visitorStruct).visitor.Delete(k)
			}
			return true
		})
		return true
	})
}

/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/1/27 10:12
 */

package models

import (
	"WebVisitor/extends/mysql"
	"encoding/json"
	"github.com/mileusna/useragent"
	"time"
)

func GetStatistics(startTime, endTime *time.Time) []byte {
	query := mysql.DB.Self.Model(&mysql.Visitor{})
	if startTime != nil {
		query.Where("updated_at > ?", startTime)
	}
	if endTime != nil {
		query.Where("updated_at < ?", endTime)
	}
	var data []mysql.Visitor
	query.Find(&data)

	type botStruct struct {
		URL string `json:"url"`
	}
	type osStruct struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	}
	type outStruct struct {
		mysql.Visitor
		Last    time.Time `json:"last"`
		Name    string    `json:"name"`
		Version string    `json:"version"`
		Device  struct {
			Name string     `json:"name"`
			Bot  *botStruct `json:"bot,omitempty"`
			OS   *osStruct  `json:"os,omitempty"`
		} `json:"device"`
	}

	out := make([]outStruct, 0, len(data))

	for _, v := range data {
		d := outStruct{}
		d.Last = v.UpdatedAt
		v.Model = nil
		d.Visitor = v
		ua := ua.Parse(v.UserAgent)
		d.Name = ua.Name
		d.Version = ua.Version
		if ua.Bot {
			d.Device.Bot = &botStruct{URL: ua.URL}
		} else {
			d.Device.OS = &osStruct{
				Name:    ua.OS,
				Version: ua.OSVersion,
			}
			d.Device.Name = ua.Device
		}
		out = append(out, d)
	}
	o, _ := json.Marshal(out)
	return o
}

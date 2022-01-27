package mysql

import "gorm.io/gorm"

type UserTables struct {
	*gorm.Model
	Name string `gorm:"not null"`
}

type Visitor struct {
	*gorm.Model
	IP        string `json:"ip"`         // ip
	Country   string `json:"country"`    // 国家
	CountryId string `json:"country_id"` // 国家id
	Province  string `json:"province"`   // 省份
	City      string `json:"city"`       // 城市
	ISP       string `json:"isp"`
	UserAgent string `json:"user_agent"`
	Frequency int64  `json:"frequency"` // 访问次数
	Host      string `json:"host"`
}

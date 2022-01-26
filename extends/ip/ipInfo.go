/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/1/26 16:12
 */

package ip

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type IpInfo struct {
	Country   string `json:"country"`
	CountryId string `json:"country_id"`
	Region    string `json:"region"`
	City      string `json:"city"`
	ISP       string `json:"isp"`
}
type Data struct {
	Data *IpInfo `json:"data"`
	Ret  int64   `json:"ret"`
	Msg  string  `json:"msg"`
}

type ip struct {
	appCode string
}

var IP *ip

func (i *ip) Init(appCode string) {
	IP = &ip{appCode: appCode}
}
func (i ip) GetIpInfo(ipp string) (ipInfo IpInfo, err error) {
	req, err := http.NewRequest("GET", "http://api01.aliyun.venuscn.com/ip?ip="+ipp, nil)
	if err != nil {
		return
	}
	req.Header.Add("Authorization", i.appCode)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if resp.StatusCode != 200 && len(body) == 0 {
		fmt.Println(resp)
		return IpInfo{}, fmt.Errorf("请求失败，状态码：%d", resp.StatusCode)
	}

	var data Data
	json.Unmarshal(body, &data)
	if data.Ret != 200 || data.Data == nil {
		return IpInfo{}, fmt.Errorf(data.Msg)
	}
	return *data.Data, nil
}

/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File:  client
 * @Date: 2020/12/15 19:53
 */
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	payload := make(url.Values)
	payload.Add("page", "1")
	payload.Add("page_size", "20")
	payload.Add("limit_status","2")
	payload.Add("start_time","0")
	payload.Add("end_time","0")
	req, err := http.NewRequest(
		http.MethodPost,
		"http://adminbob.pai500.org/merchant/member/list",
		strings.NewReader(payload.Encode()),
	)
	if err != nil {
		fmt.Println("请求失败: ",err)
	}
	req.Header.Add("token", "1078717248806794")
	// 必须要关闭req.Body
	r, err := http.DefaultClient.Do(req)
	defer func() { _ = r.Body.Close() }()
	if err != nil {
		fmt.Println("请求异常: ",err)
	}
	bytes,_ := ioutil.ReadAll(r.Body)
	fmt.Println(string(bytes))
}


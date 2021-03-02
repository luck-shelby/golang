/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File:  No1
 * @Date: 2020/12/15 21:26
 */
package main

// 横向爬取: 以页为单位 纵向爬取：在一个页面中,以不同的条目为单位

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
)
func GetCurrentDirectory() string {
	_, fileStr, _, _ := runtime.Caller(1)
	str := path.Dir(fileStr)
	return strings.Replace(str, "\\", "/", -1) //将\替换成/
}

func httpRequest(url string) (result string,err error) {
	req, errs := http.NewRequest(http.MethodGet, url, nil)
	if errs != nil {
		errs = err
		fmt.Println("请求失败: ",err)
		return "", err
	}
	r, err1 := http.DefaultClient.Do(req)
	if err1 != nil {
		err1 = err
		fmt.Println("请求异常: ",err)
		return "", err
	}
	defer func() { _ = r.Body.Close() }()
	// 循环读取数据
	buf := make([]byte,4096)
	for  {
		n,err := r.Body.Read(buf)
		if n == 0 {
			break
		}
		if err != nil && err != io.EOF {
			fmt.Println("读取出错!")
			return "", err
		}
		result += string(buf[:n])
	}
	return result,nil
}
func SpiderPage(i int,waitFinish chan<- int)  {
	url:= "https://tieba.baidu.com/f?kw=lol&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
	result,err := httpRequest(url)
	if err != nil {
		fmt.Println("httpRequest ERR : ",err)
		return
	}
	// 保存数据
	str := GetCurrentDirectory()
	file,err := os.Create(str +"/"+"第"+strconv.Itoa(i)+"页"+".html")
	if err != nil {
		fmt.Println("CreateFile ERR : ",err)
		return
	}
	defer file.Close()
	_, _ = file.WriteString(result)
	waitFinish <- i
}
func handle(start,end int)  {
	waitFinish := make(chan int)
	fmt.Printf("正在爬取第%d页到第%d页的数据\n ",start,end)
	for i:=start;i<=end;i++{
		go SpiderPage(i,waitFinish)
	}
	for i:=start;i<=end;i++{
		fmt.Printf("第%d页爬取完成\n ",<-waitFinish)
	}
}
func main() {
	var start,end int
	fmt.Print("起始位置:")
	fmt.Scan(&start)
	fmt.Print("结束位置:")
	fmt.Scan(&end)
	handle(start,end)
}

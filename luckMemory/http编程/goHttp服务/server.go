/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File:  server
 * @Date: 2020/12/15 17:58
 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

func handleRequest(resp http.ResponseWriter,fileName string)  {
	path := "D:/Luck/Luckdays" + fileName

	file,err := os.OpenFile(path,os.O_RDWR,0666)
	if err != nil {
		fmt.Println("打开文件失败:",err)
		_, _ = resp.Write([]byte("file does not exist"))
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for  {
		content,err := reader.ReadBytes('\n')
		// content,err := reader.ReadString('\n')
		if err == io.EOF || err != nil{
			break
		}
		_, _ = resp.Write(content)
	}
}
func handler(resp http.ResponseWriter, req *http.Request)  {
	requestUrl := req.URL.String()
	handleRequest(resp,requestUrl)
}
func main() {
	http.HandleFunc("/",handler)
	err := http.ListenAndServe("127.0.0.1:8809",nil)
	if err != nil {
		fmt.Println("ListenAndServe err:",err)
		return
	}
}



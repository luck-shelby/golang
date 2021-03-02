/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File:  命令行参数
 * @Date: 2020/12/11 16:38
 */
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	argList  := os.Args  // 获取命令行参数
	if len(argList) != 2{
		fmt.Println("参数格式错误! ")
		return
	}
	// 获取文件属性
	/*
	filePath := argList[1]
	fileInfo ,err := os.Stat(filePath)
	if err != nil {
		fmt.Println("获取文件属性错误! ")
		return
	}
	*/

	err := filepath.Walk(argList[1], func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Printf("文件名: %v, 文件大小: %v\n",info.Name(),info.Size())
		return nil
	})
	if err != nil {
		fmt.Println("获取文件属性错误! ")
	}
}

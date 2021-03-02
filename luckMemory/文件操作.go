/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File:  文件操作
 * @Date: 2020/12/2 17:23
 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

/*
	创建文件:
		1. create(name): 可读可写模式创建文件,根据提供的文件名创建文件,如果文件存在则清空之前文件内容
		2. open(name): 只能是只读权限打开文件,
		3. openFile(name,flg，perm): 以读写权限打开 ,flg:权限,perm: 权限范围,只有在Linux中有效
			flg:
				O_RDONLY(只读模式)
				O_WRONLY(只写模式)
				O_RDWR(可读可写)
				0_APPEND(追加模式)
				...
*/
func main4() {
	// 创建
	file,err := os.Create("luckMemory/abc.txt")
	if err != nil {
		fmt.Println("文件创建失败:",err)
		return
	}
	defer file.Close()

	// 打开文件
	file,err = os.Open("luckMemory/abc.txt")  // 只读方式,不能进行写操作
	if err != nil {
		fmt.Println("文件打开失败:",err)
		return
	}
}

func main5() {
	// openFile
	file,err := os.OpenFile("luckMemory/abc.txt",os.O_RDWR|os.O_CREATE,666)
	if err != nil {
		fmt.Println("打开文件失败:",err)
		return
	}
	defer file.Close()
	count := 10
	for i :=0;i<=count;i++ {
		_,err = file.WriteString("-----------what up----------\r\n")
		if err != nil {
			fmt.Println("文件写入失败:",err)
		}
	}

	// 按照指定位置写入  seek(偏移量,偏移起始位置)  返回从文件起始位置到文件读写位置的偏移量
	/*
		参数1: 当为正数时,向右偏移, 当为负数时,向左偏移(文件头部)
		参数2:
			io.SeekStart: 从文件起始位置
			io.SeekEnd: 文件结尾位置
			io.SeekCurrent: 文件当前文位置
	*/
	off,err :=file.Seek(10,io.SeekStart)
	// off,err :=file.Seek(-10,io.SeekEnd)
	// off,err :=file.Seek(10,io.SeekCurrent)
	if err != nil {
		fmt.Println("修改文件位置错误:",err)
	}
	fmt.Println(off)
	n,_ := file.WriteAt([]byte("use seek "),off)  // WriteAt 配合seek写入
	fmt.Println("n = ",n)
}

func main6() {
	// 文件写入其他方式:  简单覆盖式文件写入
	/*
		操作简单一个函数完成数据写入
		新内容覆盖旧的内容
		操作的文件不存在的时候会自动创建
		WriteFile(filename string, data []byte, perm os.FileMode)
	*/
	str := `I already told you before when bao is still a new born baby,`
	err := ioutil.WriteFile("luckMemory/abc.txt",[]byte(str),6)
	if err != nil {
		fmt.Println("写入文件失败")
		return
	}
}

func main7() {
	// 文件写入其他方式:  常规文件写入
	/*
		文件写入灵活 ,对文件的操作更强
		操作流程 : 打开文件(或者创建文件) ,写入内容 ,关闭文件
	*/
	f,err := os.OpenFile("luckMemory/abd.txt", os.O_RDONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("打开文件失败:",err)
		return
	}
	defer f.Close()
	// 字节方式写入
	param := "以语音为物质外壳，由词汇和语法两部分构成的符号系统 \r\n"
	_, err = f.Write([]byte("write : "+param ))
	if err != nil {
		log.Println(err)
		return
	}
	// 字符串写入
	_, err = f.WriteString("writeString : " + param)
	if err != nil {
		log.Println(err)
		return
	}
}

func main8()  {
	// 文件写入其他方式:  带有缓冲区(用户缓冲区)的文件写入
	/*
		先将数据写入缓存区,再由缓冲区写入文件中
		根据设置缓存的大小,可以存储更多数据然后一次写入文件
		数据写入的速度更快一点
	*/
	f,err := os.OpenFile("luckMemory/abe.txt", os.O_RDONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("打开文件失败:",err)
		return
	}
	defer f.Close()
	param := "以语音为物质外壳，由词汇和语法两部分构成的符号系统 \r\n"
	// NewWriter 默认缓冲区大小是 4096
	buf :=bufio.NewWriter(f)
	// 字节写入
	n,_ := buf.Write([]byte("buffer Write : " + param))
	// 字符串写入
	n,_ = buf.WriteString("buffer WriteString : " + param)
	fmt.Println(n)
	// 将缓冲中的数据写入
	err = buf.Flush()
	if err != nil {
		log.Println("flush error :", err)
	}
}

// 读操作
func main9() {
	// （1）按字节读取文件
	f,err := os.Open("luckMemory/abd.txt")
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	buf := make([]byte,1024)
	// 该字节切片用于存放文件所有字节
	var bytes []byte
	for  {
		count,err := f.Read(buf) // 返回 文件中读取到len(b)字节。
		// 检测是否到了文件末尾
		if err == io.EOF {
			break
		}
		// 取出本次读取的数据
		currBytes := buf[:count]
		// 将读取到的数据 追加到字节切片中
		bytes = append(bytes, currBytes...)
	}
	fmt.Println(string(bytes))
}

func main10()  {
	// （2）结合 ioutil 来读取
	f,err := os.Open("luckMemory/abc.txt")
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	b,err:= ioutil.ReadAll(f)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))

	// 一行代码搞定读取任务：
	bytes, _ := ioutil.ReadFile("abd.txt")
	fmt.Println(string(bytes))
}

func main11()  {
	// （3）逐行读取 可以带有缓冲区(用户缓区)的文件读取
	file,err := os.OpenFile("luckMemory/abc.txt",os.O_RDWR,0666)
	if err != nil {
		fmt.Println("打开文件失败:",err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for  {
		// content,err := reader.ReadBytes('\n')
		content,err := reader.ReadString('\n')
		if err == io.EOF || err != nil{
			break
		}
		fmt.Print(content)
	}
}

func main12()  {
	// （4）逐行读取  bufio.Scanner
	file,err := os.OpenFile("luckMemory/abc.txt",os.O_RDWR,0666)
	if err != nil {
		fmt.Println("打开文件失败:",err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var count int
	for scanner.Scan(){
		count ++
		line := scanner.Text()
		fmt.Printf(" 第%d 行: %s\n", count,line)
	}
}

// 文件拷贝 把读取的内容放在缓冲区中,在用write从缓存区写
func main13() {
	fr,err := os.Open("luckMemory/abc.txt")
	if err != nil {
		fmt.Println("打开文件失败:",err)
	}
	defer fr.Close()
	fw,err := os.Create("luckMemory/abu.txt")
	if err != nil {
		fmt.Println("创建文件失败:",err)
	}
	defer fw.Close()
	// 创建一个缓存区
	buf := make([]byte,1024)
	for  {
		n,err:= fr.Read(buf)
		if err == io.EOF || err != nil{
			fmt.Printf("文件读取完成,读取到的字节数:%d",n)
			return
		}
		_, _ = fw.Write(buf[:n]) // 读什么写多少
	}
}


// 文件目录操作
func test1(f *os.File)  {
	// 读取目录 -1代表读取全部
	fileInfo,_ := f.Readdir(-1)
	for _,info := range fileInfo{
		if info.IsDir(){
			fmt.Println(info.Name()," -- 是目录")
		}else {
			fmt.Println(info.Name()," -- 是文件")
		}
	}
}
// 语言递归获取目录下所有文件
func test2(path string)  {
	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir(){
				if strings.HasSuffix(info.Name(),"txt"){
					fmt.Println(info.Name(),"文件大小:", info.Size())
					// 拷贝到指定目录下
					FileCopy("D:/hardship/luckdays/luckMemory"+"/"+info.Name(),"E:/temp/"+info.Name())
				}
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}

func FileCopy(src,dst string)  {
	fr,err := os.Open(src)
	if err != nil {
		fmt.Println("打开文件失败:",err)
	}
	defer fr.Close()
	fw,err := os.Create(dst)
	if err != nil {
		fmt.Println("创建文件失败:",err)
	}
	defer fw.Close()
	// 创建一个缓存区
	buf := make([]byte,1024)
	for  {
		n,err:= fr.Read(buf)
		if err == io.EOF || err != nil{
			fmt.Printf("文件读取完成,读取到的字节数:%d",n)
			return
		}
		_, _ = fw.Write(buf[:n]) // 读什么写多少
	}
}
func main14() {
	// 也是使用openFile(),最后一个参数指定os.ModeDir,返回一个可以读写目录的文件指针
	fmt.Println("输入查询目录路径")
	var path string
	_, _ = fmt.Scan(&path)
	// 打开文件
	f,err := os.OpenFile(path,os.O_RDONLY,os.ModeDir)
	if err != nil {
		fmt.Println("OpenFile err:",err)
		return
	}
	defer f.Close()
	test1(f)
}

func main15() {
	fmt.Println("输入查询目录路径")
	var path string
	_, _ = fmt.Scan(&path)
	test2(path)
}

// exercise  指定一个目录找到txt文件中所有 是 xxx 单词出现的次数
func readFiles(path, s string) int {
	fp,err := os.Open(path)
	if err != nil {
		fmt.Println("Open err :",err)
		return -1
	}
	defer fp.Close()
	// 一行行的读
	reader := bufio.NewReader(fp)
	var word int = 0
	for  {
		buf,err := reader.ReadBytes('\n')
		if err == io.EOF || err != nil {
			fmt.Println("读取完成")
			break
		}
		word += wordCount(string(buf[:]),s)
	}
	return word
}
func wordCount(word,s string)int  {
	arr := strings.Fields(word)
	m := make(map[string]int)
	for i:=0;i<len(arr);i++{
		m[arr[i]]++
	}
	// 统计map中 指定单词的格式
	for k,v := range m{
		if k == s{
			fmt.Printf("%s : %d \n",k,v)
			return m[k]
		}
	}
	return 0  // 没有返回 0
}
func main() {
	fmt.Println("输入查询目录路径")
	var findPath string
	_, _ = fmt.Scan(&findPath)
	var word int = 0
	err := filepath.Walk(findPath, func(path string, info os.FileInfo, err error) error {
			fmt.Println("------------------------------------------------------------------------")
			if err != nil {
				return err
			}
			if !info.IsDir(){
				if strings.HasSuffix(info.Name(),"txt"){
					word += readFiles(findPath+"\\"+ info.Name(),"you")
				}
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("目录所有文件中 xxx 个数为：",word)
}
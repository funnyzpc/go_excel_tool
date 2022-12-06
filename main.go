package main

import (
	"flag"
	"go_excel_tool/util"
)

// build -o go_excel_tool.exe -ldflags "-w -s -X main.env=test" ./main.go
// go build -o ./go_excel_tool.exe -ldflags "-w -s -X main.env=test" ./main.go
// go run -o ./go_excel_tool.exe main.go
func main() {
	dir := flag.String("dir", "", "")
	flag.Parse()
	//fmt.Println("dir=",dir)
	//dir := "D:\\document\\需求文档\\SJ审计需求\\审计业务需求_20220303(to hana)\\202111三大店数据\\韩束202111淘宝订单数据";

	// 获取密码文件
	file, _ := util.DoGetPassFile(*dir)
	// 获取密码数据
	pass_map := util.DoGetPass(file)
	// 获取数据文件(待解密的)
	data_file := util.DoGetDataFile(*dir, file)
	// 解密并输出文件
	util.DoDecFile(data_file, pass_map, *dir+"\\DEC_FILE")

	//str1 := "Welcome to Geeks for Geeks"
	//str2 := "This is the is the article of the Go string is a replacement"
	//
	//res1 := strings.ReplaceAll(str1, " ", "")
	//res2 := strings.ReplaceAll(str2, " ", "")
	//
	//fmt.Println("Result 1: ", res1)
	//fmt.Println("Result 2: ", res2)

}

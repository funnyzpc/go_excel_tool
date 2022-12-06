package util

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
)

func DoGetPass(file string) (map[string]string){
	if ""==file {
		println("错误::读取到密码文件为空或者为目录")
		os.Exit(-1)
		return nil;
	}
	fmt.Println("已读取到密码文件:\n",file)
	// 初始化一个map并附带一个可选的初始bucket（非准确值，只是有提示意义）
	var pass_map = make(map[string]string, 16)
	//file_name, err := filepath.Abs(file)
	//if err != nil{
	//	fmt.Println(err)
	//	return nil
	//}
	f, err := excelize.OpenFile(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
		return nil
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// 读取excel 仅获取第一个sheet
	rows, err := f.GetRows(f.GetSheetName(0))
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
		return nil
	}
	i := 0
	for _, row := range rows {
		//println(row[0],"->",row[1])
		if 0!= i{
			pass_map[row[0]]=row[1]
		}
		i++
	}
	//// 遍历循环 key value
	//for key := range pass_map {
	//	println(key,pass_map[key])
	//}
	return pass_map
}

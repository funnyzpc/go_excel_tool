package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)


/**
* 获取数据文件
 */
func DoGetDataFile(dir string,skip_file string) ([]string){
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1);
	}
	//arr := make([]string,16)
	arr := []string{} // 必须是空定义的，使用make定义为定长容易抛错
	for _, file := range files {
		name := file.Name()
		if strings.HasSuffix(name,".xlsx") && !strings.HasSuffix(skip_file,name){
			arr = append(arr, dir+"\\"+file.Name())
			//return dir+"\\"+file.Name(),nil
		}
	}
	count := len(arr)
	if count==0{
		println("未能获取到任何数据文件")
		os.Exit(-1)
	}
	fmt.Println("已提取数据文件个数:",count)
	for idx,item := range arr{
		fmt.Println(idx,item)
	}
	return arr
}

/**
* 获取密码文件
 */
func DoGetPassFile(dir string) (string, error){
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("未能找到文件:",err)
		os.Exit(-1)
	}
	for _, file := range files {
		name := file.Name()
		//path, _ := filepath.Abs(file.Name())
		//fmt.Println(path,file.Name())
		if strings.HasPrefix(name, "密码") || strings.HasPrefix(name,"pass") || strings.HasPrefix(name,"PASS"){
			//println(file.Sys())
			return dir+"\\"+file.Name(),nil
		}
	}
	fmt.Println("未能找到密码文件[密码*.xlsx、pass*.xlsx、PASS*.xlsx")
	os.Exit(-1)
	return "",nil
}


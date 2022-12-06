

package util

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func DoDecFile(file_array []string,pass_map map[string]string,new_dir string){
	if nil==file_array || nil==pass_map || ""==new_dir{
		fmt.Println("待解密参数为空...")
		os.Exit(-1)
	}
	// 创建目录
	os.RemoveAll(new_dir)
	err:=os.Mkdir(new_dir,os.ModePerm)
	if  err!= nil{
		fmt.Println(err);
		os.Exit(-1);
	}
	fmt.Println("文件保存目录:",new_dir)

	//var ct sync.WaitGroup
	wg := &sync.WaitGroup{}
	start := time.Now() // 获取当前时间
	for key := range pass_map{
		for _,file := range file_array{
			if strings.Contains(file,key){
				wg.Add(1)
				go process(file,pass_map[key],new_dir,wg)
			}
		}
	}
	wg.Wait()
	//sum_time := time.Now().Sub(start);
	//fmt.Println("该函数执行完成耗时：", sum_time.Seconds())
	// 77s -> 21s
	fmt.Println("执行解密耗时：", time.Now().Sub(start).Seconds())
}

func process(file string,pass string,new_dir string,wg *sync.WaitGroup){
	f, err := excelize.OpenFile(file, excelize.Options{Password:pass})
	if err != nil {
		return
	}
	//file_name := filepath.Base(file);
	//file_name2 := strings.ReplaceAll(file_name, " ", "_")
	file_name := strings.ReplaceAll(filepath.Base(file), "：", "")
	file_name = strings.ReplaceAll(file_name, " ", "")
	file_name = strings.ReplaceAll(file_name, "-", "")
	//file_name := strconv.FormatInt(time.Now().Unix(),10)
	err2 := f.SaveAs(new_dir+"\\dec_"+file_name)
	if err2!= nil{
		fmt.Println("process错误1:",err2,file_name)
	}
	defer wg.Done()
	defer func() {
		if err := f.Close(); err != nil{
			fmt.Println("process错误2:",err)
		}
	}()

}


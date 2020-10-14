package basicinfo

import (
	"../nfuncions"
	"fmt"
	"os"
	"strconv"
	"time"
)
var  (outfile = "output.xlsx")
func handleresult(result saveres) {//保存结果
	//	fmt.Println(time.Now().Unix())
	//	fmt.Println(savefilename)
	//enc :=mahonia.NewDecoder("gbk")
	var clearip []string
	for index,infos := range result.Infos{
		clearip = nil
		for _,ip := range infos.Ip {
			if nfuncions.IsContain(ip, clearip) == false {
				clearip = append(clearip, ip)
			}
		}
		result.Infos[index].Ip = clearip
	}
	filename := strconv.FormatInt(time.Now().Unix(),10)+".html"
	savahtml(result,filename)
	fmt.Println("Output File:"+filename)
	//savefile(result)//保存到文件
	//printout(result)//屏幕输出
}
func savahtml(result saveres,filename string){  //保存exlsx
	writeHTML(result,filename)
}
func savefile(result saveres){
	filename := "result.txt"
	//urlfilename := "url.txt"
	file, err := os.OpenFile(filename, os.O_CREATE,0644)
	//fil1, err1 := os.OpenFile(filename, os.O_CREATE,0644)
	if err != nil {
		fmt.Println("Savefile Open fail")
	}
	defer file.Close()
	for _,info := range result.Infos {
		var ips string
		for _,ip := range info.Ip{
			ips = ips+ip+","
		}
		res_string := fmt.Sprintf("Url:%s    Port:%d    Code:%d    Title:%s    Server:%-s     Location:%s    ip:%s",info.Url,info.Port,info.Statuscode,info.Title,info.Server,info.Location,ips)//格式化字符串
		if err == nil {
			file.WriteString(res_string)
			file.WriteString("\n")
		}
	}
	fmt.Printf("Save to file result.txt\n")
	return

}

func printout(result saveres) {
	for _, info := range result.Infos {
		var ips string
		for index, ip := range info.Ip {
			if index == len(info.Ip)-1 {
				ips = ips + ip
			} else {
				ips = ips + ip + ","
			}
		}
		res_string := fmt.Sprintf("Url:%s    Port:%d    Code:%d    Title:%s    Server:%-s     Location:%s    ip:%s", info.Url, info.Port, info.Statuscode, info.Title, info.Server, info.Location, ips) //格式化字符串
		fmt.Println(res_string)
	}
}
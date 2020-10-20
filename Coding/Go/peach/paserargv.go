package main

import (
	"fmt"
	"os"
)
import "flag"

type callback struct {
	method int
	argv map[string]string
	scanmode string
}
var Scanmode string

func parse() *callback{//参数解析
	filepath := flag.String("filepath","","input domainfile path")
	function := flag.Int("func",0,"Select function")
	ports := flag.String("ports","80,443","Input Ports range(default 80,443)")
	firstdict := flag.String("dict1","","first dict.txt path")
	secondict := flag.String("dict2","","second dict.txt path")
	iprange := flag.String("iprange","","input IP range")
	keyword :=  flag.String("keywords","","input keywords")
	var backInfo callback
	resu := make(map[string]string)
	backInfo.argv =resu
	flag.Parse()//解析参数
	switch *function {
	case 1://获取url GET 信息
		if *filepath != ""&& *iprange== "" {//导入域名
			//fmt.Println(*filepath)
			backInfo.method = 1
			backInfo.argv["ports"] = * ports
			backInfo.argv["filepath"] = *filepath
			backInfo.scanmode = "domain"
			return &backInfo
		}
		if *iprange != ""&&*filepath == "" {//获取IP范围
			backInfo.method = 1
			backInfo.argv["ports"] = * ports
			backInfo.argv["iprange"] = *iprange
			backInfo.scanmode = "ip"
			return  &backInfo
		}else{
			fmt.Println("Please input domain.txt or iprange")
			os.Exit(0)
		}
	case 2://关键字字典
		backInfo.method = 2
		backInfo.argv["keywords"] = *keyword
		return &backInfo
	case 3://字典去重
		if firstdict != nil && secondict!=nil{
			backInfo.method = 2
			backInfo.argv["dict1"] = *firstdict
			backInfo.argv["dict2"] = *secondict
			return &backInfo
	}else {
		fmt.Println("Please input dict file path")
		os.Exit(0)
		}
	case 4://参数fuzz
	default:
		fmt.Println("Unkown function select")
		flag.PrintDefaults()
		usage()
		os.Exit(0)
	}
	return nil
}

func usage(){
	usage := "function nums:\n 1.getinfo \n 2.Generate dict \n 3.getdict \n 4.fuzzparams\n\n Ex:\n peach.exe -func 1 -filepath domain.txt (-ports 80,443,...)\n peach.exe -func 1 -iprange 192.168.1.1-255\n peach.exe -func 2 -keywords baidu,tenxun"
	fmt.Println(usage)
}
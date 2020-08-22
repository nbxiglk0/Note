package main

import (
	"fmt"
	"os"
)
import "flag"


type callback struct {
	method int
	argv map[string]string
}


func parse() *callback{//参数解析
	filepath := flag.String("filepath","","input file path")
	function := flag.Int("func",0,"Select function")
	firstdict := flag.String("dict1","","first dict.txt path")
	secondict := flag.String("dict2","","second dict.txt path")
	var backInfo callback
	resu := make(map[string]string)
	backInfo.argv =resu
	flag.Parse()//解析参数
	switch *function {
	case 1://获取url GET 信息
		if filepath != nil {
			//fmt.Println(*filepath)
			backInfo.method = 1
			backInfo.argv["filepath"] = *filepath
			return &backInfo
		}else {
			fmt.Println("Please input domain.txt")
			os.Exit(0)
		}
	case 2://字典去重
		if firstdict != nil && secondict!=nil{
			backInfo.method = 2
			backInfo.argv["dict1"] = *firstdict
			backInfo.argv["dict2"] = *secondict
			return &backInfo
	}else {
		fmt.Println("Please input dict file path")
		os.Exit(0)
		}
	case 3://参数fuzz
	default:
		fmt.Println("Unkown function select")
		flag.PrintDefaults()
		usage()
		os.Exit(0)
	}
	return nil
}

func usage(){
	usage := "function nums:\n 1.getinfo \n 2.getdict \n 3.fuzzparams"
	fmt.Println(usage)
}
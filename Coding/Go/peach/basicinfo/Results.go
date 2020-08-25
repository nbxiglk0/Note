package basicinfo

import (
	"../nfuncions"
	"fmt"
	"os"
)

func handleresult(result saveres) {//保存结果
	//	fmt.Println(time.Now().Unix())
	//	fmt.Println(savefilename)
	//enc :=mahonia.NewDecoder("gbk")
	var clearip []string
	for index,infos := range result.infos{
		clearip = nil
		for _,ip := range infos.ip{
			if nfuncions.IsContain(ip,clearip) ==false{
				clearip = append(clearip,ip)
			}
		}
		result.infos[index].ip = clearip
	}
	savefile(result)//保存到文件
	printout(result)//屏幕输出


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
	for _,info := range result.infos {
		var ips string
		for _,ip := range info.ip{
			ips = ips+ip+","
		}
		res_string := fmt.Sprintf("Url:%s    Port:%d    Code:%d    Title:%s    Server:%-s     Location:%s    ip:%s",info.url,info.port,info.Statuscode,info.Title,info.Server,info.Location,ips)//格式化字符串
		if err == nil {
			file.WriteString(res_string)
			file.WriteString("\n")
		}
	}
	fmt.Printf("Save to file result.txt\n")
	return

}

func printout(result saveres) {
	for _, info := range result.infos {
		var ips string
		for index, ip := range info.ip {
			if index == len(info.ip)-1 {
				ips = ips + ip
			} else {
				ips = ips + ip + ","
			}
		}
		res_string := fmt.Sprintf("Url:%s    Port:%d    Code:%d    Title:%s    Server:%-s     Location:%s    ip:%s", info.url, info.port, info.Statuscode, info.Title, info.Server, info.Location, ips) //格式化字符串
		fmt.Println(res_string)
	}
}
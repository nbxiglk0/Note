package basicinfo
//Todo
//1.ip
import (
	"bufio"
	"fmt"
	"os"
)
type respinfo struct {// 信息结构体
	Statuscode int // 状态码
	Server string // Server
	Title string // Ttile
	//Proto string // 协议
	Location string //跳转地址
	Url string //url
	Ip []string  //ip
	Port int //端口
}
type saveres struct {
	Infos []respinfo //每个url的信息结构体
	Len int //访问的url数量
}
type domaininfo struct {
	url string
	port []int
	ip []string
	ipflag int
 }
func Main(argv map[string]string,scanmode string)  {
	if _,ok := argv["filepath"];ok {
		filepath := argv["filepath"]
		ports := argv["ports"]
		domain,_ := getinfo(filepath)
		parseresult := domainHandle(domain,scanmode,ports)
		result := sendrequest(parseresult)
		handleresult(result)
	}
	if _,ok := argv["iprange"];ok{
		iprange := argv["iprange"]
		resultip := Handleip(iprange)
		ports := argv["ports"]
		parseresult := domainHandle(resultip,scanmode,ports)
		result := sendrequest(parseresult)
		handleresult(result)
	}
}
//从文件中读取域名或者ip
//ip 1
//domain 0
func getinfo (filedir string) ([]string, int){//
	var urls  []string//创建str类型的切片
	file, _ := os.Open(filedir)
	if file == nil{
		fmt.Println("Open domain file fail")
		os.Exit(0)
	}
	defer file.Close()//在函数返回时执行
	filescanner := bufio.NewScanner(file)
	filescanner.Split(bufio.ScanWords)
	for filescanner.Scan(){
		urls = append(urls, filescanner.Text())
	}
	return urls,0//返回url切片
}

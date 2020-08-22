package basicinfo
//Todo
//1.乱码
//3.多线程
import (
	"bufio"
	"fmt"
	"os"
)
type respinfo struct {// 信息结构体
	Statuscode int // 状态码
	Server string // Server
	Title string // Ttile
	Proto string // 协议
	Location string //跳转地址
	url string //url
	ip []string  //ip
	port int //端口
}
type saveres struct {
	infos []respinfo //每个url的信息结构体
	len int //访问的url数量
}//编码转化
type domaininfo struct {
	url string
	port []int
	ip []string
}
func Main(filepath string)  {
	domain := getinfo(filepath)
	parsedomain := domainHandle(domain)
	result := sendrequest(parsedomain)
	handleresult(result)

}


//从文件中读取域名
func getinfo (filedir string) []string {//
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
	return urls//返回url切片
}


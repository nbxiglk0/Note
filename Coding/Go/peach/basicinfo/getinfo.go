package basicinfo
//Todo
//1.乱码
//2.结果保存
//3.多线程
import (
	"../nfuncions"
	"bufio"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)
type respinfo struct {// 信息结构体
	Statuscode int // 状态码
	Server string // Server
	Title string // Ttile
	Proto string // 协议
	Location string //跳转地址
	url string //url
	ip []string  //ip
	port int
}
type portsscan struct {
	ip string
	ports []int
}
type saveres struct {
	infos []respinfo //每个url的信息结构体
	len int //访问的url数量
}//编码转化
func Main(filepath string)  {
	urls := getinfo(filepath)
	sendrequest(urls)
}

func getinfo (filedir string) []string {
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
func domain2ip(domain string) *[]string{//解析域名
	ips :=make([]string,0)
	ns, err := net.LookupHost(domain)
	if err != nil{
		return nil
	}else {
		for _, ip := range ns{
			ips = append(ips,ip)
			}
		}
	return &ips
}
func checkport(ip string) portsscan{
	scanports :=[] int {80,443,8080}//扫描的端口
	var results portsscan
	results.ip = ip
	for _,port := range scanports {
		reqip := ip+":"+strconv.Itoa(port)
		timeout := 3*time.Second
		_, err := net.DialTimeout("tcp", reqip,timeout)
		if err == nil{
			results.ports = append(results.ports,port)
		}
	}
	return results
}

func sendrequest(urls []string){
	var saveres saveres
	for url := range urls {//每一个domain
		ip := domain2ip(urls[url])//解析为ip []string
		if ip != nil {
			ip := *ip
			portresult := checkport(ip[0])//多个解析IP只扫描第一个默认IP端口
			if nfuncions.IsContain(80,portresult.ports) {
				requrl := "http://"+urls[url]
				result := request(requrl)
				result.ip = ip
				result.port = 80
				if result != nil {
					saveres.infos = append(saveres.infos, *result)
					saveres.len += 1
				} else {
					break
				}
			}
			if nfuncions.IsContain(443,portresult.ports) {
				requrl := "https://" + urls[url]
				result := request(requrl)
				result.ip = ip
				result.port = 443
				if result != nil {
					saveres.infos = append(saveres.infos, *result)
					saveres.len += 1
				} else {
					break
				}
			}
		}else {//域名未解析到IP
			var nulldoamin respinfo
			nulldoamin.url= urls[url]
			saveres.infos =append(saveres.infos,nulldoamin)
		}
	}
	saveresult(saveres)
}

func request(url string) *respinfo{
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true},}//忽略https证书
	clent := &http.Client{Transport:tr}
	resp, err := clent.Get(url)//发送请求,获取响应
	defer resp.Body.Close()//函数返回时执行
	if err != nil{
		fmt.Println("Get request error")
		return nil
	}	else {
		resps := handleresp(resp)
		resps.url = url
		return resps
	}
}
func handleresp(response *http.Response) *respinfo{
	redirect := []int{302, 307, 301, 308}//跳转状态码
	var result respinfo
	//enc :=mahonia.NewDecoder("gbk")
	body, err := ioutil.ReadAll(response.Body)//读取返回内容
	if err != nil{
		fmt.Println("Get response body error")
		return nil
	}
	convtbody:=string(body)
	reg := regexp.MustCompile(`<title>(.*)</title>`)
	if reg.FindStringSubmatch(convtbody) == nil {
		result.Title = "NULL"
	}else{
		result.Title = reg.FindStringSubmatch(convtbody)[1]//title
	}
	result.Statuscode = response.StatusCode //状态码
	if nfuncions.IsContain(response.StatusCode,redirect){//获取跳转地址
		result.Location = response.Header.Get("Location")
	}else {
		result.Location ="NULL"
	}
	result.Proto = response.Proto //协议
	result.Server = response.Header.Get("Server") //Server
	return &result
}


func saveresult(result saveres) {//保存结果
//	fmt.Println(time.Now().Unix())
//	fmt.Println(savefilename)
	//enc :=mahonia.NewDecoder("gbk")
	file, err := os.OpenFile("result.txt", os.O_CREATE|os.O_APPEND,2)
	if err != nil {
		fmt.Println("Savefile Open fail")
	}
	defer file.Close()
	for _,info := range result.infos {
		var ips string
		for _,ip := range info.ip{
			ips = ips+ip+","
		}
		res_string := fmt.Sprintf("Url:%s    Port:%d    Code:%d    Title:%s    Server:%-s    Proto:%s    Location:%s    ip:%s",info.url,info.port,info.Statuscode,info.Title,info.Server,info.Proto,info.Location,ips)//格式化字符串
		fmt.Println(res_string)
		if err == nil {
			file.WriteString(res_string)
			file.WriteString("\n")
		}
	}
	fmt.Printf("Save to file result.txt")
	return
}

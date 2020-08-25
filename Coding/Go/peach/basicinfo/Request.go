package basicinfo

import (
	"../nfuncions"
	"crypto/tls"
	"fmt"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"net/http"
	"regexp"
	"sync"
)

func sendrequest(domain []domaininfo) saveres {
	var saveres saveres
	var mutex1 sync.Mutex
	var mutex2 sync.Mutex
	var mutex3 sync.Mutex
	var wg sync.WaitGroup
	for _, urlinfo := range domain { //每一个domain
		wg.Add(1)
		if urlinfo.ip != nil {
			go func(inurlinfo domaininfo) {
				defer wg.Add(-1)
				if nfuncions.IsContain(80, inurlinfo.port) {
					requrl := "http://" + inurlinfo.url
					result := request(requrl)
					if result == nil{
						return
					}
					result.ip = inurlinfo.ip
					result.port = 80
					mutex1.Lock()
					saveres.infos = append(saveres.infos, *result)
					saveres.len += 1
					mutex1.Unlock()
				}
				if nfuncions.IsContain(443, inurlinfo.port) {
					requrl := "https://" + inurlinfo.url
					result := request(requrl)
					if result == nil{
						return
					}
					result.ip = inurlinfo.ip
					result.port = 443
					mutex2.Lock()
					saveres.infos = append(saveres.infos, *result)
					saveres.len += 1
					mutex2.Unlock()
				}
			}(urlinfo)
			} else { //域名未解析到IP
				var nulldoamin respinfo
				nulldoamin.url = urlinfo.url
				mutex3.Unlock()
				saveres.infos = append(saveres.infos, nulldoamin)
				mutex3.Unlock()
			}
}
	wg.Wait()
	return saveres
}


func request(url string) *respinfo{
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true},}//忽略https证书
	clent := &http.Client{Transport:tr}
	fmt.Print("Request url: "+url+"\n")
	resp, err := clent.Get(url)//发送请求,获取响应
	if err != nil{
//		fmt.Println("Get request error")
		return nil
	}	else {
		resps := handleresp(resp)
		if resps == nil{
			return nil
		}
		resps.url = url
		resp.Body.Close()
		return resps
	}
}

func handleresp(response *http.Response) *respinfo{
	redirect := []int{302, 307, 301, 308}//跳转状态码
	var result respinfo
	//enc :=mahonia.NewDecoder("gbk")
	reg1 = regexp.MustCompile(`charset=utf-8`)
	body, err := ioutil.ReadAll(response.Body)//读取返回内容
	if err != nil{
	//	fmt.Println("Get response body error")
		return nil
	}
	convtbody:=mahonia.NewDecoder("gbk").ConvertString(string(body))
	reg := regexp.MustCompile(`<title>(.*)</title>`)
	if reg.FindStringSubmatch(convtbody) == nil {
		result.Title = "None"
	}else{
		result.Title = reg.FindStringSubmatch(convtbody)[1]//title
	}
	result.Statuscode = response.StatusCode //状态码
	if nfuncions.IsContain(response.StatusCode,redirect){//获取跳转地址
		result.Location = response.Header.Get("Location")
	}else {
		result.Location ="None"
	}
	//result.Proto = response.Proto //协议
	result.Server = response.Header.Get("Server") //Server
	return &result
}
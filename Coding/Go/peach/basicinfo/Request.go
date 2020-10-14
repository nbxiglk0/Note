package basicinfo

import (
	"../nfuncions"
	"crypto/tls"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"sync"
	"time"
)

func sendrequest(domain []domaininfo) saveres {
	var saveres saveres
	var mutex1 sync.Mutex
	var mutex2 sync.Mutex
	var mutex3 sync.Mutex
	var wg sync.WaitGroup
	fmt.Println("Start Request")
	for _, urlinfo := range domain { //每一个domain
		wg.Add(1)
		if urlinfo.ip != nil {
			go func(inurlinfo domaininfo) {
				defer wg.Add(-1)
/*				if nfuncions.IsContain(80, inurlinfo.port) {
					requrl := "http://" + inurlinfo.url
					result := request(requrl)
					if result == nil{
						return
					}
					result.Ip = inurlinfo.ip
					result.Port = 80
					mutex1.Lock()
					saveres.Infos = append(saveres.Infos, *result)
					saveres.Len += 1
					mutex1.Unlock()
				}
				if nfuncions.IsContain(443, inurlinfo.port) {
					requrl := "https://" + inurlinfo.url
					result := request(requrl)
					if result == nil{
						return
					}
					result.Ip = inurlinfo.ip
					result.Port = 443
					mutex2.Lock()
					saveres.Infos = append(saveres.Infos, *result)
					saveres.Len += 1
					mutex2.Unlock()
				}*/
				for _,port := range inurlinfo.port{
					if 80 == port||nfuncions.Istls(inurlinfo.ip[0],port)==false && port != 443 {
						requrl := "http://" + inurlinfo.url+":"+strconv.Itoa(port)
						result := request(requrl)
						if result == nil{
							return
						}
						result.Ip = inurlinfo.ip
						result.Port = port
						mutex1.Lock()
						saveres.Infos = append(saveres.Infos, *result)
						saveres.Len += 1
						mutex1.Unlock()
					}
					if 443 == port||nfuncions.Istls(inurlinfo.ip[0],port)==true {
						requrl := "https://" + inurlinfo.url+":"+strconv.Itoa(port)
						result := request(requrl)
						if result == nil {
							return
						}
						result.Ip = inurlinfo.ip
						result.Port = port
						mutex2.Lock()
						saveres.Infos = append(saveres.Infos, *result)
						saveres.Len += 1
						mutex2.Unlock()
					}
				}
			}(urlinfo)
			} else { //域名未解析到IP
				var nulldoamin respinfo
				nulldoamin.Url = urlinfo.url
				mutex3.Unlock()
				saveres.Infos = append(saveres.Infos, nulldoamin)
				mutex3.Unlock()
			}
}
	wg.Wait()
	return saveres
}

func request(url string) *respinfo{
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true},}//忽略https证书
	clent := &http.Client{Transport:tr,Timeout: time.Duration(3* time.Second)}
	//fmt.Print("Request url: "+url+"\n")
	resp, err := clent.Get(url)//发送请求,获取响应
	if err != nil{
//		fmt.Println("Get request error")
		return nil
	}	else {
		resps := handleresp(resp)
		if resps == nil{
			return nil
		}
		resps.Url = url
		resp.Body.Close()
		return resps
	}
}

func handleresp(response *http.Response) *respinfo{
	redirect := []int{302, 307, 301, 308}//跳转状态码
	var result respinfo
	body, err := ioutil.ReadAll(response.Body)//读取返回内容
	if err != nil{
	//	fmt.Println("Get response body error")
		return nil
	}
	convtbody := string(body)
	reg := regexp.MustCompile(`<title>(.*)</title>`)
	if reg.FindStringSubmatch(convtbody) == nil {
		result.Title = "None"
	}else{
		result.Title = reg.FindStringSubmatch(convtbody)[1]//title
		if nfuncions.IsUtf8([]byte(result.Title))==false {//转换title编码
			tempdata, err := simplifiedchinese.GBK.NewDecoder().Bytes([]byte(result.Title))
			if err == nil {
				result.Title = string(tempdata)
			}
		}
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
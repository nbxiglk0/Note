package basicinfo

import (
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

func domainHandle(domain []string,scanmode string,ports string) []domaininfo {
	var channel []domaininfo
	var mutex sync.Mutex
	var wg sync.WaitGroup
	for _, url := range domain {
		wg.Add(1)
		go func(insideurl string) {
			defer wg.Add(-1)
			var results domaininfo
			var ips []string
			if scanmode == "domain" {
				ip := domain2ip(insideurl)
				if ip == nil {
					return
				}
				ips = *ip
			}else {
				ip := insideurl
				ips = append(ips,ip)
			}
			portresult := checkport(ips[0],ports)
			results.url = insideurl
			results.port = portresult
			results.ip = ips
			mutex.Lock()
			channel = append(channel, results)
			mutex.Unlock()
		}(url)
	}
	wg.Wait()
	return channel
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
func checkport(ip string,scanports string) []int{

	scanport := strings.Split(scanports,",")//扫描的端口
	var ports []int
	for _,port := range scanport {
		reqip := ip+":"+port
		timeout := 3*time.Second
		_, err := net.DialTimeout("tcp", reqip,timeout)
		if err == nil{
			temp,_ := strconv.Atoi(port)
			ports = append(ports,temp)
		}
	}
	return ports
}
func Handleip(iprange string )[]string{
	var parseres []string
	ip := strings.Split(iprange,"-")
	index := strings.LastIndex(ip[0],".")
	indexip := ip[0][:index]+"."
	start,_ := strconv.Atoi(strings.Split(ip[0],".")[3])//起始IP
	end,_ := strconv.Atoi(ip[1])//结束IP
	for i := start;i<=end; i++{
			temp := indexip+strconv.Itoa(i)
			parseres = append(parseres,temp)
	}
	return parseres
}
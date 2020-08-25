package basicinfo

import (
	"net"
	"strconv"
	"sync"
	"time"
)

func domainHandle(domain []string) []domaininfo {
	var channel []domaininfo
	var mutex sync.Mutex
	var wg sync.WaitGroup
	var ips []string
	for _, url := range domain {
		wg.Add(1)
		go func(insideurl string) {
			defer wg.Add(-1)
			var results domaininfo
			ip := domain2ip(insideurl)
			if ip == nil{
				return
			}
			ips = *ip
			portresult := checkport(ips[0])
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
func checkport(ip string) []int{
	scanports :=[] int {80,443}//扫描的端口
	var ports []int
	for _,port := range scanports {
		reqip := ip+":"+strconv.Itoa(port)
		timeout := 3*time.Second
		_, err := net.DialTimeout("tcp", reqip,timeout)
		if err == nil{
			ports = append(ports,port)
		}
	}
	return ports
}

func ipHandle(ip []string){

}
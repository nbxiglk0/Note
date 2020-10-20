package getdict

import (
    "fmt"
    "os"
    "strings"
    "time"
)

var suffix =[]string{".7z",".rar",".war",".tar",".gz",".zip",".tar.gz",".bak"}
var normalwords =[]string{"bak","bakeup","beifen","bf","rar","war","source","ym","web","www","备份","data","webroot","1"}


func Main(keywords string ){
    var final []string
    newords := strings.Split(keywords,",")
    for _,value := range newords{
       finaldict := Generate(value)
       final  = append(final,finaldict...)
 }
 final = append(final,date()...)
 final = append(final,normal()...)
 save2file(&final)
}


func Generate(keyword string)[]string{
    dicts := forkeyword(keyword)
    return  dicts
}
//近一年时间
func date()[]string{
    var result []string
    var finalresult []string
    timeFormatTpl := "20060102"//格式化模板
    dateTime := time.Now()
    before := dateTime.AddDate(-1,0,0) //获取前一年时间
    for {
        before = before.AddDate(0,0,1)
        result  = append(result,before.Format(timeFormatTpl))
        if before.Format(timeFormatTpl) == dateTime.Format(timeFormatTpl){
            break
        }
    }
    for _,time := range result {
         for _, value := range suffix {
            new := time + value
            finalresult = append(finalresult, new)
        }
    }
    return  finalresult
}
//关键字
func forkeyword(keyword string)[]string{
    var result []string
    for _,value := range suffix{
        new := keyword+value
        result = append(result,new)
    }
    return  result
}
//常见文件名
func normal()[]string{
    var result []string
    for _,words := range normalwords{
        for _,value := range suffix{
            new := words + value
            result = append(result, new)
        }
    }
    //var result []string
    return  result
}

func save2file(result *[]string){
    filename := "dicts.dic"
    results := *result
    count := len(results)
    file, err := os.OpenFile(filename, os.O_WRONLY | os.O_CREATE|os.O_TRUNC,0644)
    defer file.Close()
    if err != nil{
        fmt.Println("Save to file Fail")
        return
    }else {
        for _,value := range results{
            file.WriteString(value+"\n")
        }
    }
    fmt.Printf("Save to File %s\nTotal Count: %d\n", filename, count)
    return

}
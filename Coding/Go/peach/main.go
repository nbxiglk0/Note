package main

import (
	"./basicinfo"
	"./getdict"
	"fmt"
	"time"
)

func main(){
//	gui.Main()
	start := time.Now()
	callback := parse()
	switch callback.method {
	case 1:
		basicinfo.Main(callback.argv,callback.scanmode)
	case 2:
		getdict.Main(callback.argv["keywords"])
	}
	cost := time.Since(start)
	fmt.Printf("The Task cost time %s",cost)
}

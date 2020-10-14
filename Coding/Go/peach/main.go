package main

import (
	"./basicinfo"
	"fmt"
	"time"
)
import "./getdict"

func main(){
//	gui.Main()
	start := time.Now()
	callback := parse()
	switch callback.method {
	case 1:
		basicinfo.Main(callback.argv)
	case 2:
		getdict.Main(callback.argv["dict1"],callback.argv["dict2"])
	}
	cost := time.Since(start)
	fmt.Printf("The Task cost time %s",cost)
}

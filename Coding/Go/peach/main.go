package main

import "./basicinfo"
import "./getdict"

func main(){
//	gui.Main()
	callback := parse()
	switch callback.method {
	case 1:
		basicinfo.Main(callback.argv["filepath"])
	case 2:
		getdict.Main(callback.argv["dict1"],callback.argv["dict2"])
	}
}

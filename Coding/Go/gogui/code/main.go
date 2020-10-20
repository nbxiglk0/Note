// 由res2go IDE插件自动生成。
package main

import (
    "github.com/ying32/govcl/vcl"
)

func main() {
    vcl.Application.SetScaled(true)
    vcl.Application.SetTitle("project1")
    vcl.Application.Initialize()
    vcl.Application.SetMainFormOnTaskBar(true)
    vcl.Application.CreateForm(&Peach)
    vcl.Application.Run()
}

// Automatically generated by the res2go IDE plug-in.
package main

import (
    "github.com/ying32/govcl/vcl"
)

func main() {
    vcl.Application.SetScaled(true)
    vcl.Application.SetTitle("demo1")
    vcl.Application.Initialize()
    vcl.Application.SetMainFormOnTaskBar(true)
    vcl.Application.CreateForm(&Form1)
    vcl.Application.Run()
}

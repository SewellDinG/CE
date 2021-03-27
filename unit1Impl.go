package main

import (
    "github.com/ying32/govcl/vcl"
)

//::private::
type TForm1Fields struct {
}

func (f *TForm1) OnCheckClick(sender vcl.IObject) {
    url := f.URL.Text()
    if url == "" {
        vcl.ShowMessage("None")
        return
    }
    f.Output.SetText(url)
}

func (f *TForm1) OnSetThreadClick(sender vcl.IObject) {

}


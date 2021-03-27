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
    var resp *Response
    resp = HttpReqer(url, "GET", "", Headers{})
    f.Output.SetText(string(resp.Body))
}

func (f *TForm1) OnSetThreadClick(sender vcl.IObject) {

}

func (f *TForm1) OnCheckGroup1Click(sender vcl.IObject) {

}

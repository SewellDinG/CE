package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "reflect"
    "strings"
    "testing"
)

func TestHTTP(t *testing.T) {
    target := "https://httpbin.org"
    var poc = &Poc{}
    poc = poc.PocGetter("./httpbin.org-post.yml")
    for i, rule := range poc.Rules {
        fmt.Println("------------------ Rule NO.", i+1, "------------------")
        // 初始化rule
        url := target + rule.Path
        ruleHeaders := rule.Headers
        ruleBody := rule.Body
        fmt.Printf("[*] rule: %+v\n", rule)
        // 判断请求类型
        var resp *Response
        switch rule.Method {
        case "GET":
            resp = HttpReqer(url, "GET", "", Headers{})
        case "POST":
            resp = HttpReqer(url, "POST", ruleBody, ruleHeaders)
        case "PUT":
            resp = HttpReqer(url, "PUT", "", Headers{})
        case "MOVE":
            resp = HttpReqer(url, "MOVE", "", Headers{})
        default:
            fmt.Println("No allowed method...")
        }
        fmt.Printf("[*] resp: %+v\n", resp)
        fmt.Println("[*] resp.Body: ", string(resp.Body))
    }
}

type Response struct {
    Url     string
    Status  int
    Body    []byte
    Headers http.Header
}

func HttpReqer(url, httpMethod, ruleBody string, ruleHeaders Headers) *Response {
    client := &http.Client{}
    contentBody := strings.NewReader(ruleBody)
    req, err := http.NewRequest(httpMethod, url, contentBody)
    if err != nil {
        fmt.Println("http.NewRequest err:", err)
    }
    // 设置请求头
    // Headers struct 空判断
    if !reflect.DeepEqual(ruleHeaders, Headers{}) {
        ruleHeader := StructToMap(ruleHeaders)
        for contentTypeKey, contentTypeValue := range ruleHeader {
            //fmt.Println(contentTypeKey, "--------", contentTypeValue)
            if contentTypeKey == "ContentType" {
                contentTypeKey = "Content-Type"
            }
            req.Header.Set(contentTypeKey, contentTypeValue)
        }
    }
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("client.Do err:", err)
    }
    defer func() { _ = resp.Body.Close() }()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("ioutil.ReadAll err:", err)
    }
    var respStruct = &Response{url, resp.StatusCode, body, resp.Header}
    return respStruct
}

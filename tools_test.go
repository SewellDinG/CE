package main

import (
    "fmt"
    "testing"
)

func TestTools(t *testing.T) {
    type structTest struct {
        one   string
        two   int
        three bool
    }
    fmt.Println("Struct to Map: ", StructToMap(structTest{"one", 2, true}))

    var s1 = "helloworld"
    var s2 = "world"
    fmt.Println("Contains: ", Contains(s1, s2))
    fmt.Println("StartsWith: ", StartsWith(s1, s2))
    fmt.Println("EndsWith: ", EndsWith(s1, s2))
    var s3 = []byte("helloworld")
    var s4 = []byte("world")
    fmt.Println("Bcontains: ", Bcontains(s3, s4))

    fmt.Println("Md5: ", Md5(s1))
    fmt.Println("Base64: ", Base64(s1))
    var s5 = "aGVsbG93b3JsZA=="
    fmt.Println("Base64Decode: ", Base64Decode(s5))

    fmt.Println("RandomInt: ", RandomInt(0, 9))
    fmt.Println("RandomLowercase: ", RandomLowercase(32))

    var s6 = "http://baidu.com/index.php?id=1&ok=1"
    fmt.Println("UrlEncode: ", UrlEncode(s6))
    var s7 = "http://baidu.com/index.php?id=1%26ok=1"
    fmt.Println("UrlDecode: ", UrlDecode(s7))
}

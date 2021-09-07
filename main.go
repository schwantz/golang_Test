package main

import (
	httpser "cruise.com/maintest/MyhttpSerTest"
	tt "cruise.com/maintest/MysqlTest"
	"encoding/json"
	"fmt"
	"reflect"
	"unsafe"

	"io/ioutil"
	"net/http"
	"runtime"
)

type J struct {
	a    string
	B    string `json:"name"`
	Ctag string
	D    string `json:"DDe"`
}

func main() {
	j := J{
		a:    "11",
		B:    "22",
		Ctag: "333",
		D:    "44",
	}

	fmt.Printf("j before: = %+v \n", j)
	jsonInfo, vv := json.Marshal(j)
	fmt.Printf("json format output！ %+v,%#v \n", string(jsonInfo), vv)

	fmt.Printf("----------------------------------------------------------\n")

	a := "aaa"
	ssh := *(*reflect.StringHeader)(unsafe.Pointer(&a))
	//b := *(*[]byte)(unsafe.Pointer(&ssh))
	fmt.Printf("content= %#v,%#v,%#v \n", ssh, *(*int)(unsafe.Pointer(&ssh.Len)), *(*[]byte)(unsafe.Pointer(&ssh)))

	fmt.Printf("----------------------------------------------------------\n")
	num := 1
	var stringbody string
	for index := 0; index < num; index++ {
		resp, _ := http.Get("https://www.google.com")
		s1, _ := ioutil.ReadAll(resp.Body)
		stringbody = string(s1)
		resp.Body.Close()
	}

	fmt.Printf("now goroutine= %d\n", runtime.NumGoroutine())
	fmt.Printf("body length=%d\n", len(stringbody))


	fmt.Printf("----------------------------------------------------------\n")

	tt.MysqlTest()
	httpser.MyhttpSerTest1()
}

package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	//"reflect"
	"reflect"
)

type js_str struct  {
	Id string `json : "id"`
	State string `json : "state"`
}
func main() {
	resp, _ := http.Get("http://10.85.43.1:8080/compute/v1/apps/megh/instances/11")
	b, _ := ioutil.ReadAll(resp.Body)
	var data interface{}
	err := json.Unmarshal(b, &data)
	fmt.Printf("%s",err)
	//fmt.Fprintln(data)
	//fmt.Printf("%v", data["primary_ip"])
	//var data1 js_str
	fmt.Println(reflect.TypeOf(data))
	//data1.Id := data
	//fmt.Println(data1)
	if err != nil {
	panic(err.Error())
	//fmt.Printf(data["hostname"])
	//fmt.Printf("%v",data)
	}

	//fmt.Printf("%s", b )
}
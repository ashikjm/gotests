
package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	//"reflect"
	//"bytes"
	"encoding/json"
)

type bdy struct  {

	Resource_type string `json:"resource_type"`
	Event_type string `json:"event_type"`
	Ctime string `json:"resource_id"`
}

func url_get (a string)  {
	resp, _ := http.Get(a)
	body, _ := ioutil.ReadAll(resp.Body)
	var js_obj bdy
	json.Unmarshal(body, &js_obj)
	fmt.Printf("%v\n", js_obj.Ctime)
	//fmt.Printf("%s\n", body)
	//r := bytes.NewReader(body)
	//fmt.Printf("%s\n", r)
	//var js_obj interface{}
	//json.Unmarshal(body, &js_obj)
	//fmt.Printf("%v\n",js_obj.(map[string]interface{})["ctime"])
	//fmt.Printf("%v", body)
	//return reflect_type
}
func main() {
//x := map[string]interface{}{
//"foo": []string{"a","b"},
//"bar": "foo",
//"baz": 10.4,
//}
	url_get("http://wbuild-0002.nm.flipkart.com:9000/v1/servers/Dell1/events/2")

//fmt.Printf("%s\n", type_of_string)
//fmt.Printf("%#v\n", x)
//	fmt.Printf("% v", x["foo"])
}


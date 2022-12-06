package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"projectcreation/createHarbor"
)

type HarborProject struct {
	PN string
}

func myHarbor(w http.ResponseWriter, r *http.Request) {
	//defer fmt.Println(w, "OK")
	fmt.Println("method:", r.Method)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read body err:%v", err)
	} else {
		println("json内容是:\n", string(body))
	}
	jsonData := []byte(string(body))

	var name HarborProject
	err = json.Unmarshal(jsonData, &name)
	if err != nil {
		fmt.Printf("error:%v", err)

	} else {
		fmt.Printf("web端输入的项目名是:%v \n", name.PN)
	}
	createHarbor.Logon(name.PN)
}

func main() {
	server := http.Server{
		Addr:         ":8090",
		ReadTimeout:  0,
		WriteTimeout: 0,
	}
	http.HandleFunc("/", myHarbor)
	server.ListenAndServe()
}

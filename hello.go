package main

import (
	"fmt"
	"io/ioutil"
  "net/http"
  "log"

)

func httpGet() {
	resp, err := http.Get("https://recommd.xyq.cbg.163.com/cgi-bin/recommend.py?act=recommd_by_role&server_id=554&areaid=58&server_name=%E5%85%B0%E4%BA%AD%E5%BA%8F&page=1&view_loc=equip_list&count=15")
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
  http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
    writer.Header().Set("Access-Control-Allow-Origin", "*")
    fmt.Fprintln(writer, string(body))
  })
  log.Fatal(http.ListenAndServe(":8080",nil))

}

func main() {
	httpGet()

}
package main

import "fmt"
import "net/http"
import "io/ioutil"

func main() {
	var username string = "admin"
	var passwd string = "geoserver"
	var url string = "http://192.168.33.10:8080/geoserver/rest/"
	var element string = "datastores"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url+element, nil)
	req.SetBasicAuth(username, passwd)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

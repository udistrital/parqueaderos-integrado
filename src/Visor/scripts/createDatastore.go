package main

import "fmt"
import "net/http"
import "io/ioutil"
import "bytes"

func main() {
	var username string = "admin"
	var passwd string = "geoserver"
	var url string = "http://192.168.33.10:8080/geoserver/rest/"
	var element string = "workspaces/"
	var workspace string = "parqueaderos/"
	var element2 string = "datastores"
	var jsondata string = "{'dataStore':{'name':'circe','type':'h2','enabled':'true'," +
		"'connectionParameters':{'dbtype':'postgis','host':'localhost','port':'5432','database':'circe','schema':'public','user':'usercirce','passwd':'oascirce'}}}"

	client := &http.Client{}
	b := bytes.NewBuffer([]byte(jsondata))
	req, err := http.NewRequest("POST", url+element+workspace+element2+".json", b)
	req.Header.Set("Content-Type", "application/json")
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

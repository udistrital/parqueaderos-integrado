package main

import "fmt"
import "net/http"
import "io/ioutil"
import "bytes"

func main() {
	var username string = "admin"
	var passwd string = "geoserver"
	var url string = "http://192.168.33.10:8080/geoserver/rest/"
	var element string = "layers"
	var jsondata string = "{'layer':{'name':'grupo_isla','enable':'true','advertised':'true','type':'VECTOR','defaultStyle':{'name':'polygon','href':'http://192.168.33.10:8080/geoserver/rest/styles/polygon.json'},'resource':{'@class':'featureType','name':'grupo_isla','href':'http://192.168.33.10:8080/geoserver/rest/workspaces/parqueaderos/datastores/circe/featuretypes/grupo_isla.json'},'attribution':{'logoWidth':0,'logoHeight':0}}}"

	client := &http.Client{}
	b := bytes.NewBuffer([]byte(jsondata))
	req, err := http.NewRequest("POST", url+element+".json", b)
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

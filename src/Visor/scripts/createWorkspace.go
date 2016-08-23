package main

import "fmt"
import "net/http"
import "io/ioutil"
import "bytes"

func main() {
    var username string = "admin"
    var passwd string = "geoserver"
    var url string = "http://192.168.33.10:8080/geoserver/rest/"
    client := &http.Client{}
    var jsondata string = "{'workspace':{'name':'parqueaderos'}}"
    b := bytes.NewBuffer([]byte(jsondata))
    req, err := http.NewRequest("POST", url + "workspaces.json", b)
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

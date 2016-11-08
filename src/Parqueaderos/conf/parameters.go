package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type _Parameters struct {
	DATABASE_HOST     string
	DATABASE_PORT     string
	DATABASE_NAME     string
	DATABASE_SCHEMA   string
	DATABASE_USER     string
	DATABASE_PASSWORD string
}

var (
	Parameters = _Parameters{}
)

func init() {
	//secret := beego.AppConfig.String("secret")
	data, err := ioutil.ReadFile("parameters.yml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal([]byte(data), &Parameters)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	//log.Printf("parameters:\n%v\n\n", p)
}

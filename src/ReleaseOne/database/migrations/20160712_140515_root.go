package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Root_20160712_140515 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Root_20160712_140515{}
	m.Created = "20160712_140515"
	migration.Register("Root_20160712_140515", m)
}

// Run the migrations
func (m *Root_20160712_140515) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *Root_20160712_140515) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}

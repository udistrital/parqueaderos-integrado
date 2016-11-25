package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type UpdateNfcType_20161125_125149 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &UpdateNfcType_20161125_125149{}
	m.Created = "20161125_125149"
	migration.Register("UpdateNfcType_20161125_125149", m)
}

// Run the migrations
func (m *UpdateNfcType_20161125_125149) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	var sql string = "ALTER TABLE public.vehiculo " +
	"ALTER COLUMN id_nfc TYPE bigint USING id_nfc::bigint";
	m.SQL(sql);
}

// Reverse the migrations
func (m *UpdateNfcType_20161125_125149) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}

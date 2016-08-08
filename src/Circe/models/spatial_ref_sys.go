package models

type SpatialRefSys struct {
	Srid      int    `orm:"column(srid)"`
	AuthName  string `orm:"column(auth_name);null"`
	AuthSrid  int    `orm:"column(auth_srid);null"`
	Srtext    string `orm:"column(srtext);null"`
	Proj4text string `orm:"column(proj4text);null"`
}

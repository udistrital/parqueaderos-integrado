package models

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type Isla struct {
	Id          int        `orm:"column(id);pk;auto"`
	Ocupado     bool       `orm:"column(ocupado)"`
	Geometria   string     `orm:"column(geometria)"`
	IdGrupoIsla *GrupoIsla `orm:"column(id_grupo_isla);rel(fk)"`
}

func (t *Isla) TableName() string {
	return "isla"
}

func init() {
	orm.RegisterModel(new(Isla))
}

// AddIsla insert a new Isla into database and returns
// last inserted Id on success.
func AddIsla(m *Isla) (id int64, err error) {
	o := orm.NewOrm()
	valid := validation.Validation{}
	valid.Required(m.Ocupado, "Ocupado")
	valid.Required(m.Geometria, "Geometria")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	} else {
		log.Println("Insert New Register")
		m.IdGrupoIsla, _ = GetGrupoIslaById(m.IdGrupoIsla.Id)
		id, err = o.Insert(m)
	}
	return
}

// GetIslaById retrieves Isla by Id. Returns error if
// Id doesn't exist
func GetIslaById(id int) (v *Isla, err error) {
	o := orm.NewOrm()
	v = &Isla{Id: id}
	if err = o.Read(v); err == nil {
		v.IdGrupoIsla, _ = GetGrupoIslaById(v.IdGrupoIsla.Id)
		return v, nil
	}
	return nil, err
}

// GetAllIsla retrieves all Isla matches certain condition. Returns empty list if
// no records exist
func GetAllIsla(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Isla))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Isla
	qs = qs.OrderBy("id")
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				v.IdGrupoIsla, _ = GetGrupoIslaById(v.IdGrupoIsla.Id)
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateIsla updates Isla by Id and returns error if
// the record to be updated doesn't exist
func UpdateIslaById(m *Isla) (err error) {
	o := orm.NewOrm()
	v := Isla{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteIsla deletes Isla by Id and returns error if
// the record to be deleted doesn't exist
func DeleteIsla(id int) (err error) {
	o := orm.NewOrm()
	v := Isla{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Isla{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

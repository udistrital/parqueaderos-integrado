package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Registro struct {
	Id          int       `orm:"column(id);pk";auto`
	IdVehiculo  *Vehiculo `orm:"column(id_vehiculo);rel(fk)"`
	IdIsla      *Isla     `orm:"column(id_isla);rel(fk)"`
	HoraEntrada time.Time `orm:"column(hora_entrada);type(timestamp without time zone);null"`
	HoraSalida  time.Time `orm:"column(hora_salida);type(timestamp without time zone);null"`
}

func (t *Registro) TableName() string {
	return "registro"
}

func init() {
	orm.RegisterModel(new(Registro))
}

// AddRegistro insert a new Registro into database and returns
// last inserted Id on success.
func AddRegistro(m *Registro) (id int64, err error) {
	o := orm.NewOrm()
	m.HoraEntrada = time.Now()
	m.IdVehiculo, _ = GetVehiculoById(m.IdVehiculo.Id)
	m.IdIsla, _ = GetIslaById(m.IdIsla.Id)
	id, err = o.Insert(m)
	return
}

// GetRegistroById retrieves Registro by Id. Returns error if
// Id doesn't exist
func GetRegistroById(id int) (v *Registro, err error) {
	o := orm.NewOrm()
	v = &Registro{Id: id}
	if err = o.Read(v); err == nil {
		v.IdVehiculo, _ = GetVehiculoById(v.IdVehiculo.Id)
		v.IdIsla, _ = GetIslaById(v.IdIsla.Id)
		return v, nil
	}
	return nil, err
}

// GetAllRegistro retrieves all Registro matches certain condition. Returns empty list if
// no records exist
func GetAllRegistro(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Registro))
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

	var l []Registro
	qs = qs.OrderBy("id")
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				v.IdVehiculo, _ = GetVehiculoById(v.IdVehiculo.Id)
				v.IdIsla, _ = GetIslaById(v.IdIsla.Id)
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

// UpdateRegistro updates Registro by Id and returns error if
// the record to be updated doesn't exist
func UpdateRegistroById(m *Registro) (err error) {
	o := orm.NewOrm()
	v := Registro{Id: m.Id}
	m.HoraSalida = time.Now()
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		m.HoraEntrada = v.HoraEntrada
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteRegistro deletes Registro by Id and returns error if
// the record to be deleted doesn't exist
func DeleteRegistro(id int) (err error) {
	o := orm.NewOrm()
	v := Registro{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Registro{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

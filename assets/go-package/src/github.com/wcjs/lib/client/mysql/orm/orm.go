package orm

import (
	"errors"
	"fmt"
	"reflect"
	"git.chemm.com/backend/lib/client/mysql"
)

type Join struct {
	TableName string
	Alias     string
	OnStr     string
}

type Table struct {
	TableName string
	Alias     string
	Join      []*Join
}

type Model interface {
	TableInfo() *Table
}

type InsertDictStruct struct {
	Field            string
	NeedYinQuotation bool
	Index            int
}

var Mysqlclint *mysql.MultiMysqlClient

const SqlFormat string = "select %s from %s%s%s%s%s%s%s%s"
const InsertSqlFormat string = "insert into %s (%s) values %s %s" //tablename,InsertField(eg : `id`, `avatar`),InsertValue(eg:NULL, 'a'),duplicate(eg:nickname=nickname;)

//INSERT INTO `w_robot` (`id`, `avatar`, `nickname`, `admin_id`, `robot_id`) VALUES (NULL, 'a', 'a', '11', '11') ON DUPLICATE KEY UPDATE nickname='a';

//SELECT %WHAT% FROM %TABLE% %FORCE% %JOIN% %WHERE% %GROUP% %HAVING% %ORDER% %LIMIT%

type Orm struct {
	CountDistict string
	WhatStr      string
	Table        *Table
	WhereStr     string
	ForceStr     string
	JoinStr      string
	GroupStr     string
	HavingStr    string
	OrderStr     string
	LimitStr     string

	InsertFields          map[string]bool
	InsertDuplicateFields map[string]bool

	Mysqlclint *mysql.MultiMysqlClient
}

func SetClient(mc *mysql.MultiMysqlClient) {
	Mysqlclint = mc
}

func NewOrm() *Orm {
	return &Orm{
		InsertFields:          make(map[string]bool, 0),
		InsertDuplicateFields: make(map[string]bool, 0),
		Mysqlclint:Mysqlclint,
	}
}

func (o *Orm) SetClient(mc *mysql.MultiMysqlClient) *Orm {
	o.Mysqlclint = mc
	return o
}


func (o *Orm) Distinct(s string) *Orm {
	o.CountDistict = s
	return o
}

func (o *Orm) clearPasre() {
	o.CountDistict = ""
	o.WhatStr = ""
	o.WhereStr = ""
	o.ForceStr = ""
	o.JoinStr = ""
	o.GroupStr = ""
	o.HavingStr = ""
	o.OrderStr = ""
	o.LimitStr = ""

	o.InsertFields = make(map[string]bool, 0)
	o.InsertDuplicateFields = make(map[string]bool, 0)
}

func (o *Orm) pasreSql() (sql string) {
	if o.WhereStr != "" {
		o.WhereStr = " where " + o.WhereStr
	}

	var tablename string
	if o.Table.Alias != "" {
		tablename = o.Table.TableName + " as " + o.Table.Alias
	} else {
		tablename = o.Table.TableName
	}

	for _, j := range o.Table.Join {
		o.JoinStr += " left join " + j.TableName + " as " + j.Alias + " on " + j.OnStr
	}
	sql = fmt.Sprintf(SqlFormat, o.WhatStr, tablename, o.ForceStr, o.JoinStr, o.WhereStr, o.GroupStr, o.HavingStr, o.OrderStr, o.LimitStr)
	o.clearPasre()
	return
}

func (o *Orm) Limit(offset int64, limit int64) *Orm {
	o.LimitStr = fmt.Sprintf(" limit %d,%d", offset, limit)
	return o
}

func (o *Orm) Having(format string, i ...interface{}) *Orm {
	var err error
	if o.GroupStr == "" {
		panic(errors.New("must call Group first"))
	}
	o.HavingStr, err = mysql.MysqlFilter(format, i...)
	if err != nil {
		panic(err)
	}
	return o
}

func (o *Orm) Group(format string, i ...interface{}) *Orm {
	var err error
	o.GroupStr, err = mysql.MysqlFilter(format, i...)
	if err != nil {
		panic(err)
	}
	return o
}

func (o *Orm) Force(indexstr string) *Orm {
	o.ForceStr = " force index(" + indexstr + ")"
	return o
}

func (o *Orm) Order(format string, i ...interface{}) *Orm {
	var err error
	o.OrderStr, err = mysql.MysqlFilter(format, i...)
	if err != nil {
		panic(err)
	}
	return o
}

func (o *Orm) CallMethon(ns string) (n float64) {
	o.WhatStr = ns
	sql := o.pasreSql()
	n, e := o.Mysqlclint.GetRead().QueryCallMethon(sql)
	if e != nil {
		panic(e)
	}
	return
}

func (o *Orm) Count() (n int64) {
	what := ""
	if o.CountDistict != "" {
		what = fmt.Sprintf("count(distinct %s) as num", o.CountDistict)
	} else {
		what = "count(*) as num"
	}

	o.WhatStr = what
	sql := o.pasreSql()
	n, e := o.Mysqlclint.GetRead().QueryCount(sql)
	if e != nil {
		panic(e)
	}
	return
}

func (o *Orm) insertPasreSql() (sql string) {
	if o.WhereStr != "" {
		o.WhereStr = " where " + o.WhereStr
	}

	var tablename string
	if o.Table.Alias != "" {
		tablename = o.Table.TableName + " as " + o.Table.Alias
	} else {
		tablename = o.Table.TableName
	}

	for _, j := range o.Table.Join {
		o.JoinStr += " left join " + j.TableName + " as " + j.Alias + " on " + j.OnStr
	}
	sql = fmt.Sprintf(SqlFormat, o.WhatStr, tablename, o.ForceStr, o.JoinStr, o.WhereStr, o.GroupStr, o.HavingStr, o.OrderStr, o.LimitStr)
	o.clearPasre()
	return
}

func (o *Orm) InsertFeild(f ...string) *Orm {
	for _, ff := range f {
		o.InsertFields[ff] = true
	}
	return o
}

func (o *Orm) InsertDuplicateField(f ...string) *Orm {
	for _, ff := range f {
		o.InsertDuplicateFields[ff] = true
	}
	return o
}

func (o *Orm) Insert(md Model) int64 {
	val := reflect.ValueOf(md)
	if val.Kind() != reflect.Ptr {
		panic("orm Read(md interface{}) md must be ptr")
	}

	ind := reflect.Indirect(val)
	o.Table = md.TableInfo()
	insertFieldStr := ""
	insertValueStr := ""
	duplicateValueStr := ""
	for i := 0; i < ind.NumField(); i++ {
		field := ind.Field(i)
		sf := ind.Type().Field(i)
		tag := sf.Tag

		fieldname := tag.Get("field")
		if fieldname == "" {
			continue
		}

		_, ok := o.InsertFields[fieldname]
		needYinQuotation := false
		if ok {
			if insertFieldStr != "" {
				insertFieldStr += ","
			}
			if insertValueStr != "" {
				insertValueStr += ","
			}
			switch field.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64:
				insertValueStr += fmt.Sprint(field)
			case reflect.String:
				needYinQuotation = true
				fieldvalue, _ := mysql.MysqlFilter("%s", fmt.Sprintf("%s", field))
				insertValueStr += "'" + fieldvalue + "'"
			case reflect.Slice:
				if "[]uint8" != field.Type().String() {
					panic("不支持这个格式")
				}
				needYinQuotation = true
				fieldvalue, _ := mysql.MysqlFilter("%s", fmt.Sprintf("%s", field))
				insertValueStr += "'" + fieldvalue + "'"
			}

			insertFieldStr += "`" + fieldname + "`"
		}
		_, ok2 := o.InsertDuplicateFields[fieldname]

		if ok2 {
			if duplicateValueStr != "" {
				duplicateValueStr += ","
			}
			if needYinQuotation {
				fieldvalue, _ := mysql.MysqlFilter("%s", fmt.Sprintf("%s", field))
				duplicateValueStr += "`" + fieldname + "`='" + fieldvalue + "'"
			} else {
				duplicateValueStr += "`" + fieldname + "`=" + fmt.Sprint(field)
			}
		}
	}
	if duplicateValueStr != "" {
		duplicateValueStr = " ON DUPLICATE KEY UPDATE " + duplicateValueStr
	}

	sql := fmt.Sprintf(`insert into %s (%s) value (%s) %s`, o.Table.TableName, insertFieldStr, insertValueStr, duplicateValueStr)

	o.clearPasre()
	rr, ee := o.Mysqlclint.GetWrite().Exec(sql)
	if ee != nil {
		return 0
	}
	insertId, e := rr.LastInsertId()
	if e != nil {
		return 0
	} else {
		return insertId
	}

}

func (o *Orm) InsertList(ml interface{}) error {
	val := reflect.ValueOf(ml)

	if val.Kind() != reflect.Slice {
		panic("orm InsertList(ml interface{}) md must be Slice")
	}

	l := val.Len()
	if l == 0 {
		return errors.New("ml len is 0")
	}

	e := val.Index(0)

	tablemethod := e.MethodByName("TableInfo")
	if tablemethod.Kind() == reflect.Invalid {
		panic("orm InsertList(ml interface{}) md must has Method 'TableInfo'")
	}

	args := make([]reflect.Value, 0)
	tablemethodresult := tablemethod.Call(args)
	tablename := tablemethodresult[0].Elem().Field(0)

	ee := e.Elem()

	kv := make([]*InsertDictStruct, 0)

	duplicateValueStr := ""
	insertFieldStr := ""

	for i := 0; i < ee.NumField(); i++ {
		field := ee.Field(i)
		sf := ee.Type().Field(i)
		tag := sf.Tag
		fieldname := tag.Get("field")

		if fieldname == "" {
			continue
		}

		_, ok := o.InsertFields[fieldname]
		needYinQuotation := false
		if ok {
			switch field.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64:
				needYinQuotation = false
			case reflect.String:
				needYinQuotation = true
			case reflect.Slice:
				if "[]uint8" != field.Type().String() {
					panic("不支持这个格式")
				}
				needYinQuotation = true
			}

			if insertFieldStr != "" {
				insertFieldStr += ","
			}
			insertFieldStr += "`" + fieldname + "`"

			kv = append(kv, &InsertDictStruct{
				fieldname,
				needYinQuotation,
				i,
			})
		}

		_, ok2 := o.InsertDuplicateFields[fieldname]

		if ok2 {
			if duplicateValueStr != "" {
				duplicateValueStr += ","
			}
			duplicateValueStr += fmt.Sprintf("`%s` = VALUES(%s)", fieldname, fieldname)
		}

	}

	valueStr := ""
	var vStr, vvStr, fieldvalue string
	for j := 0; j < l; j++ {
		d := val.Index(j)
		de := d.Elem()
		vStr = "("
		vvStr = ""
		for _, idict := range kv {
			if vvStr != "" {
				vvStr += ","
			}

			findex := idict.Index
			ff := de.Field(findex)
			if idict.NeedYinQuotation {
				fieldvalue, _ = mysql.MysqlFilter(`"%s"`, fmt.Sprintf("%s", ff))
			} else {
				fieldvalue = fmt.Sprintf("%d", ff)
			}

			vvStr += fieldvalue
		}
		vStr += vvStr
		vStr += ")"
		if valueStr != "" {
			valueStr += ","
		}
		valueStr += vStr
	}

	if duplicateValueStr != "" {
		duplicateValueStr = " ON DUPLICATE KEY UPDATE " + duplicateValueStr
	}
	sql := fmt.Sprintf(`insert into %s (%s) values %s %s`, tablename, insertFieldStr, valueStr, duplicateValueStr)

	o.clearPasre()
	_, err := o.Mysqlclint.GetWrite().Exec(sql)
	if err != nil {
		return err
	}
	return nil

}

func (o *Orm) Save(md Model) {
	val := reflect.ValueOf(md)
	if val.Kind() != reflect.Ptr {
		panic("orm Save(md interface{}) md must be ptr")
	}

	ind := reflect.Indirect(val)
	o.Table = md.TableInfo()

	setstr := ""
	primarywhere := ""
	uniquewhere := ""

	isprimarykey := false
	isuniquekey := false
	for i := 0; i < ind.NumField(); i++ {
		field := ind.Field(i)
		sf := ind.Type().Field(i)
		tag := sf.Tag

		fieldname := tag.Get("field")
		if fieldname == "" {
			continue
		}

		key := tag.Get("key")
		isprimarykey = false
		isuniquekey = false
		if key == "primary" {
			isprimarykey = true
		} else if key == "unique" {
			isuniquekey = true
		}

		_, ok := o.InsertFields[fieldname]

		switch field.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			v := field.Int()
			if isprimarykey {
				if v != 0 {
					if primarywhere != "" {
						primarywhere += " and "
					}
					primarywhere += fmt.Sprintf("%s = %d", fieldname, v)
				}
			} else if isuniquekey {
				if v != 0 {
					if uniquewhere != "" {
						uniquewhere += " and "
					}
					uniquewhere += fmt.Sprintf("%s = %d", fieldname, v)
				}
			}
			if ok {
				if setstr != "" {
					setstr += ","
				}
				setstr += fmt.Sprintf("%s = %d", fieldname, v)
			}

		case reflect.Float32, reflect.Float64:
			v := field.Float()
			if isprimarykey {
				if v != 0 {
					if primarywhere != "" {
						primarywhere += " and "
					}
					primarywhere += fmt.Sprintf("%s = %d", fieldname, v)
				}
			} else if isuniquekey {
				if v != 0 {
					if uniquewhere != "" {
						uniquewhere += " and "
					}
					uniquewhere += fmt.Sprintf("%s = %d", fieldname, v)
				}
			}
			if ok {
				if setstr != "" {
					setstr += ","
				}
				setstr += fmt.Sprintf("%s = %d", fieldname, v)
			}
		case reflect.String:
			v := field.String()
			var ms string
			var e1 error
			if isprimarykey {
				if v != "" {
					ms, e1 = mysql.MysqlFilter("%s = '%s'", fieldname, v)
					if e1 != nil {
						if primarywhere != "" {
							primarywhere += " and "
						}
						primarywhere += ms
					}
				}
			} else if isuniquekey {
				if v != "" {
					ms, e1 = mysql.MysqlFilter("%s = '%s'", fieldname, v)
					if e1 != nil {
						if uniquewhere != "" {
							uniquewhere += " and "
						}
						uniquewhere += ms
					}
				}
			}
			if ok {
				if isprimarykey || isuniquekey {
					if e1 != nil {
						if setstr != "" {
							setstr += ","
						}
						setstr += ms
					}
				} else {
					ms, e1 := mysql.MysqlFilter("%s = '%s'", fieldname, v)
					if e1 != nil {
						if setstr != "" {
							setstr += ","
						}
						setstr += ms
					}
				}
			}

		case reflect.Slice:
			if "[]uint8" != field.Type().String() {
				panic("不支持这个格式")
			}
			v := field.Bytes()
			var ms string
			var e1 error

			if isprimarykey {
				if len(v) != 0 {
					ms, e1 = mysql.MysqlFilter("%s = '%s'", fieldname, v)
					if e1 != nil {
						if primarywhere != "" {
							primarywhere += " and "
						}
						primarywhere += ms
					}
				}
			} else if isuniquekey {
				if len(v) != 0 {
					ms, e1 = mysql.MysqlFilter("%s = '%s'", fieldname, v)
					if e1 != nil {
						if primarywhere != "" {
							primarywhere += " and "
						}
						primarywhere += ms
					}
				}
			}
			if ok {
				if isprimarykey || isuniquekey {
					if e1 != nil {
						if setstr != "" {
							setstr += ","
						}
						setstr += ms
					}
				} else {
					ms, e1 := mysql.MysqlFilter("%s = '%s'", fieldname, v)
					if e1 != nil {
						if setstr != "" {
							setstr += ","
						}
						setstr += ms
					}
				}
			}
		}

	}

	if setstr == "" {
		panic("no set field,should call orm function InsertFeild first")
	}

	var sql string
	if primarywhere != "" {
		sql = fmt.Sprintf("update %s set %s where %s", o.Table.TableName, setstr, primarywhere)
	} else if uniquewhere != "" {
		sql = fmt.Sprintf("update %s set %s where %s", o.Table.TableName, setstr, uniquewhere)
	} else {
		panic("no where,maybe not set primarykey or uniquekey")
	}

	o.clearPasre()
	o.Mysqlclint.GetWrite().Exec(sql)
}

func (o *Orm) Read(md Model) {
	val := reflect.ValueOf(md)
	if val.Kind() != reflect.Ptr {
		panic("orm Read(md interface{}) md must be ptr")
	}

	ind := reflect.Indirect(val)
	//typ := ind.Type()
	o.Table = md.TableInfo()

	dest := make([]interface{}, 0)

	where := ""
	what := ""
	for i := 0; i < ind.NumField(); i++ {
		field := ind.Field(i)
		sf := ind.Type().Field(i)
		tag := sf.Tag

		fieldname := tag.Get("field")
		if fieldname == "" {
			continue
		}
		if what != "" {
			what += ","
		}
		what += fieldname
		dest = append(dest, field.Addr().Interface())

		switch field.Kind() {
		case reflect.Int:
			v := field.Int()
			if v > 0 {
				if where != "" {
					where += " and "
				}
				where += fmt.Sprintf("%s = %d", fieldname, v)
			}
		case reflect.Int8:
			v := field.Int()
			if v > 0 {
				if where != "" {
					where += " and "
				}
				where += fmt.Sprintf("%s = %d", fieldname, v)
			}
		case reflect.Int16:
			v := field.Int()
			if v > 0 {
				if where != "" {
					where += " and "
				}
				where += fmt.Sprintf("%s = %d", fieldname, v)
			}
		case reflect.Int32:
			v := field.Int()
			if v > 0 {
				if where != "" {
					where += " and "
				}
				where += fmt.Sprintf("%s = %d", fieldname, v)
			}
		case reflect.Int64:
			v := field.Int()
			if v > 0 {
				if where != "" {
					where += " and "
				}
				where += fmt.Sprintf("%s = %d", fieldname, v)
			}
		case reflect.Float32:
			v := field.Float()
			if v > 0 {
				if where != "" {
					where += " and "
				}
				where += fmt.Sprintf("%s = %f", fieldname, v)
			}
		case reflect.Float64:
			v := field.Float()
			if v > 0 {
				if where != "" {
					where += " and "
				}
				where += fmt.Sprintf("%s = %f", fieldname, v)
			}
		case reflect.String:
			v := field.String()
			if v != "" {
				if where != "" {
					where += " and "
				}
				w1, _ := mysql.MysqlFilter("%s = '%s'", fieldname, v)
				where += w1
			}
		case reflect.Slice:
			if "[]uint8" != field.Type().String() {
				panic("不支持这个格式")
			}
		default:
			panic("不支持这个格式")
		}
	}

	if o.WhereStr == "" {
		o.WhereStr = where
	}

	o.WhatStr = what
	sql := o.pasreSql()
	err := o.Mysqlclint.GetRead().MysqlSetOneData(dest, sql)
	if err != nil {
		panic(err)
	}
}

func (o *Orm) Fitle(format string, i ...interface{}) *Orm {
	var err error
	o.WhereStr, err = mysql.MysqlFilter(format, i...)
	if err != nil {
		panic(err)
	}
	return o
}

func (o *Orm) FitleStruct(md Model) *Orm {
	o.Table = md.TableInfo()
	val := reflect.ValueOf(md)
	ind := reflect.Indirect(val)
	where := ""
	for i := 0; i < ind.NumField(); i++ {
		field := ind.Field(i)
		sf := ind.Type().Field(i)
		tag := sf.Tag
		fieldname := tag.Get("field")
		if fieldname == "" {
			continue
		}

		switch field.Kind() {
		case reflect.Int:
			v := field.Int()
			if v > 0 {
				if where != "" {
					where += " and "
				}
				where += fmt.Sprintf("%s = %d", fieldname, v)
			}
		case reflect.Int8:
			v := field.Int()
			if v > 0 {
				if where != "" {
					where += " and "
				}
				where += fmt.Sprintf("%s = %d", fieldname, v)
			}
		case reflect.Int16:
			v := field.Int()
			if v > 0 {
				if where != "" {
					where += " and "
				}
				where += fmt.Sprintf("%s = %d", fieldname, v)
			}
		case reflect.Int32:
			v := field.Int()
			if v > 0 {
				if where != "" {
					where += " and "
				}
				where += fmt.Sprintf("%s = %d", fieldname, v)
			}
		case reflect.Int64:
			v := field.Int()
			if v > 0 {
				if where != "" {
					where += " and "
				}
				where += fmt.Sprintf("%s = %d", fieldname, v)
			}
		case reflect.Float32:
			v := field.Float()
			if v > 0 {
				if where != "" {
					where += " and "
				}
				where += fmt.Sprintf("%s = %f", fieldname, v)
			}
		case reflect.Float64:
			v := field.Float()
			if v > 0 {
				if where != "" {
					where += " and "
				}
				where += fmt.Sprintf("%s = %f", fieldname, v)
			}
		case reflect.String:
			v := field.String()
			if v != "" {
				if where != "" {
					where += " and "
				}
				w1, _ := mysql.MysqlFilter("%s = '%s'", fieldname, v)
				where += w1
			}
		case reflect.Slice:
			if "[]uint8" != field.Type().String() {
				panic("不支持这个格式")
			}
		default:
			panic("不支持这个格式")
		}
	}
	if where != "" {
		o.WhereStr = where
	}
	return o

}

func (o *Orm) All(ml interface{}) {
	t := reflect.TypeOf(ml)
	v := reflect.ValueOf(ml)
	ind := reflect.Indirect(v)

	if t.Kind() != reflect.Ptr {
		panic("ml must be Ptr")
	}

	slice := ind

	if t.Elem().Elem().Kind() == reflect.Ptr {
		d := reflect.New(t.Elem().Elem().Elem())
		if tabler, ok := d.Interface().(Model); ok {
			o.Table = tabler.TableInfo()
		} else {
			panic("model must have function TableName")
		}

		dind := reflect.Indirect(d)
		dindlen := dind.NumField()
		fslice := make([]int, dindlen)
		var fslicei int = 0

		what := ""
		for i := 0; i < dindlen; i++ {
			sf := dind.Type().Field(i)
			tag := sf.Tag

			fieldname := tag.Get("field")
			if fieldname == "" {
				continue
			} else {
				fslice[fslicei] = i
				fslicei ++
			}
			if what != "" {
				what += ","
			}
			what += fieldname
			//dest = append(dest, field.Addr().Interface())
		}

		o.WhatStr = what
		sql := o.pasreSql()
		rows, err := o.Mysqlclint.GetRead().QueryRows(sql)
		if err != nil {
			panic(err)
		}

		cols, _ := rows.Columns()
		lennum := len(cols)
		for rows.Next() {
			d1 := reflect.New(t.Elem().Elem().Elem())
			ind1 := reflect.Indirect(d1)
			dest := make([]interface{}, 0)
			for i := 0; i < lennum; i++ {
				field := ind1.Field(fslice[i])
				dest = append(dest, field.Addr().Interface())
			}
			rows.Scan(dest...)
			slice = reflect.Append(slice, d1)
		}
		rows.Close()
		ind.Set(slice)

	} else {

		d := reflect.New(t.Elem().Elem())
		if tabler, ok := d.Interface().(Model); ok {
			o.Table = tabler.TableInfo()
		} else {
			panic("model must have function TableName")
		}

		dind := reflect.Indirect(d)
		dindlen := dind.NumField()
		fslice := make([]int, dindlen)
		var fslicei int = 0

		what := ""
		for i := 0; i < dindlen; i++ {
			sf := dind.Type().Field(i)
			tag := sf.Tag

			fieldname := tag.Get("field")

			if fieldname == "" {
				continue
			} else {
				fslice[fslicei] = i
				fslicei ++
			}

			if what != "" {
				what += ","
			}
			what += fieldname
			//dest = append(dest, field.Addr().Interface())
		}

		o.WhatStr = what

		sql := o.pasreSql()
		rows, err := o.Mysqlclint.GetRead().QueryRows(sql)
		if err != nil {
			panic(err)
		}

		cols, _ := rows.Columns()
		lennum := len(cols)

		for rows.Next() {
			d1 := reflect.New(t.Elem().Elem())
			ind1 := reflect.Indirect(d1)
			dest := make([]interface{}, 0)
			for i := 0; i < lennum; i++ {
				field := ind1.Field(fslice[i])
				dest = append(dest, field.Addr().Interface())
			}
			rows.Scan(dest...)
			slice = reflect.Append(slice, d1.Elem())
		}

		rows.Close()
		ind.Set(slice)
	}
}

package helper

import (
	"errors"
	"fmt"

	"reflect"
	"strings"

	"gitlab.com/go-helpers/simple-sql-helper/utility"
)

// NamedExecHelper ... a simple query builder for single table
//doesnt take joins into consideration
type NamedExecHelper struct {
	tablename  string
	model      interface{}            // cannot be a pointer
	mapValues  map[string]interface{} // tag-value of struct
	ignoreTags []string               // tags to be ignored
	AliasMap   map[string]interface{}
}

// NewNamedExecHelper ...
func NewNamedExecHelper(tablename string, model interface{}) (*NamedExecHelper, error) {
	if kind := reflect.ValueOf(model).Kind(); kind == reflect.Ptr {
		return nil, errors.New("model cannot be a pointer")
	}
	me := &NamedExecHelper{
		model:      model,
		tablename:  tablename,
		ignoreTags: []string{"-"},
		mapValues:  map[string]interface{}{},
		AliasMap:   map[string]interface{}{},
	}
	return me, nil
}

// SetIgnoreTags ...sets db tags to ignore.
// use cases are for time where a default value is provided by the db, eg datetime current
func (me *NamedExecHelper) SetIgnoreTags(tags ...string) *NamedExecHelper {
	me.ignoreTags = append(me.ignoreTags, tags...)
	return me
}

// InsertQuery builds an insert query...
func (me *NamedExecHelper) InsertQuery() string {
	var (
		keys  []string
		alias []string
	)
	me.setMapValues()
	for k := range me.mapValues {
		keys = append(keys, k)
		alias = append(alias, fmt.Sprintf(":%s", k))
	}
	keyS := strings.Join(keys, ",")
	aliasS := strings.Join(alias, ",")
	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", me.tablename, keyS, aliasS)
}

// SelectQuery ...select everything by default
func (me *NamedExecHelper) SelectQuery(condition *ConditionBuilder, fields ...string) string {
	var (
		sField     = "*"
		sCondition string
	)
	if len(fields) > 0 {
		//use provided fields
		sField = strings.Join(fields, ",")
	}
	if condition != nil {
		sCondition = condition.UnnamedConditionString()
	}

	return fmt.Sprintf("SELECT %s FROM %s %s;", sField, me.tablename, sCondition)
}

// UpsertQuery /.. returns a named exec
// if fields are provided, only upsert fields
// else upsert entire model within mapValues
func (me *NamedExecHelper) UpsertQuery(condition *ConditionBuilder, fields ...string) string {
	var (
		setters    []string
		settersS   = ""
		conditionS = ""
	)

	me.setMapValues()

	// 	m := map[string]interface{}{"city": "Johannesburg"}
	// result, err := db.NamedExec(`SELECT * FROM place WHERE city=:city`, m)
	if len(fields) > 0 {
		for _, field := range fields {
			if val, ok := me.mapValues[field]; ok {
				setters = append(setters, fmt.Sprintf("%s=:%s", field, field))
				me.AliasMap[field] = val
			}
			settersS = strings.Join(setters, ",")

		}
	} else {
		for tmp, val := range me.mapValues {
			setters = append(setters, fmt.Sprintf("%s=:%s", tmp, tmp))
			me.AliasMap[tmp] = val
		}
		settersS = strings.Join(setters, ",")
	}

	if condition != nil {
		conditionS = condition.NamedConditionString()
		for k, v := range condition.AliasMap {
			me.AliasMap[k] = v
		}
	}

	return fmt.Sprintf("UPDATE %s SET %s %s;", me.tablename, settersS, conditionS)
}

func (me *NamedExecHelper) setMapValues() {
	v := reflect.ValueOf(me.model)
	typeOf := v.Type()
	for i := 0; i < v.NumField(); i++ {
		dbtag := typeOf.Field(i).Tag
		tag := dbtag.Get("db")
		if len(tag) > 0 && tag != "-" {
			if found, _ := utility.Find(tag, me.ignoreTags); !found {
				me.mapValues[tag] = v.Field(i).Interface()
			}
		}
	}
}

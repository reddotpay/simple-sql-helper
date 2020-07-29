package helper

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type db struct {
	Name string     `db:"name"`
	Time *time.Time `db:"time"`
}

func Test_ExecInsertHelper(t *testing.T) {
	d := db{}
	assert := assert.New(t)
	helper, e := NewNamedExecHelper("dbTable", d)
	assert.NoError(e)
	helper.SetIgnoreTags("time")
	s := helper.InsertQuery()
	fmt.Println("s", s)
}

func Test_SelectQuery(t *testing.T) {

	d := db{}
	assert := assert.New(t)
	helper, e := NewNamedExecHelper("dbTable", d)
	assert.NoError(e)
	helper.SetIgnoreTags("time")
	s := helper.SelectQuery(nil, "name")
	fmt.Println("s", s)

}
func Test_SelectQueryWithCond(t *testing.T) {

	d := db{}
	assert := assert.New(t)
	helper, e := NewNamedExecHelper("dbTable", d)
	assert.NoError(e)
	helper.SetIgnoreTags("time")

	c1 := ConditionEqual("name", "david")
	c2 := ConditionNotEqual("age", 15)
	cb := NewConditionBuilder(c1)
	cb.And(c2)

	s := helper.SelectQuery(cb)
	fmt.Println("s", s)
}

func Test_Upsert(t *testing.T) {
	d := db{}
	assert := assert.New(t)
	helper, e := NewNamedExecHelper("dbTable", d)
	assert.NoError(e)
	helper.SetIgnoreTags("time")

	c1 := ConditionEqual("name", "david")
	c2 := ConditionNotEqual("age", 15)
	cb := NewConditionBuilder(c1)
	cb.And(c2)

	s := helper.UpsertQuery(cb, "name")
	aliases := helper.AliasMap
	fmt.Println(s, aliases)
}
func Test_UpsertDefault(t *testing.T) {
	d := db{}
	assert := assert.New(t)
	helper, e := NewNamedExecHelper("dbTable", d)
	assert.NoError(e)
	// helper.SetIgnoreTags("time")

	c1 := ConditionEqual("name", "david")
	c2 := ConditionNotEqual("age", 15)
	cb := NewConditionBuilder(c1)
	cb.And(c2)

	s := helper.UpsertQuery(cb)
	aliases := helper.AliasMap
	fmt.Println(s, aliases)
}

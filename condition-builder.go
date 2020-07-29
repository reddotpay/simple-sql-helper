package helper

import (
	"fmt"
	"strings"
)

// ConditionBuilder /...
type ConditionBuilder struct {
	unnamedConditions []string
	namedConditions   []string
	AliasMap          map[string]interface{}
	Values            []interface{}
	conditions        []*Condition
}

// NewConditionBuilder ...
func NewConditionBuilder(c *Condition) *ConditionBuilder {
	me := &ConditionBuilder{
		AliasMap: map[string]interface{}{},
	}
	return me.build(c, "WHERE")
}

// And ...
func (me *ConditionBuilder) And(c *Condition) *ConditionBuilder {
	return me.build(c, "AND")
}

// Or ...
func (me *ConditionBuilder) Or(c *Condition) *ConditionBuilder {
	return me.build(c, "OR")
}

// ConditionBuilder .. where verb = WHERE, AND, OR
func (me *ConditionBuilder) build(c *Condition, verb string) *ConditionBuilder {

	// check current size of condition store
	i := len(me.conditions)

	//add c to the collection of conditions
	me.conditions = append(me.conditions, c)
	alias := fmt.Sprintf("%v", i)

	// build unnamed condition name=?
	unnamedCondition := c.string()
	unnamedCondition = fmt.Sprintf("%v %v ", verb, unnamedCondition)
	me.unnamedConditions = append(me.unnamedConditions, unnamedCondition)

	//set named conditions
	namedCondition := c.string(i)
	namedCondition = fmt.Sprintf("%v %v ", verb, namedCondition)
	me.namedConditions = append(me.namedConditions, namedCondition)
	//set alias
	me.AliasMap[alias] = c.Value
	me.Values = append(me.Values, c.Value)

	return me
}

// NamedConditionString ...
func (me *ConditionBuilder) NamedConditionString() string {
	return strings.Join(me.namedConditions, "")
}

// UnnamedConditionString ...
func (me *ConditionBuilder) UnnamedConditionString() string {
	return strings.Join(me.unnamedConditions, "")
}

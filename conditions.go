package helper

import "fmt"

// Condition ...
type Condition struct {
	named    bool
	Key      string
	Operator string
	Value    interface{}
}

// ConditionEqual ///
func ConditionEqual(key string, value interface{}) *Condition {
	return &Condition{
		Key:      key,
		Value:    value,
		Operator: "=",
	}
}

// ConditionLessThan ...
func ConditionLessThan(key string, value interface{}) *Condition {
	return &Condition{
		Key:      key,
		Value:    value,
		Operator: "<",
	}
}

// ConditionNotEqual ///
func ConditionNotEqual(key string, value interface{}) *Condition {
	return &Condition{
		Key:      key,
		Value:    value,
		Operator: "!=",
	}
}

func (me *Condition) string(i ...int) string {
	if len(i) > 0 {
		// key=:named_value
		return fmt.Sprintf("%v%v:%v", me.Key, me.Operator, i[0])
	}
	//key=?
	return fmt.Sprintf("%v%v?", me.Key, me.Operator)

}

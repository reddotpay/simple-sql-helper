package helper

import (
	"fmt"
	"testing"
)

func Test_ConditionBuilder(t *testing.T) {
	c1 := ConditionEqual("name", "david")
	c2 := ConditionEqual("name", "david")
	c3 := ConditionEqual("name", "david")
	cb := NewConditionBuilder(c1)
	cb.And(c2)
	cb.Or(c3)
	conditionString := cb.NamedConditionString()
	unnamed := cb.UnnamedConditionString()
	fmt.Println(conditionString)
	fmt.Println(unnamed)
}

package internal

import (
	"github.com/beego/beego/v2/client/orm"
	"strings"
)

type Condition struct {
	*orm.Condition
}

func NewCond() *Condition {
	return &Condition{orm.NewCondition()}
}

func (c *Condition) And(expr string, args ...interface{}) *Condition {
	return &Condition{Condition: c.Condition.And(expr, args...)}
}

func (c *Condition) AndNot(expr string, args ...interface{}) *Condition {
	return &Condition{Condition: c.Condition.AndNot(expr, args...)}
}

func (c *Condition) AndCond(cond *Condition) *Condition {
	return &Condition{Condition: c.Condition.AndCond(cond.Condition)}
}

func (c *Condition) AndNotCond(cond *Condition) *Condition {
	return &Condition{Condition: c.Condition.AndNotCond(cond.Condition)}
}

func (c *Condition) Or(expr string, args ...interface{}) *Condition {
	return &Condition{Condition: c.Condition.Or(expr, args...)}
}

func (c *Condition) OrNot(expr string, args ...interface{}) *Condition {
	return &Condition{Condition: c.Condition.OrNot(expr, args...)}
}

func (c *Condition) OrCond(cond *Condition) *Condition {
	return &Condition{Condition: c.Condition.OrCond(cond.Condition)}
}

func (c *Condition) OrNotCond(cond *Condition) *Condition {
	return &Condition{Condition: c.Condition.OrNotCond(cond.Condition)}
}

func (c *Condition) IsEmpty() bool {
	return c.Condition.IsEmpty()
}

type Field string

type AliasField string

func (f Field) Exact() string {
	return string(f)
}

func (f Field) Greater() string {
	return string(f) + "__gt"
}

func (f Field) GreaterEqual() string {
	return string(f) + "__gte"
}

func (f Field) Less() string {
	return string(f) + "__lt"
}

func (f Field) LessEqual() string {
	return string(f) + "__lte"
}

func (f Field) In() string {
	return string(f) + "__in"
}

func (f Field) Contains() string {
	return string(f) + "__icontains"
}

func (f AliasField) In(n int) string {
	b := strings.Builder{}
	b.WriteString(string(f))
	b.WriteString(" in (")
	for i := 0; i < n; i++ {
		if i == 0 {
			b.WriteString("?")
		} else {
			b.WriteString(",?")
		}
	}
	b.WriteString(")")
	return b.String()
}

func (f Field) Desc() Field {
	return "-" + f
}

func (f Field) Alias() AliasField {
	return AliasField("T0." + f)
}

func (f Field) AliasT1() AliasField {
	return AliasField("T1." + f)
}

func (f Field) AliasT2() AliasField {
	return AliasField("T2." + f)
}

func (f Field) AliasT3() AliasField {
	return AliasField("T3." + f)
}

func (f AliasField) Greater() string {
	return string(f) + " > ? "
}

func (f AliasField) GreaterEqual() string {
	return string(f) + " >= ? "
}

func (f AliasField) Less() string {
	return string(f) + " < ? "
}

func (f AliasField) LessEqual() string {
	return string(f) + " <= ? "
}

func (f AliasField) Contains() string {
	return string(f) + " LIKE ? "
}

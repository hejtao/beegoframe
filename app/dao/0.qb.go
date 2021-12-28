package dao

import (
	"beegoframe/app/dao/internal"
	"reflect"
)

type joinBuilder struct {
	builder internal.QueryBuilder
}

type onBuilder struct {
	builder internal.QueryBuilder
	alias   string
}

type whereBuilder struct {
	builder internal.QueryBuilder
	args    []interface{}
}

type andOrBuilder struct {
	builder internal.QueryBuilder
	args    []interface{}
}

type orderBuilder struct {
	builder internal.QueryBuilder
	args    []interface{}
}

type offsetBuilder struct {
	builder internal.QueryBuilder
	args    []interface{}
}

type limitBuilder struct {
	builder internal.QueryBuilder
	args    []interface{}
}

func (r *onBuilder) On(foreignKey internal.Field) *joinBuilder {
	fk := string(foreignKey)
	return &joinBuilder{
		builder: r.builder.
			On("T0." + fk + " = " + r.alias + ".id")}
}

func (r *joinBuilder) InnerJoin(d Dao, alias string) *onBuilder {
	tb := d.GetTable()
	return &onBuilder{
		alias: alias,
		builder: r.builder.
			InnerJoin(tb + " " + alias)}
}

func (r *joinBuilder) LeftJoin(d Dao, alias string) *onBuilder {
	tb := d.GetTable()
	return &onBuilder{
		alias:   alias,
		builder: r.builder.LeftJoin(tb + " " + alias)}
}

func (r *joinBuilder) Where(cond string, arg interface{}) *whereBuilder {
	return &whereBuilder{
		builder: r.builder.Where(cond),
		args:    []interface{}{arg},
	}
}

func (r *whereBuilder) And(cond string, arg interface{}) *andOrBuilder {
	r.args = append(r.args, arg)
	return &andOrBuilder{
		builder: r.builder.And(cond),
		args:    r.args,
	}
}

func (r *whereBuilder) Or(cond string, arg interface{}) *andOrBuilder {
	r.args = append(r.args, arg)
	return &andOrBuilder{
		builder: r.builder.Or(cond),
		args:    r.args,
	}
}

func (r *whereBuilder) Limit(num int) *limitBuilder {
	return &limitBuilder{
		builder: r.builder.Limit(num),
		args:    r.args,
	}
}

func (r *andOrBuilder) And(cond string, arg interface{}) *andOrBuilder {
	r.args = append(r.args, arg)
	return &andOrBuilder{
		builder: r.builder.And(cond),
		args:    r.args,
	}
}

func (r *andOrBuilder) Or(cond string, arg interface{}) *andOrBuilder {
	r.args = append(r.args, arg)
	return &andOrBuilder{
		builder: r.builder.And(cond),
		args:    r.args,
	}
}

func (r *andOrBuilder) Offset(num int) *offsetBuilder {
	return &offsetBuilder{
		builder: r.builder.Offset(num),
		args:    r.args,
	}
}

func (r *andOrBuilder) Limit(num int) *limitBuilder {
	return &limitBuilder{
		builder: r.builder.Limit(num),
		args:    r.args,
	}
}

func (r *andOrBuilder) OrderBy(fields ...internal.Field) *orderBuilder {
	sql := ""
	for _, field := range fields {
		sql += ", " + string(field)
	}
	return &orderBuilder{
		builder: r.builder.OrderBy(sql[1:]),
		args:    r.args,
	}
}

func (r *orderBuilder) Offset(num int) *offsetBuilder {
	return &offsetBuilder{
		builder: r.builder.Offset(num),
		args:    r.args,
	}
}

func (r *offsetBuilder) Limit(num int) *limitBuilder {
	return &limitBuilder{
		builder: r.builder.Offset(num),
		args:    r.args,
	}
}

func (r *limitBuilder) Rows(container interface{}) (int64, error) {
	if reflect.TypeOf(container).Kind() != reflect.Ptr {
		return 0, errQueryRowsWithWrongContainerType
	}
	fields := getSelectFields(container)
	sql := internal.NewQueryBuilder().Select(fields...).String() + " " + r.builder.String()
	return internal.DB.Raw(sql, r.args...).QueryRows(container)
}

// Package dao
// ===========================
// This file is auto-generated
// ===========================
package dao

import (
	"beegoframe/app/dao/internal"
	"beegoframe/app/model"
	"beegoframe/pkg/encoding/json"
	"time"
)

type accountDao struct {
	setter internal.QuerySetter
	tx     Tx
	data   interface{}
	table  string
	Fields *accountFields
}

type accountFields struct {
	Id         internal.Field
	Deleted    internal.Field
	CreateTime internal.Field
	UpdateTime internal.Field
	Email      internal.Field
	Phone      internal.Field
	Password   internal.Field
}

var Account = &accountDao{
	setter: internal.DB.QueryTable((*model.Account)(nil)),
	table:  "account",
	Fields: &accountFields{
		Id:         "id",
		Deleted:    "deleted",
		CreateTime: "create_time",
		UpdateTime: "update_time",
		Email:      "email",
		Phone:      "phone",
		Password:   "password",
	},
}

func (r *accountDao) GetTable() string {
	return r.table
}

func (r *accountDao) Filter(expr string, args ...interface{}) *accountDao {
	return &accountDao{
		setter: r.setter.Filter(expr, args...),
		tx:     r.tx,
		data:   r.data,
		Fields: r.Fields,
	}
}

func (r *accountDao) FilterRaw(expr string, arg string) *accountDao {
	return &accountDao{
		setter: r.setter.FilterRaw(expr, arg),
		tx:     r.tx,
		data:   r.data,
		Fields: r.Fields,
	}
}

func (r *accountDao) Exclude(expr string, args ...interface{}) *accountDao {
	return &accountDao{
		setter: r.setter.Exclude(expr, args...),
		tx:     r.tx,
		data:   r.data,
		Fields: r.Fields,
	}
}

func (r *accountDao) SetCond(cond *internal.Condition) *accountDao {
	return &accountDao{
		setter: r.setter.SetCond(cond.Condition),
		tx:     r.tx,
		data:   r.data,
		Fields: r.Fields,
	}
}

func (r *accountDao) GetCond() *internal.Condition {
	return &internal.Condition{Condition: r.setter.GetCond()}
}

// Pagination set the limit and offset for the dao
func (r *accountDao) Pagination(pageNum, pageSize int) *accountDao {
	if pageNum < 1 {
		pageNum = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	return &accountDao{
		setter: r.setter.Limit(pageSize).Offset((pageNum - 1) * pageSize),
		tx:     r.tx,
		data:   r.data,
		Fields: r.Fields,
	}
}

func (r *accountDao) GroupBy(expr ...internal.Field) *accountDao {
	var fields []string
	for _, v := range expr {
		fields = append(fields, string(v))
	}
	return &accountDao{
		setter: r.setter.GroupBy(fields...),
		tx:     r.tx,
		data:   r.data,
		Fields: r.Fields,
	}
}

func (r *accountDao) OrderBy(expr ...internal.Field) *accountDao {
	var fields []string
	for _, v := range expr {
		fields = append(fields, string(v))
	}
	return &accountDao{
		setter: r.setter.OrderBy(fields...),
		tx:     r.tx,
		data:   r.data,
		Fields: r.Fields,
	}
}

func (r *accountDao) ForceIndex(indexes ...string) *accountDao {
	return &accountDao{
		setter: r.setter.ForceIndex(indexes...),
		tx:     r.tx,
		data:   r.data,
		Fields: r.Fields,
	}
}

func (r *accountDao) UseIndex(indexes ...string) *accountDao {
	return &accountDao{
		setter: r.setter.UseIndex(indexes...),
		tx:     r.tx,
		data:   r.data,
		Fields: r.Fields,
	}
}

func (r *accountDao) IgnoreIndex(indexes ...string) *accountDao {
	return &accountDao{
		setter: r.setter.IgnoreIndex(indexes...),
		tx:     r.tx,
		data:   r.data,
		Fields: r.Fields,
	}
}

func (r *accountDao) Distinct() *accountDao {
	return &accountDao{
		setter: r.setter.Distinct(),
		tx:     r.tx,
		data:   r.data,
		Fields: r.Fields,
	}
}

func (r *accountDao) ForUpdate() *accountDao {
	return &accountDao{
		setter: r.setter.ForUpdate(),
		tx:     r.tx,
		data:   r.data,
		Fields: r.Fields,
	}
}

func (r *accountDao) Count() (int64, error) {
	return r.setter.Count()
}

func (r *accountDao) Exist() bool {
	return r.setter.Exist()
}

// SetData set the data for the dao update or insert
func (r *accountDao) SetData(data interface{}) *accountDao {
	return &accountDao{
		setter: r.setter,
		tx:     r.tx,
		data:   data,
		Fields: r.Fields,
	}
}

// SetTx set the transaction for the dao
func (r *accountDao) SetTx(tx internal.Tx) *accountDao {
	return &accountDao{
		setter: r.setter,
		tx:     tx,
		data:   r.data,
		Fields: r.Fields,
	}
}

// Update is the dao update operation, the transaction param is optional
func (r *accountDao) Update() (int64, error) {
	cond := r.setter.GetCond()
	if cond == nil || cond.IsEmpty() {
		return 0, errUpdateWithEmptyCond
	}
	row := internal.NewRow()
	if err := json.Convert(r.data, &row); err != nil {
		return 0, err
	}
	if len(row) == 0 {
		return 0, errUpdateWithEmptyData
	}
	delete(row, string(r.Fields.Id))
	row[string(r.Fields.UpdateTime)] = time.Now()
	if r.tx != nil {
		return r.tx.QueryTable((*model.Account)(nil)).SetCond(cond).Update(row)
	}
	return r.setter.Update(row)
}

// Delete is the dao delete operation, the transaction param is optional
func (r *accountDao) Delete() (int64, error) {
	cond := r.setter.GetCond()
	if cond == nil || cond.IsEmpty() {
		return 0, errDeleteWithEmptyCond
	}
	if r.tx != nil {
		return r.tx.QueryTable((*model.Account)(nil)).SetCond(cond).Delete()
	}
	return r.setter.Delete()
}

// Insert is the dao insert operation, the transaction param is optional
func (r *accountDao) Insert() (int64, error) {
	if _, ok := r.data.(*model.Account); !ok {
		var row model.Account
		if err := json.Convert(r.data, &row); err != nil {
			return 0, err
		}
		r.data = &row
	}
	if r.data == nil {
		return 0, errInsertWithEmptyData
	}
	if r.tx != nil {
		return r.tx.Insert(r.data)
	}
	return internal.DB.Insert(r.data)
}

// FirstRow queries the first row data
// container is map[string]interface{}, all fields are selected
// container is *struct, the struct json tags specify the selected fields
func (r *accountDao) FirstRow(container interface{}) error {
	if err := checkPtr(container); err != nil {
		return err
	}
	cond := r.setter.GetCond()
	if cond == nil || cond.IsEmpty() {
		return errQueryWithEmptyCond
	}
	tags := getJsonTags(container)
	rows := internal.NewRows()
	if _, err := r.setter.Limit(1).Values(&rows, tags...); err != nil {
		return err
	}
	if len(rows) == 0 {
		return errQueryRowNotFound
	}
	if err := json.Convert(rows[0], container); err != nil {
		return err
	}
	return nil
}

// LastRow queries the last row data
// container is map[string]interface{}, all fields are selected
// container is *struct, the struct json tags specify the selected fields
func (r *accountDao) LastRow(container interface{}) error {
	if err := checkPtr(container); err != nil {
		return err
	}
	cond := r.setter.GetCond()
	if cond == nil || cond.IsEmpty() {
		return errQueryWithEmptyCond
	}
	tags := getJsonTags(container)
	rows := internal.NewRows()
	if _, err := r.setter.Limit(1).OrderBy(string(r.Fields.Id.Desc())).Values(&rows, tags...); err != nil {
		return err
	}
	if len(rows) == 0 {
		return errQueryRowNotFound
	}
	if err := json.Convert(rows[0], container); err != nil {
		return err
	}
	return nil
}

// Rows query multiple rows data
// container is []map[string]interface{}, all fields are selected
// container is []struct or []*struct, the struct json tags specify the selected fields
func (r *accountDao) Rows(container interface{}) (int64, error) {
	if err := checkSlicePtr(container); err != nil {
		return 0, err
	}
	cond := r.setter.GetCond()
	if cond == nil || cond.IsEmpty() {
		return 0, errQueryWithEmptyCond
	}
	tags := getJsonTags(container)
	rows := internal.NewRows()
	n, err := r.setter.Values(&rows, tags...)
	if err != nil {
		return 0, err
	}
	if err := json.Convert(rows, container); err != nil {
		return 0, err
	}
	return n, nil
}

// Column query one column data
func (r *accountDao) Column(field internal.Field) []interface{} {
	column := internal.NewColumn()
	cond := r.setter.GetCond()
	if cond != nil && !cond.IsEmpty() {
		r.setter.ValuesFlat(&column, string(field))
	}
	return column
}

func (r *accountDao) LeftJoin(d Dao, alias string) *onBuilder {
	table := d.GetTable()
	return &onBuilder{
		alias: alias,
		builder: internal.NewQueryBuilder().From(r.table + " T0").
			LeftJoin(table + " " + alias)}
}

func (r *accountDao) InnerJoin(d Dao, alias string) *onBuilder {
	table := d.GetTable()
	return &onBuilder{
		alias: alias,
		builder: internal.NewQueryBuilder().From(r.table + " T0").
			InnerJoin(table + " " + alias)}
}

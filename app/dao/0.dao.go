package dao

type Dao interface {
	GetTable() string
	Insert() (int64, error)
	Delete() (int64, error)
	Update() (int64, error)
	FirstRow(container interface{}) error
	LastRow(container interface{}) error
	Rows(container interface{}) (int64, error)
	Count() (int64, error)
	Exist() bool
}

package internal

import (
	"context"
	"github.com/beego/beego/v2/client/orm"
)

type Tx interface {
	orm.TxOrmer
}

func DoTx(task func(ctx context.Context, tx Tx) error) error {
	return DB.DoTx(func(ctx2 context.Context, tx2 orm.TxOrmer) error {
		return task(ctx2, tx2)
	})
}

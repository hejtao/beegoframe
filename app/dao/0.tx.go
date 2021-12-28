package dao

import (
	"beegoframe/app/dao/internal"
	"context"
)

type Tx interface {
	internal.Tx
}

func DoTx(task func(ctx context.Context, tx Tx) error) error {
	return internal.DoTx(func(ctx2 context.Context, tx2 internal.Tx) error {
		return task(ctx2, tx2)
	})
}

func BeginTx(ctx ...context.Context) (Tx, error) {
	if len(ctx) > 0 {
		return internal.DB.BeginWithCtx(ctx[0])
	}
	return internal.DB.Begin()
}

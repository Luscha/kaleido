package storage

import "github.pitagora/pkg/ent"

type QueryOption func(*queryOptions)

type queryOptions struct {
	tx *ent.Tx
}

func WithTransactionalClient(tx *ent.Tx) QueryOption {
	return func(o *queryOptions) {
		o.tx = tx
	}
}

func DefaultQueryOption() *queryOptions {
	return &queryOptions{}
}

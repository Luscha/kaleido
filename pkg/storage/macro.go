package storage

import (
	"context"

	"github.pitagora/pkg/ent"
	"github.pitagora/pkg/ent/macro"
)

func (c *Client) GetMacros(ctx context.Context, names []string, opts ...QueryOption) ([]*ent.Macro, error) {
	o := DefaultQueryOption()
	for _, opt := range opts {
		opt(o)
	}

	query := c.GetClient(o).Macro.Query().
		Where(
			macro.NameIn(names...),
		)
	result, err := query.All(ctx)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) ListMacros(ctx context.Context, opts ...QueryOption) ([]*ent.Macro, error) {
	o := DefaultQueryOption()
	for _, opt := range opts {
		opt(o)
	}

	query := c.GetClient(o).Macro.Query()
	result, err := query.All(ctx)

	if err != nil {
		return nil, err
	}

	return result, nil
}

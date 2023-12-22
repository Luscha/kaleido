package storage

import (
	"context"

	"github.pitagora/pkg/ent"
	"github.pitagora/pkg/ent/procedure"
)

func (c *Client) GetProcedure(ctx context.Context, name string, opts ...QueryOption) (*ent.Procedure, error) {
	o := DefaultQueryOption()
	for _, opt := range opts {
		opt(o)
	}

	query := c.GetClient(o).Procedure.Query().
		Where(
			procedure.Name(name),
		)
	result, err := query.Only(ctx)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) ListProcedures(ctx context.Context, opts ...QueryOption) ([]*ent.Procedure, error) {
	o := DefaultQueryOption()
	for _, opt := range opts {
		opt(o)
	}

	query := c.GetClient(o).Procedure.Query()
	result, err := query.All(ctx)

	if err != nil {
		return nil, err
	}

	return result, nil
}

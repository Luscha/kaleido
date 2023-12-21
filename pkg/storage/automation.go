package storage

import (
	"context"

	"github.pitagora/pkg/ent"
	"github.pitagora/pkg/ent/automation"
)

func (c *Client) GetAutomation(ctx context.Context, name string, opts ...QueryOption) (*ent.Automation, error) {
	o := DefaultQueryOption()
	for _, opt := range opts {
		opt(o)
	}

	query := c.GetClient(o).Automation.Query().
		Where(
			automation.Name(name),
		)
	result, err := query.Only(ctx)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) ListAutomations(ctx context.Context, opts ...QueryOption) ([]*ent.Automation, error) {
	o := DefaultQueryOption()
	for _, opt := range opts {
		opt(o)
	}

	query := c.GetClient(o).Automation.Query()
	result, err := query.All(ctx)

	if err != nil {
		return nil, err
	}

	return result, nil
}

package storage

import (
	"context"

	"entgo.io/ent/dialect"
	"github.com/pkg/errors"
	"github.pitagora/pkg/ent"
	"gitlab.com/technity/go-x/pkg/connection"

	_ "github.com/lib/pq"
)

type ClientFactory struct {
}

func (cf *ClientFactory) Spawn(ctx context.Context, cfg *connection.DBConnectionConfig) (*Client, error) {
	return NewClient(ctx, cfg)
}

func NewClientFactoy() *ClientFactory {
	return &ClientFactory{}
}

type Client struct {
	db *ent.Client
}

var (
	ErrNotFound         = errors.New("storage: not found")
	ErrAlreadyExists    = errors.New("storage: already exists")
	ErrAlreadyDeleted   = errors.New("storage: already deleted")
	ErrInvalidPageToken = errors.New("storage: invalid page token")
	ErrInvalidFilter    = errors.New("storage: invalid filter")
)

func FromEntClient(ctx context.Context, c *ent.Client, cfg *connection.DBConnectionConfig) (*Client, error) {
	return &Client{db: c}, nil
}

func NewClient(ctx context.Context, cfg *connection.DBConnectionConfig) (*Client, error) {
	db, err := ent.Open(dialect.Postgres, cfg.Endpoint)
	if err != nil {
		return nil, err
	}

	if cfg.DebugMode {
		db = db.Debug()
	}

	return FromEntClient(ctx, db, cfg)
}

func (c *Client) Close() error {
	var closeErr error

	if err := c.db.Close(); err != nil {
		closeErr = err
	}

	return closeErr
}

func (c *Client) GetDB() *ent.Client {
	return c.db
}

func (c *Client) GetClient(opts *queryOptions) *ent.Client {
	if opts.tx != nil {
		return opts.tx.Client()
	}
	return c.db
}

func (c *Client) Tx(ctx context.Context, fn func(tx *ent.Tx) error) error {
	tx, err := c.db.Tx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()

	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = errors.Wrapf(err, "rolling back transaction: %v", rerr)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return errors.Wrapf(err, "committing transaction: %v", err)
	}

	return nil
}

package ISP

import "context"

type FatDbInterface[T any] interface {
	BatchGetItem(IDs ...int) ([]T, error)
	BatchGetItemWithContext(ctx context.Context, IDs ...int) ([]T, error)

	BatchPutItem(items ...T) error
	BatchPutItemWithContext(ctx context.Context, items ...T) error

	DeleteItem(ID int) error
	DeleteItemWithContext(ctx context.Context, ID int) error

	GetItem(ID int) (T, error)
	GetItemWithContext(ctx context.Context, ID int) (T, error)

	PutItem(item T) error
	PutItemWithContext(ctx context.Context, item T) error

	Query(query string, args ...interface{}) ([]T, error)
	QueryWithContext(ctx context.Context, query string, args ...interface{}) ([]T, error)

	UpdateItem(ID int, item T) error
	UpdateItemWithContext(ctx context.Context, ID int, item T) error
}

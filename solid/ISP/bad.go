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

/*
위 Interface의 모든 메서드를 구현해야 하는데, 이는 매우 비효율적이다.
GetItem과 GetItemWithContext 같은 메서드 쌍은 전반적이지는 않더라도 부분적으로 동일한 코드를 공유한다.
반면에, GetItem()의 사용자는 GetItemWithContext()를 사용하지 않을 것이다.
*/

package ISP

type myDb[T any] interface {
	GetItem(ID int) (T, error)
	PutItem(item T) error
}

type CacheV2[T any] struct {
	db myDb[T]
}

type User struct {
	ID   int
	Name string
}

func (c *CacheV2[User]) Get(key string) interface{} {
	_, _ = c.db.GetItem(0)

	return nil
}

/*
이처럼 Thin interface를 사용할 경우, 함수 시크니처가 좀 더 명확하고 유연해진다.
*/

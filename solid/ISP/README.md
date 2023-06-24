Interface Segregation Principle
---
> Clients should not be forced to depend upon interfaces that they do not use. - Robert C. Martin
> 
큰 덩어리의 Interface는 많은 메서드를 갖고 있으므로 이해하기 어렵다.
또한 이러한 Interface는 구현, 모듸, 스텁 등 사용하기 위해 많은 작업이 필요하다.
Interface가 변경될 경우에는 모든 사용자에게 미치는 파급 효과가 클 것이며,
이는 OCP 원칙을 위반하고 무수히 많은 변경의 분산 상황이 발생한다.

```go 
type FatDbInterface interface {
	BatchGetItem(IDs...int)([]Item, error)
	BatchGetItemWithContext(ctx context.Context, IDs...int)([]Item, error)
	
	BatchPutItem(items...Item)([]Item, error)
	BatchPutItemWithContext(ctx context.Context, items...Item)([]Item, error)
	
	BatchDeleteItem(IDs...int)([]Item, error)
	BatchDeleteItemWithContext(ctx context.Context, IDs...int)([]Item, error)
	
	GetItem(ID int)(Item, error)
	GetItemWithContext(ctx context.Context, ID int)(Item, error)
	
	PutItem(item Item)(Item, error)
	PutItemWithContext(ctx context.Context, item Item)(Item, error)
	
	Query(query string, args...interface{})([]Item, error)
	QueryWithContext(ctx context.Context, query string, args...interface{})([]Item, error)
	
	UpdateItem(item Item)(Item, error)
	UpdateItemWithContext(ctx context.Context, item Item)(Item, error)
}
```
위 Interface에서 `GetItem()`과, `GetItemWithContext()` 같은 메서드 쌍은 전반적이지는 않더라도 부분적으로 도잉ㄹ한 코드를 공유한다.
반면에, `GetItem()`의 사용자는 `GetItemWithContext()`를 사용하지 않는다.

이러한 경우에는 Interface를 분리하여 사용하는 것이 좋다.
```go
type myDB interface {
    GetItem(ID int)(Item, error)
    PutItem(item Item)(Item, error)
}
```
이처럼 새롭게 정의된 'Thin Interface'를 사용할 경우, 함수 시그니터가 좀 더 명확해지고 유연해진다. 이것이 ISP의 핵심이다.
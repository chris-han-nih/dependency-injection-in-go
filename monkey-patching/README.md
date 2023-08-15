Monkey patching
---
> Monkey patching은 일반적으로 프로그램의 런타임에 함수 또는 변수 등을 변경하는 것을 의미한다.
> 자동차 충돌 사고가 인체에 미치는 영향을 테스트 할 때 차 내부에 있을 사람을 crash test dummy로 교체(monkey patch)할 수 있다.

### Monkey patching의 장점
#### 1. Monkey patching을 통한 DI는 구현하기 매우 쉽다.
> 코드를 분리한다는 것은 서로 사용하거나 의존 관계에 있더라도 코드의 조각을 구분해 관리하는 것이다. 코드를 분리하면 테스트를 용이하게 할 수 있을 뿐만 아니라
> 코드를 각기 따로 발전시켜서 작은 그룹과 생각해볼 수 있는 단위로 제공하게 하고 이를 통해 개별적으로 코드의 다른 부분을 생각해볼 수 있는 기회를 제공한다.
> Monkey patching을 적용할 수 있는 것은 잘 분리된 코드다.

EX) 1
```go
// 적용전
func SaveConfig(filename string, cfg *Config) error {
	// JSON 형식으로 변환
	data, err := json.Marshal(cfg)
	if err != nil {
        return err
    }
	
	// 파일에 저장
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		log.Printf("failed to save file '%s' with err: %s", filename, err)
        return err
    }
	
	return nil
}

// 적용후
func SaveConfig(filename string, cfg *Config) error {
	// JSON 형식으로 변환
	data, err := json.Marshal(cfg)
	if err != nil {
        return err
    }
	
	// 파일에 저장
	err := writeFile(filename, data, 0644)
	if err != nil {
		log.Printf("failed to save file '%s' with err: %s", filename, err)
        return err
    }
	
	return nil
}

// 다음과 같이 writeFile 함수를 분리해서 writeFile을 모의로 손쉽게 변경할 수 있다.
var WriteFile = ioutil.WriteFile
```
#### 2. 다른 패키지의 내부 구조를 완전히 이해하지 못하더라도 그 패키지에 대한 모의를 생성할 수 있다.
> Monkey patching은 다른 모의의 형태와 마찬가지로, 의존성 내부에는 신경 쓰지 않고 원하는 대로 동작하도록 하는 능력을 제공한다.

#### 3. Monkey patching을 통한 DI는 기존 코드에 미치는 영향을 최소화한다.
#### 4. Monkey patching을 활용한 DI를 통해 전역 및 싱글톤 테스트 가능
EX) 2
다음은 rand package의 convenience 함수인 Int를 테스트하는 코드다.
```go
func TestInt(t *tesging.T) {
    // Monkey patch
    defer func(original *Rand) {
        globalRand = original
    }(globalRand)

    // 예측 가능한 결과를 위해 교환한다.
    globalRand = New(&stubSource {})
    // Monkey patch 종료

    // 함수를 호출한다.
    result := Int()
    assert.Equal(t, 42, result)
}

// 다음은 예측 가능한 값을 반환하는 스텁 구현체
type stubSource struct {}

func (s *stubSource) Int63() int64 {
    return 42
}
```
### Monkey patching 적용

```go
func TestData_happyPath(t *testing.T) {
    in := &Person {
        FullName: "Grace Hopper",
        Phone: "555-123-4567",
        Currency: "USD",
        Price: 100,
    }

    // Save
    resultID, err := Save(in)
    require.Nil(t, err)
    assert.True(t, resultID > 0)

   // Load
    returned, err := Load(resultID)
    require.NoError(t, err)

    in.ID = resultID
    assert.Equal(t, in, returned)

    // LoadAll
    all, err := LoadAll()
    require.NoError(t, err)
    assert.True(t, len(all) > 0)
}
```
위 테스트 코드의 문제점은 다음과 같다.
1. 행목 경로에 대한 테스트만 진행하고 있을 뿐, 에러 처리에 대한 부분은 전혀 테스트하지 않고 있다.
2. 테스트 코드가 데이터베이스에 의존하고 있다.
3. 모든 함수를 격리시키지 않고 함께 테스트하고 있다.
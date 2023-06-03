Single Responsibility Principle
---
> A class should have only one reason to change. - Robert C. Martin

## SRP는 DI와 어떤 관련이 있는가?
DI를 코드에 적용할 때는 일반적으로 의존성 객체를 함수의 매개변수로 받은 뒤 사용한다. 코드에 삽입된 의존성이 많은 함수가 있는 경우는 해당 메서드가 너무 많은 
역할을 수행하고 있다는 것을 의미한다.

보통 공통으로 필요한 코드 영역을 추출해 commons 또는 utils 패키지로 만들고 싶은 유혹에 빠지게 된다. 하지만 프로그래머는 이러헌 유혹에서 빠져나와야 한다.
코드를 재사용하는 것은 전적으로 옳지만, 이처럼 일반적인 패키지 이름을 사용하는 것은 옳지 않다. 이러한 패키지는 기본적으로 명확한 목적이 없기 때문에 SRP원칙에 위배된다. :sparkles:

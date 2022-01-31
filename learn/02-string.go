package main

import "fmt"

func main() {

	// 문자열 리터럴은 Back Quote(``) 혹은 쌍따옴표("")를 사용하여 표현
	/* 1. Back Quote (` `)로 둘러 싸인 문자열은 Raw String Literal이라 부르는데,
	이 안에 있는 문자열은 별도로 해석되지 않고 Raw String 그대로의 값을 갖는다.
	예를 들어, 문자열 안에 \n 이 있을 경우 이는 NewLine으로 해석되지 않는다.
	또한, Back Quote은 복수 라인의 문자열을 표현할 때 자주 사용된다*/

	/* 2. 이중인용부호(" ")로 둘러 싸인 문자열은 Interpreted String Literal이라 부르는데,
	복수 라인에 걸쳐 쓸 수 없으며, 인용부호 안의 Escape 문자열들은 특별한 의미로 해석된다.
	예를 들어, 문자열 안에 \n 이 있을 경우 이는 NewLine으로 해석된다.
	이중인용부호를 이용해 문자열을 여러 라인에 걸쳐 쓰기 위해서는 + 연산자를 이용해 결합하여 사용한다.*/

	rawLiteral := `아리랑\n
	아리랑\n
	아라리오`

	interLiteral := "아리랑아리랑\n아라리오"

	fmt.Println(rawLiteral)
	fmt.Println(interLiteral)
}

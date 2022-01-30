package main

func main() {
	var a int
	var f float32 = 11
	// 사용하지 않는 변수는 에러를 발생시킨다
	a = 10
	f = 12.0

	var i, j, k int

	var i, j, k int = 1, 2, 3

	// 변수를 선언하지 않으면 Go는 Zero Value를 기본적으로 할당한다.
	// 숫자는 0, bool false, string ""(빈문자열)

	var i = 1    //정수형으롷 할당
	var s = "HI" //문자열로 할당

	// 상수는 const
	const c int = 10
	const s string = "HI"

	const c = 10
	const s = "Hi"

	const (
		Visa   = "Visa"
		Master = "MasterCard"
		Amex   = "American Express"
	)

	const (
		Apple = iota
		Grape
		Orange
	)
}

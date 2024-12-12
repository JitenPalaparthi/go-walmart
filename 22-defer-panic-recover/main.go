package main

import "errors"

func main() {

	num := 100

	for i := num; ; i-- {
		if i%2 == 1 {
			continue
		}
		if i != 0 { // comment and uncomment based on panics, to check
			println(100 / i)
		}
		if i == 0 {
			break
		}
	}

	//num = 100

	f1 := func(a, b int) error {
		if a == 0 || b == 0 {
			return errors.New("a or b is zero")
		}
		println("anonymous function", a/b)
		return nil
	}

	if err := f1(20, 0); err != nil { // give a or b  zero to check panic
		panic(err)
	}
}